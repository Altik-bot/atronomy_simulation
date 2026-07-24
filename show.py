import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
from maths import bodies

fig, ax = plt.subplots()
lines = []
scatters = []
for line in range(0,len(bodies)):
    line, = ax.plot([], [])
    lines.append(line)
for i in range(0,len(bodies)):
    scatter = ax.scatter(bodies[i].path_x[0],bodies[i].path_y[0],bodies[i].mass)
    scatters.append(scatter)

 

def update(frame):
    N = 250
    start = max(0,frame-N)

    cx,cy = bodies[1].path_x[frame],bodies[1].path_y[frame]
    
    ax.set_xlim(cx-20,cx+20)
    ax.set_ylim(cy-20,cy+20)
    for line in range(0,len(bodies)):
        lines[line].set_data(bodies[line].path_x[start:frame], bodies[line].path_y[start:frame])
    for i in range(0,len(bodies)):
        scatters[i].set_offsets([[bodies[i].path_x[frame-1],bodies[i].path_y[frame-1]]])
        
    return lines,scatters
ani = FuncAnimation(fig, update, frames=len(bodies[0].path_x), interval=10)
plt.axis('equal')
plt.show()
plt.ion()