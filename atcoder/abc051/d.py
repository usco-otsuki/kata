edges = []
n, m = map(int, input().split())
d = [[1000] * n for _ in range(n)]
for i in range(n):
    d[i][i] = 0

for _ in range(m):
    a, b, c = map(int, input().split())
    a -= 1
    b -= 1
    d[a][b] = c
    d[b][a] = c
    edges.append((a, b, c))

for k in range(n):
    for i in range(n):
        for j in range(i + 1, n):
            d[j][i] = d[i][j] = min(d[i][j], d[i][k] + d[j][k])

count = 0
for i, j, val in edges:
    isIncluded = False
    for s in range(n):
        if d[s][i] + val == d[s][j] or d[s][j] + val == d[s][i]:
            isIncluded = True
            break

    if not isIncluded:
        count += 1

print(count)
