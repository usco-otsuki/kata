import fileinput

s = 0

for line in fileinput.input():
    n = int(line)
    s += n // 3 - 2

print(s)
