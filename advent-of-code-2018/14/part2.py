import sys

num = "37"
target = sys.stdin.readline().strip()
p1, p2 = 0, 1

while target not in num[-len(target)-1:]: 
    num += str(int(num[p1]) + int(num[p2]))
    p1 = (p1 + int(num[p1]) + 1) % len(num)
    p2 = (p2 + int(num[p2]) + 1) % len(num)
    
index = "".join(num).index(target)
print(index)