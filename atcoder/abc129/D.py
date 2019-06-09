H, W = map(int, input().split())

m = []
for i in range(H):
    m.append(list(input()))

left = [[0] * W for _ in range(H)]
up = [[0] * W for _ in range(H)]
right = [[0] * W for _ in range(H)]
down = [[0] * W for _ in range(H)]

total = [[0] * W for _ in range(H)]
for h in range(H):
    for w in range(W):
        if m[h][w] == "#":
            left[h][w] = 0
            up[h][w] = 0
            continue

        left[h][w] = 1
        if w > 0:
            left[h][w] += left[h][w-1]

        up[h][w] = 1
        if h > 0:
            up[h][w] += up[h-1][w]


for h in range(H-1, -1, -1):
    for w in range(W-1, -1, -1):
        if m[h][w] == "#":
            right[h][w] = 0
            down[h][w] = 0
            continue

        right[h][w] = 1
        if w < W - 1:
            right[h][w] += right[h][w+1]

        down[h][w] = 1
        if h < H - 1:
            down[h][w] += down[h+1][w]

for h in range(H):
    for w in range(W):
        total[h][w] = left[h][w] + right[h][w] + up[h][w] + down[h][w]
        total[h][w] -= 3

maxVal = max([max(row)for row in total])
print(maxVal)
