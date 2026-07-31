import matplotlib.pyplot as plt
from rate import res,bench
import csv
import math

def writedown(results):
    with open("results.csv",mode = 'w')as f :
        a = csv.writer(f,delimiter = ",")
        a.writerow(["number_of_bodies","time"])        
        for k in sorted(results.keys()):
            a.writerow([k,results[k]])

n = list()

for _ in [10,20,30,40,50,60,70,80,90,100,110,120,130,140,150]:
    n.append(_)

results = bench(n)

x = sorted(results.keys())
y = [results[k] for k in x]

fig,ax = plt.subplots()


writedown(results)
plt.title("Big O for python simulation")
plt.xlabel('Number of bodies')
plt.ylabel('Time')
plt.plot(x,y)
plt.savefig('python_results.png')
plt.show()

data = results
items = sorted(data.items())
slopes = []
for i in range(1, len(items)):
    n1, t1 = items[i-1]
    n2, t2 = items[i]
    if t1 <= 0 or t2 <= 0:
        continue
    k = (math.log(t2) - math.log(t1)) / (math.log(n2) - math.log(n1))
    slopes.append(k)
    sum = 0
    for i in slopes:
        sum += i
    av =  sum / len(slopes)    
a,b = slopes,av
print(a)
print(f"The efficiency of this program is defined by O(n^{b:.3f})")

