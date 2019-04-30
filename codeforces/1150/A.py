n, m, r = list(map(int, input().strip().split(" ")))
s = list(map(int, input().strip().split(" ")))
b = list(map(int, input().strip().split(" ")))
min_s = min(*s)
max_b = max(*b)

maxValue = max((r // min_s) * max_b + (r % min_s), r)
print(maxValue)
