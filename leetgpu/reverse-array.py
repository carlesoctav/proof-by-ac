import jax
import jax.numpy as jnp


# input is a tensor on device
@jax.jit
def solve(input: jax.Array, N: int) -> jax.Array:
    output = input[::-1]
    # return output tensor directly
    return output
    pass




key = jax.random.key(12)
test = jax.random.normal(key, (10, ))
# print(f"DEBUGPRINT[1]: reverse-array.py:17: test={test}")
# print(f"DEBUGPRINT[3]: reverse-array.py:19: len(test)={len(test)}")
# out = solve(test, len(test))
lower = solve.lower(test, len(test))
hlo = lower.compile().as_text()
with open("/content/compile.hlo", "w") as f:
    f.write(hlo)
