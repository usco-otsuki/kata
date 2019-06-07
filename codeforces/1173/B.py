import math

n = int(input())

m = int(math.ceil((n + 1) / 2))
print(m)
for i in range(1, min(n+1, m+1)):
    print(f"1 {i}")

for i in range(2, n - m + 2):
    print(f"{i} {m}")
