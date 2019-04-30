n = int(input())
Vi = map(int, input().split())
Ci = map(int, input().split())

print(sum([max(0, v-c) for (v, c) in zip(Vi, Ci)]))
