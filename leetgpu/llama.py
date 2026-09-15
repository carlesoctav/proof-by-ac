import jax
import jax.numpy as jnp
from functools import partial


# x, weights, cos, sin are tensors on device
@partial(jax.jit, static_argnums = 4)
def solve(
    x: jax.Array,
    weights: jax.Array,
    cos: jax.Array,
    sin: jax.Array,
    seq_len: int,
) -> jax.Array:

    # return output tensor directly
    # x = (T, nh)
    # w1 = (T, )
    # wq = (nh, nh) -> x wq = (T, nh) -> (T, n, h)
    # wk = (nh/4, nh) -> x wk = (T, n/4h) -> repeat (T, nh)
    # attn_weights = (T, nh ) (T, nh)^T  -> (T, T)
    # out = (T, T) (T, nh) -> (T, nh)
    # 

    T, S= seq_len, seq_len
    N, K = 8, 2

    def dot_product_attention(q, k, v):
        T, N, H = q.shape
        S, N, H = k.shape
        min_score = jnp.finfo(q.dtype).min
        mask = jnp.arange(T)[:, None] >= jnp.arange(S)[None, :] #(T,S)
        mask = mask[:, None, :]
        attn_score = jnp.einsum("tnh,snh-> tns", q, k) / (H**0.5)
        attn_score = jnp.where(mask, attn_score, min_score)
        attn_weights = jax.nn.softmax(attn_score, axis = -1)
        attn_out = jnp.einsum("tns,snh -> tnh", attn_weights, v)
        return attn_out

    def rms_norm(x, w): 
        #(T, D)
        rms = jnp.sqrt(jnp.square(x).mean(-1, keepdims = True) + 1e-5) # (T, 1)
        x = (x/rms) * w
        return x

    w = {}
    increment = {
            "w1": (512, ), 
            "wq": (512, 512), 
            "wk": (128, 512),
            "wv": (128, 512),
            "wo": (512, 512),
            "w2": (512, ),
            "gate": (1408, 512),
            "up": (1408, 512),
            "down": (512, 1408),
    }

    start = 0
    for wname, wshape in increment.items():
        stride =  wshape[0] * wshape[1] if len(wshape) == 2 else wshape[0]
        w[f"{wname}"] = weights[start: start + stride].reshape(*wshape)
        start = start + stride

    x_out = rms_norm(x, w["w1"]) #(T, D) 

    q = jnp.einsum("td,od->to", x_out, w["wq"])
    k = jnp.einsum("sd,od-> so", x_out, w["wk"])
    v = jnp.einsum("sd,od-> so", x_out, w["wv"])

    q = q.reshape(T, N, -1) #(T, N, H)
    k = k.reshape(T, K, -1)
    v = v.reshape(T, K, -1)

    H = q.shape[-1]
    sin, cos = sin[:, None, :], cos[:, None, :]
    q1, q2 = q[:, :, :H//2], q[:, :, H//2:] #(T, N, H/2) cos :  (T, N)
    k1, k2 = k[:, :, :H//2], k[:, :, H//2:]

    q = jnp.concatenate([q1 * cos - q2 * sin, q2 * cos + q1 * sin], axis=-1)
    k = jnp.concatenate([k1 * cos - k2 * sin, k2 * cos + k1 * sin], axis=-1)

    k  = jnp.repeat(k, N // K, axis = -2)
    v  = jnp.repeat(v, N // K, axis = -2)

    attn = dot_product_attention(q, k, v).reshape(T, -1) #(t, nh)
    x_out = jnp.einsum("to,do -> td", attn, w["wo"]) + x

    x_norm2 = rms_norm(x_out, w["w2"])
    upgate = jax.nn.silu(jnp.einsum("td,fd-> tf", x_norm2, w["gate"])) * jnp.einsum("td,fd ->tf", x_norm2, w["up"])
    x_out = jnp.einsum("tf,df -> td", upgate, w["down"]) + x_out
    return x_out

