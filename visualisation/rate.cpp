#include <iostream>
#include <vector>
#include <chrono>
#include "math.hpp"
#include <unordered_map>
#include <cmath>
#include <algorithm>

using namespace std;
using namespace std::chrono;

float run_sim(int n){
    vector <Body> bodies = degenerate(n,10000,250,250,20,20);

    auto start = high_resolution_clock::now();

    simulate(bodies,10000); 

    auto end = high_resolution_clock::now();

    return duration_cast<milliseconds>(end-start).count();

}
unordered_map <int,float> bench(){
    int a[5] = {2,4,16,32,64};
    unordered_map <int,float> results;
    for (int i = 0; i < 5; i++){
    
    results[a[i]] = (run_sim(a[i])/1000);
    cout<<"For "<<a[i]<<" bodies, time is: "<<results[a[i]]<<" seconds"<<endl;
    }
    return results;
}
float calc_bigO(){
    unordered_map <int, float> data = bench();
    vector <float> slopes;
    vector<pair<int,float>> values;
    for (const auto& pair: data){
        values.push_back(pair);
    }
    sort(values.begin(), values.end(), [](const auto& a, const auto& b) {
        return a.first < b.first; 
    });

    for (size_t i = 1; i < values.size(); i++){
    auto [n1, t1] = values[i - 1];  
    auto [n2, t2] = values[i];
    float k = (log(t2)-log(t1))/(log(n2)-log(n1));
    slopes.push_back(k);

    }
    float sum = 0;
    for(size_t i = 0; i < slopes.size(); i++){
        sum+=slopes[i];
    }
    float av = sum/(slopes.size());
    return av;
}    

int main(){
   float res = calc_bigO();
   cout<<"The efficiency of this program is defined by O(n^"<<res<<") ";
}