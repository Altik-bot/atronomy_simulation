import matplotlib.pyplot as plt
from rate import res,bench
import csv
import math
import pandas as pd
data = pd.read_csv("results_cpp.csv")
fig, ax = plt.subplots()

data["bodies"] = pd.to_numeric(data["bodies"])
data["time"] = pd.to_numeric(data["time"])

x = data["bodies"]
y = data["time"]
plt.title("Big O for C++ simulation")
plt.xlabel('Number of bodies')
plt.ylabel('Time(seconds)')
plt.plot(x,y)
plt.savefig('cpp_results.png')
plt.show()