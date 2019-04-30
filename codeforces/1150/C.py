n  = int(input())
a = list(map(int, input().split()))

count_one = sum(val == 1 for val in a)
count_two = n - count_one
if count_two > 0:
    print("2", end=" ")
    count_two -= 1

if count_one > 0:
    print("1", end=" ")
    count_one -= 1

print("2 " * count_two, end="")
print("1 " * count_one)
