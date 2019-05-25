n, m = map(int, input().split())
a = list(map(int, input().split()))
a.sort()

li = []
for _ in range(m):
    b, c = map(int, input().split())
    li.append([c, b])

li.sort(reverse=True)

p = 0
q = 0
while p < len(a) and q < len(li):
    if a[p] < li[q][0]:
        a[p] = li[q][0]
        li[q][1] -= 1
        if li[q][1] == 0:
            q += 1
    else:
        break

    p += 1
            
print(sum(a))
