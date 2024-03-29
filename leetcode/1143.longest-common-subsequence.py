from typing import *

# @leet start
class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        text1 = "0" + text1
        text2 = "0" + text2
        dp = [[0 for i in range (len(text2))] for j in range(len(text1))] 
        for i in range (1,len(text1)):
            for j in range(1,len(text2)):
                if text1[i] == text2[j]:
                    dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
                else:
                    dp[i][j] = max(dp[i-1][j], dp[i][j-1])


        return dp[len(text1)-1][len(text2)-1]
        
# @leet end
