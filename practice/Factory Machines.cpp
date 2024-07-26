#include <bits/stdc++.h>
#include <vector>
using namespace std;
#define ll long long 

bool condition_mid(ll mid, vector<ll> &k, ll t){
    ll sum = 0;
    for (auto a: k){
        sum+= mid / a;
        if (sum >= t){
            return true;
        }
    }
    if (sum >= t){
        return true;
    }

    return false;
}

void solve(){
    ll n, t;
    cin >>n>>t;
    vector<ll> k(n);
    for (int i = 0; i < n; i++){
        cin>> k[i];
    }
    ll lo = 0, hi = 1e18L; 
    ll ans = -1;
    while (lo <= hi){
        ll mid = lo + (hi - lo)/2;
        if (condition_mid(mid, k, t)){
            hi = mid -1;
            ans = mid;
        } else {
            lo = mid +1;
        }
    }
    cout<< ans;
}

int main(){
    solve();
}



