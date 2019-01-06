import sys

num = [3, 7]
after = int(sys.stdin.readline().strip())
p1, p2 = 0, 1

while True:
    num += list(map(int, list(str(num[p1] + num[p2]))))
    p1 = (p1 + num[p1] + 1) % len(num)
    p2 = (p2 + num[p2] + 1) % len(num)
    if len(num) >= after + 10:
        print("".join(map(str, num[after:after+10])))
        break
