d = [[1000] * n for _ in range(n)]
n = int(input())
m = int(input())
for _ in range(m):
    a, b, c = map(int, input().split())
    a, b = min(a, b), max(a, b)
    d[a][b] = c


for k in range(n):
    for i in range(n):
        for j in range(i + 1, n):
            d[i][j] = min(d[i][j], d[min(i, k)][max(i, k)] + d[min(j, k)][max(j, k)])
