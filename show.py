import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
from maths import B1,B2

fig, ax = plt.subplots()
line1, = ax.plot([], [])
line2, = ax.plot([], [])
body1 = ax.scatter(B1.path_x[0],B1.path_y[0],10)
body2 = ax.scatter(B2.path_x[0],B2.path_y[0],10)
ax.set_xlim(-1050,1050)
ax.set_ylim(-1050,1050)

def update(frame):
    N = 250
    start = max(0,frame-N)
    line1.set_data(B1.path_x[start:frame], B1.path_y[start:frame])
    line2.set_data(B2.path_x[start:frame], B2.path_y[start:frame])
    body1.set_offsets([[B1.path_x[frame-1],B1.path_y[frame-1]]])
    body2.set_offsets([[B2.path_x[frame-1],B2.path_y[frame-1]]])
    return line1, line2, body1, body2

ani = FuncAnimation(fig, update, frames=len(B1.path_x), interval=10)
plt.axis('equal')
plt.show()
plt.ion()