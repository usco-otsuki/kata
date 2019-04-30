from math import gcd
from functools import reduce

def solve():
    input()
    x =list( map(int, input().split()))
    dx = [t - s for s, t in zip(x, x[1:])] 
    p = list(map(int, input().split()))
    d = reduce(gcd, dx)
    for i in range(len(p)):
        if d % p[i] == 0:
            print("YES")
            print(f"{x[0]} {i+1}")
            return
    
    print("NO")
    return

solve()
