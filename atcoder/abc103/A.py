a1, a2, a3 = list(map(int, input().split()))
v1 = abs(a1 - a2) + abs(a2 - a3)
v2 = abs(a1 - a3) + abs(a2 - a3)
v3 = abs(a2 - a1) + abs(a1 - a3)
v4 = abs(a2 - a3) + abs(a3 - a1)
v5 = abs(a3 - a1) + abs(a1 - a2)
v6 = abs(a3 - a2) + abs(a2 - a1)
print(min(v1, v2, v3, v4, v5, v6))

