from functools import reduce

def gcd(a,b): 
    if(b==0): 
        return a 
    else: 
        return gcd(b,a%b) 

n = int(input())
A = list(map(int, input().split()))
maxGcd = 0

L = [0] * (n + 2)
R = [0] * (n + 2)

for i in range(2, n+2):
    L[i] = gcd(L[i-1], A[i-2])

for i in range(n, 0, -1):
    R[i] = gcd(R[i+1], A[i-1])

for i in range(1, n+1):
    maxGcd = max(maxGcd, gcd(L[i-1], R[i]))

print(maxGcd)
