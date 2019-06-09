N, M = map(int, input().split())
a = {}
for i in range(M):
    a[int(input())] = True

comb = [0] * (N + 1)
comb[0] = 1
comb[1] = 0 if 1 in a else 1

for i in range(2, N+1):
    if i in a:
        comb[i] = 0
    else:
        comb[i] = (comb[i-1] + comb[i-2]) % 1000000007 

print(comb[N])
