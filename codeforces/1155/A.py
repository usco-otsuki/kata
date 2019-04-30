def solve():
    n = int(input())
    s = list((input()))
    for i in range(len(s)-1):
        if s[i] > s[i+1]:
            print("YES")
            print(f"{i+1} {i+2}")
            return

    print("NO")

solve()

# aaabbb
# aabab
# caabc
#
# abacab
#
# abaczb
# bacz
