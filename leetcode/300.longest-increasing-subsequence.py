from typing import *

# @leet start
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        dp = [0  for i in range (len(nums))]

        for i in range(0, len(dp)):
            for j in range(i):
                if nums[j] < nums[i]:
                    dp[i] = max(dp[i], 1+dp[j])

        return max(dp)+1


        
# @leet end
