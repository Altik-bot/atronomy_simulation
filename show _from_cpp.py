import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
import pandas as pd

data = pd.read_csv("output.csv", header=None, names=["step","id","x","y"])

fig, ax = plt.subplots()
ax.set_aspect('equal')

# get unique bodies
bodies = data["id"].unique()

lines = []
points = []

for _ in bodies:
    line, = ax.plot([], [])
    point, = ax.plot([], [], 'o')
    lines.append(line)
    points.append(point)

def init():
    ax.set_xlim(-20,20)
    ax.set_ylim(-20,20)
    return lines + points

def update(frame):
    for i, body_id in enumerate(bodies):
        subset = data[(data["id"] == body_id) & (data["step"] <= frame)]

        if len(subset) == 0:
            continue

        x = subset["x"].values
        y = subset["y"].values

        lines[i].set_data(x, y)
        points[i].set_data([x[-1]], [y[-1]])

  
    main = data[(data["id"] == 0) & (data["step"] == frame)]

    if not main.empty:
        cx = main["x"].values[0]
        cy = main["y"].values[0]

        scale = 20 
        ax.set_xlim(cx - scale, cx + scale)
        ax.set_ylim(cy - scale, cy + scale)
    
    return lines + points

ani = FuncAnimation(
    fig,
    update,
    frames=data["step"].max(),
    init_func=init,
    interval=10
)

plt.show()