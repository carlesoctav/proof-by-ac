#include <bits/stdc++.h>
#include <math.h>
#include <vector>
using namespace std;
#define ll long long 


bool condition_mid(ll mid, vector<ll>a, vector<ll>b){
    ll i = 0;
    ll j = 0;
    ll n = a.size();
    ll m = b.size();
    while (i < n && j < m){
        if (abs(a[i] - b[j]) <= mid){
            i++;
        } else {
            j++;
        }
    }

    if (i == n){
        return true;
    }
    if ( j== m){
        return false;
    }

}


int main(){
    ll n, m;
    cin>>n>>m;
    vector<ll> a(n);
    vector<ll> b(m);
    ll temp;
     
    for (ll i=0; i<n;i++){
        cin>>temp;
        a[i] = temp;
    }

    for (ll i = 0; i< m; i++){
        ll temp;
        cin>> temp;
        b[i] = temp;
    }
    
    ll ans = -1;
    ll low = 0;
    ll hi = 1e18L;
    while (low <= hi){
        ll mid = low + (hi - low)/2;
        if (condition_mid(mid, a, b)){
            hi = mid -1;
            ans = mid;
        } else{
            low = mid+1;
        }
    }

    cout<<ans<<endl;
}
