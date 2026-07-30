#include <iostream>
#include <vector>
#include <cmath>
#include <fstream>
#include <random>
using namespace std;
struct Body{
    float mass;
    float pos_x;
    float pos_y;
    float vel_x;
    float vel_y;
    float force_x;
    float force_y;
    vector <float> acceleration;
    vector <float> path_x;
    vector <float> path_y;

};
vector <Body> degenerate(float n,float mass,float pos_x,float pos_y,float vel_x,float vel_y){
    vector <Body> bodies;
    random_device rd;
    mt19937 gen(rd());
    uniform_real_distribution<float> massS(0,mass);
    uniform_real_distribution<float> pos_X(-pos_x,pos_x);
    uniform_real_distribution<float> pos_Y(-pos_y,pos_y);
    uniform_real_distribution<float> vx(-vel_x,vel_x);
    uniform_real_distribution<float> vy(-vel_y,vel_y);
    for (int i = 0; i < n; i++){
    float massa = massS(gen);
    float x = pos_X(gen);
    float y = pos_Y(gen);
    float vpx = vx(gen);
    float vpy = vy(gen);
    
    bodies.push_back({massa,x,y,vpx,vpy});
    }
    return bodies;
}
void save_to_file(const vector<Body>& bodies) {
    ofstream file("output.csv");

    for (int i = 0; i < bodies.size(); i++) {
        for (int j = 0; j < bodies[i].path_x.size(); j++) {
            file << j << ","      // step
                 << i << ","      // body id
                 << bodies[i].path_x[j] << ","
                 << bodies[i].path_y[j] << "\n";
        }
    }

    file.close();
}
void calc(float g, vector <Body>&arr,float t,float eps){
    for (int i = 0; i < arr.size();i++){
        arr[i].force_x = 0;
        arr[i].force_y = 0; 
    }
    for(int i = 0; i < arr.size();i++){
        for(int j = 0; j < arr.size();j++){
            if (i == j){
                continue;
            }
            float r = sqrt(pow(arr[i].pos_x - arr[j].pos_x,2) + pow(arr[i].pos_y - arr[j].pos_y,2)+ pow(eps,2));
            float d[2] = {arr[i].pos_x-arr[j].pos_x , arr[i].pos_y-arr[j].pos_y};
            float F = g * (arr[i].mass * arr[j].mass)/(pow(r,3));
            arr[i].force_x += d[0] * -F;
            arr[i].force_y += d[1] * -F;
    }
    }
    for(int i = 0; i < arr.size();i++){
        arr[i].acceleration.resize(2);
        arr[i].acceleration[0] = arr[i].force_x/arr[i].mass;
        arr[i].acceleration[1] = arr[i].force_y/arr[i].mass;
        
        arr[i].vel_x += (arr[i].acceleration[0]*t);
        arr[i].vel_y += (arr[i].acceleration[1]*t);

        arr[i].pos_x += (arr[i].vel_x * t);
        arr[i].pos_y += (arr[i].vel_y * t);

        arr[i].path_x.push_back(arr[i].pos_x);
        arr[i].path_y.push_back(arr[i].pos_y);
    }
}
int main(){

    
    vector<Body> b = degenerate(5,10000,250,250,20,20);
    
    for (int i = 0; i <= 10000 ; i++ ){
    calc(1,b,0.01,0.001);

    }
    cout<<b.size()<<endl;
    save_to_file(b);
    return 0;
}
