from functools import partial

import jax
import jax.numpy as jnp


# Q, K, V are tensors on device
@partial(jax.jit, static_argnums = (3, 4, 5))
def solve(Q: jax.Array, K: jax.Array, V: jax.Array, N: int, d_model: int, h: int) -> jax.Array:
    # return output tensor directly
    # q: n h, k : nh, v: nh
    # h is numhead
    # N is sequence
    # dmodel is full head size
    # dk is head_size
    T, N, H = N, h, d_model//h


    Q = Q.reshape(T, N, H).astype(jnp.float32)
    K = K.reshape(T, N, H).astype(jnp.float32)
    V = V.reshape(T, N, H).astype(jnp.float32)

    attn_score = jnp.einsum("tnh,snh -> tns", Q, K, preferred_element_type= jnp.float32) // (H ** 0.5)
    attn_weights = jax.nn.softmax(attn_score, axis = -1, ).astype(jnp.float32)
    out = jnp.einsum("tns,snh->tnh", attn_weights, V).reshape(T, -1)
    return out

