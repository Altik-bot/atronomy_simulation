import matplotlib.pyplot as plt
import pandas as pd


cpp = pd.read_csv("results_cpp.csv")


py = pd.read_csv("results.csv")

cpp = cpp.sort_values("bodies")
py = py.sort_values("number_of_bodies")

plt.figure()

plt.plot(cpp["bodies"], cpp["time"], label="C++")
plt.plot(py["number_of_bodies"], py["time"], label="Python")

plt.title("C++ vs Python N-body Performance")
plt.xlabel("Number of bodies")
plt.ylabel("Time (seconds)")
plt.legend()

plt.savefig("comparison.png")
plt.show()