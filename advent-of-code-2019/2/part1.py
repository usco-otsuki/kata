li = list(map(int, input().split(",")))
p = 0

li[1] = 12
li[2] = 2

while True:
    if li[p] == 99:
        break
    elif li[p] == 1:
        li[li[p+3]] = li[li[p+1]] + li[li[p+2]]
    elif li[p] == 2:
        li[li[p+3]] = li[li[p+1]] * li[li[p+2]]

    p += 4

print(li[0])

