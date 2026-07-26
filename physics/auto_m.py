import math,time
import random

class Body:
    def __init__(self,mass,init_position_x,init_position_y,init_velocity_x,init_velocity_y,force_x,force_y):
        self.mass = mass
        self.position = [init_position_x,init_position_y]
        self.velocity = [init_velocity_x,init_velocity_y]
        self.force = [force_x , force_y]
        self.path_x = []
        self.path_y = []
def generate(n,mass,pos_x,pos_y,vel_x,vel_y):  
        bodies = []
        for i in range(n):
            mass =random.uniform(1,mass)
            x = random.uniform(-pos_x,pos_x)
            y = random.uniform(-pos_y,pos_y)
            vx = random.uniform(-vel_x,vel_x)
            vy = random.uniform(-vel_y,vel_y)

            bodies.append(Body(mass,x,y,vx,vy,0,0))
        return bodies           
                               

bodies = generate(5,10000,250,250,20,20)
g = 1
t = 0.01
def calc(g,bodies,t,eps):   
    for body in bodies:
        body.force[0] = 0
        body.force[1] = 0 
    for i in range(len(bodies)):
        for j in range(len(bodies)):   
            if i == j:
                continue 
            r = math.sqrt(( ( bodies[i].position[0] - bodies[j].position[0] ) ** 2 ) + ( ( bodies[i].position[1] - bodies[j].position[1] ) ** 2 ) + eps ** 2)
            d = (bodies[i].position[0]-bodies[j].position[0] , bodies[i].position[1]-bodies[j].position[1])
            F = g * (bodies[i].mass * bodies[j].mass) / r ** 3
            bodies[i].force[0] += d[0] * -F 
            bodies[i].force[1] += d[1] * -F
           

    for body in bodies:
        body.acceleration = (body.force[0]/body.mass, body.force[1]/body.mass)

        body.velocity[0] += (body.acceleration[0] * t)
        body.velocity[1] += (body.acceleration[1] * t) 

        body.position[0] += (body.velocity[0] * t)
        body.position[1] += (body.velocity[1] * t)  

        body.path_x.append(body.position[0])
        body.path_y.append(body.position[1])
def simulate(bodies,n):
    g = 1
    t = 0.001
    for _ in  range(n):
        calc(g,bodies,t,0.0001)
    print("done")
