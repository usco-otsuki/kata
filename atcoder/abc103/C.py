import functools

def gcd(a,b):
    while b > 0:
        a, b = b, a % b
    return a
    
def lcm(a, b):
    return a * b / gcd(a, b)

N = int(input())
a = list(map(int, input().split()))

lcmVal = int(functools.reduce(lcm, a, 0))
print(sum(map(lambda x: (lcmVal-1) % x, a)))
