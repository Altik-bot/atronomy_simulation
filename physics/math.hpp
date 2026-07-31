#ifndef MATH_HPP
#define MATH_HPP

#include <vector>
using namespace std;

struct Body {
    float mass;
    float pos_x;
    float pos_y;
    float vel_x;
    float vel_y;
    float force_x;
    float force_y;
    vector<float> path_x;
    vector<float> path_y;
};

void calc(float g, vector<Body>& arr, float t, float eps);
vector <Body> degenerate(float n,float mass,float pos_x,float pos_y,float vel_x,float vel_y);
void save_to_file(const vector<Body>& bodies);
void simulate(vector <Body> b, int n);

#endif