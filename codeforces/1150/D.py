from collections import defaultdict

n, q = list(map(int, input().split()))
s = list(input())
indices = defaultdict([])
char_index_ptrs = defaultdict(None)
region_ptrs = (None, None, None)
for i, c in enumerate(s):
    indices[c].append(i)

for c in indices: 
    char_index_ptrs[c] = 0

for i in range(n):
    sign, region, char = input().split()
    if sign == "+":
        if char_index_ptrs[char]
        pass
    else:
        pass
