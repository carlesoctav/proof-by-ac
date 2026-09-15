import jax
import jax.numpy as jnp
T = 10
S = 5

q = jnp.ones(( T, T ))
k = jnp.ones((S, S)).

mask = q <= k
print(f"DEBUGPRINT[4]: masking.py:8: mask={mask}")
print(f"DEBUGPRINT[4]: masking.py:8: mask={mask.shape}")

