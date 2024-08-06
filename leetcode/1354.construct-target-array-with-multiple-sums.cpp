#include <bits/stdc++.h>
#include <numeric>

#include <queue>
using namespace std;

// @leet start
class Solution {
public:
    bool isPossible(vector<int>& target) {
        priority_queue<long long>pq;
        long long total_sum =0;
        for (auto a: target){
            pq.push(a);
            total_sum += a;
        }

        while (!pq.empty() && pq.top() > 1) {
            long long top = pq.top();
            pq.pop();
            long long remaining_sum = total_sum - top;

            if (remaining_sum == 0 || top <= remaining_sum) return false;

            long long new_top = top % remaining_sum;
            if (new_top == 0) new_top = remaining_sum; 

            pq.push(new_top);
            total_sum = remaining_sum + new_top;
        }

        return true;
    }

        
};
// @leet end

int main() {}
