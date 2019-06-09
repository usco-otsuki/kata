N = int(input())
W = list(map(int, input().split()))
total = sum(W)

subTotal = 0
minVal = 100 * 100
for i in range(N):
    subTotal += W[i]
    if minVal > abs((total - subTotal) - subTotal):
        minVal = abs((total - subTotal) - subTotal)

print(minVal)

