import math,time

class Body:
    def __init__(self,mass,init_position_x,init_position_y,init_velocity_x,init_velocity_y,force_x,force_y):
        self.mass = mass
        self.position = (init_position_x,init_position_y)
        self.velocity = (init_velocity_x,init_velocity_y)
        self.force = (force_x , force_y)
        self.path_x = []
        self.path_y = []

B1 = Body(100,10,10,0,10,0,0)
B2 = Body(100000,100,100,0,0,0,0)
g = 1
t = 0.01
def calc(g,B1,B2,t):    
    r = math.sqrt(( ( B1.position[0] - B2.position[0] ) ** 2 ) + ( ( B1.position[1] - B2.position[1] ) ** 2 ))
    if r <= 0.00001:
        r = 0.00001
    d = (B2.position[0]-B1.position[0] , B2.position[1]-B1.position[1])
    F = g * (B1.mass * B2.mass) / r ** 3

    B1.force = (d[0]* F , d[1] * F)
    B2.force = (d[0] * -F , d[1] * -F)

    B1.acceleration = (B1.force[0]/B1.mass, B1.force[1]/B1.mass)
    B2.acceleration = (B2.force[0]/B2.mass, B2.force[1]/B2.mass)

    B1.velocity = (B1.velocity[0] + (B1.acceleration[0] * t), B1.velocity[1] + (B1.acceleration[1] * t)) 
    B2.velocity = (B2.velocity[0] + (B2.acceleration[0] * t), B2.velocity[1] + (B2.acceleration[1] * t)) 

    B1.position = (B1.position[0] + (B1.velocity[0] * t), B1.position[1] + (B1.velocity[1] * t))
    B2.position = (B2.position[0] + (B2.velocity[0] * t), B2.position[1] + (B2.velocity[1] * t))
    
    B1.path_x.append(B1.position[0])
    B1.path_y.append(B1.position[1])

    B2.path_x.append(B2.position[0])
    B2.path_y.append(B2.position[1])
for _ in  range(10000):
    calc(g,B1,B2,t)
    time.sleep(t)
    
