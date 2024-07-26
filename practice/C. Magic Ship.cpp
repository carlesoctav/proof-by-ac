#include <bits/stdc++.h>
#include <utility>
#include <vector>
using namespace std;
#define ll long long
#define x first 
#define y second

pair<int, int> st, en;
ll dx[4] = {1, -1, 0, 0};
ll dy[4] = {0, 0, 1, -1};
string mvs = "RLUD";
ll n;
string s;

void solve(){
    cin >> st.first >> st.second;
    cin >> en.first >> en.second;
    cin >> n;
    cin >> s;

    // Initialize prev_move after n is known
    vector<pair<int, int>> prev_move(n+1);

    for (int i = 0; i < n; i++){
        ll id = -1;
        for (int j = 0; j < 4; j++){
            if (s[i] == mvs[j]){
                id = j;
                break;  
            }
        }
        prev_move[i+1] = make_pair(prev_move[i].first + dx[id], prev_move[i].second + dy[id]);
    }

    ll ans = -1;
    ll lo = 0, hi = 1e18L; 

    while (lo <= hi){
        ll mid = lo + (hi - lo)/2;
        ll div = mid / n;  
        ll mod = mid % n;  
        ll xmid = st.first + div * prev_move[n].first + prev_move[mod].first;
        ll ymid = st.second + div * prev_move[n].second + prev_move[mod].second;

        if (abs(xmid - en.first) + abs(ymid - en.second) <= mid){
            ans = mid;
            hi = mid - 1;  
        } else {
            lo = mid + 1;  
        }
    }
    cout << ans << endl;
}

int main(){
    solve();
}
