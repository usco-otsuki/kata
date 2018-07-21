s = input()
t = input()

def match(a, b):
    for i in range(len(a)):
        if a[i:] + a[:i] == b:
            return True

    return False

if match(s, t):
    print("Yes")
else:
    print("No")

