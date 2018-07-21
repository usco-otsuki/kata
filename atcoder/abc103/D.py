N, M = list(map(int, input().split()))
intervals = []
for i in range(M):
    a, b = map(int, input().split())
    newInterval = None
    index = None
    for j, (s, e) in enumerate(intervals):
        if b <= s or e <= a:
            continue

        newInterval = (max(a, s), min(b, e))
        index = j
        break

    if newInterval is None:
        intervals.append((a, b))
    else:
        intervals[index] = newInterval

print(len(intervals))
