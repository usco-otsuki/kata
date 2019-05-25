n, m = map(int, input().split())
l, r = 1, 10**5
for _ in range(m):
    l_cur, r_cur = map(int, input().split())
    l = max(l_cur, l)
    r = min(r_cur, r)
    if r < l:
        break

print(max(r-l+1, 0))
