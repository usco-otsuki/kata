n = int(input())
A = list(map(int, input().split()))

s = sum(map(abs, A))
if sum(1 for a in A if a < 0) % 2 == 0:
    print(s)
else:
    print(s - 2 * min(map(abs, A)))
