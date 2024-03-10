from typing import *

# @leet start
class Solution:
    def countBits(self, n: int) -> List[int]:
        sol = [0]*(n+1)
        offset = 1
        for x in range(1,n+1):
            if offset * 2 == x:
                offset = x
            sol[x] = 1 +sol[x-offset]

        return sol
        
# @leet end
