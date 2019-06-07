x, y, z = map(int, input().split())
t = x - y
if abs(t) > z:
    if t > 0:
        print("+")
    else:
        print("-")
elif z == 0:
    if t > 0:
        print("+")
    if t == 0:
        print("0")
    else:
        print("-")   
else:
    print("?")
