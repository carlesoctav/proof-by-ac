#include <bits/stdc++.h>
#include <functional>
#include <queue>
#include <utility>
using namespace std;

// @leet start
class Solution {
public:
    int rangeSum(vector<int>& nums, int n, int left, int right) {
        priority_queue<pair<int,int>, vector<pair<int, int>>, std::greater<pair<int,int>>> pq;
        int ans = 0;
        for (int i = 0; i< n; i++){
            pq.push(make_pair(nums[i], i));
        };

        for (int i = 0 ; i < right; i++ ){
            pair<int, int> x;
            x = pq.top();
            pq.pop();
            if (i >= left -1){
                ans += x.first; 
                ans %= int(1e9+7);
            }

            if ( x.second + 1 < n ){
                pq.push(make_pair(x.first + nums[x.second + 1], x.second +1));
            }
        };

        return ans;
    }
};
// @leet end

int main() {}
