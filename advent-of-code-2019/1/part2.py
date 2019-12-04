import fileinput

s = 0

for line in fileinput.input():
    n = int(line)
    while True:
        n = n // 3 - 2
        if n <= 0:
            break

        s += n

print(s)
