#include <iostream>
#include <vector>
#include "math.hpp"
#include <unordered_map>
#include <fstream>
#include <algorithm>
#include <cmath>

using namespace std;

void save_to_files(unordered_map <int,float> result) {
    ofstream file("results_cpp.csv");
    vector<pair<int,float>> values;
    file << "bodies" << "," << "time" << "\n";
    for (const auto& pair: result){
        values.push_back(pair);
    }
    sort(values.begin(), values.end(),
     [](const auto& a, const auto& b){
         return a.first < b.first;
     });
    for (size_t i = 0; i < values.size(); i++) {
            auto [n2, t2] = values[i];
            file << n2 << ","
                 << t2 << "\n";
        
    }

    file.close();
}
int main(){
    vector <int> n = {10,20,30,40,50,60,70,80,90,100,110,120,130,140,150};
    unordered_map <int, float> result = bench(n);
    save_to_files(result);
    vector <float> slopes;
    vector<pair<int,float>> values;
    for (const auto& pair: result){
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
    if (slopes.size() == 0) return 0;
    float av = sum/(slopes.size());
    cout<<"The efficiency of this program is defined by O(n^"<<av<<")"<<endl;

}
