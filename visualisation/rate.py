import time
import math
from auto_m import generate,simulate
def run_sim(n):
    bodies = generate(n,1000,250,250,20,20)
    start = time.time()
    simulate(bodies,10000)
    end = time.time()
    return abs(start - end)
def bench():
    results = {}
    for v in [2,4,16,32,64]:
        t = run_sim(v)
        results[v] = t
        print(f"For {v} bodies, time is {t:.4f}s")
    return results     
def calc_bigO():
    data = bench()
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

    return slopes,av
a,b = calc_bigO()
print(a)
print(f"The efficiency of this program is defined by O(n^{b:.3f})")
