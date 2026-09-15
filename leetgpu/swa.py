import jax
import jax.numpy as jnp
from functools import partial



# Q, K, V are tensors on device
@partial(jax.jit, static_argnums = (3,4,5))
def solve(Q: jax.Array, K: jax.Array, V: jax.Array, M: int, d: int, window_size: int) -> jax.Array:
    # return output tensor directly
    min = jnp.finfo(jnp.float32).min


    # mask = jnp.ones((M, M))
    qarange = jnp.arange(M)[:, None]
    karange = jnp.arange(M)[None, :]
    mask = (qarange >= karange - window_size) & (qarange <= karange + window_size)
    jax.debug.print("{mask}", mask = mask.astype(jnp.int8))
    print(f"DEBUGPRINT[5]: swa.py:17: mask={mask.astype(jnp.int8)}")


    attn_score = jnp.einsum("td,sd-> ts", Q, K) / (M**0.5)
    attn_score = jnp.where(mask, attn_score, min)
    attn_weights = jax.nn.softmax(attn_score, axis = -1)


    return jnp.einsum("ts,sd-> td", attn_weights, V)



key = jax.random.key(10)
m = 7
d = 4
window = 2
q = jax.random.normal(key, (m, d))
k = jax.random.normal(key, (m, d))
v = jax.random.normal(key, (m, d))

solve(q, k, v, m, d, window)
