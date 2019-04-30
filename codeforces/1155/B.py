# 8380011223344
# 00000000008444
# 0808080800008088888

# 000000000008880

def solve():
    n = int(input())
    s = list(input())
    a = []
    i, cnt = 0, 0
    while i < n:
        if cnt >= (n - 11)>>1:
            break

        if s[i] != "8":
            cnt += 1
        else:
            a.append("8")

        i += 1

    s = a + s[i:]
    a = []
    i, cnt = 0, 0
    while i < len(s):
        if cnt >= (n - 11)>>1:
            break

        if s[i] == "8":
            cnt += 1
        else:
            a.append(s[i])

        i += 1

    a += s[i:]

    if a[0] == "8":
        print("YES")
    else:
        print("NO")

solve()
