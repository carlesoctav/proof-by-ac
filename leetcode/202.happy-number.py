from typing import *

# @leet start
class Solution:
    def isHappy(self, n: int) -> bool:
        def sum_digit_square(n: int):
            sum = 0
            for x in str(n):
                sum+= int(x)**2
            return sum
        hashmap = {}
        while True:
            n = sum_digit_square(n)
            if hashmap.get(n):
                return False
            else:
                hashmap[n] = True

            if n == 1:
                return True



        
# @leet end
