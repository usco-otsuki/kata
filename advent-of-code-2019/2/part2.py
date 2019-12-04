
def main():
    orig_li = list(map(int, input().split(",")))
    for noun in range(100):
        for verb in range(100):
            p = 0
            li = orig_li[:]
            li[1] = noun
            li[2] = verb
            while True:
                if li[p] == 99:
                    break
                elif li[p] == 1:
                    li[li[p+3]] = li[li[p+1]] + li[li[p+2]]
                elif li[p] == 2:
                    li[li[p+3]] = li[li[p+1]] * li[li[p+2]]

                p += 4

            if li[0] == 19690720:
                print(f"100 * {noun} + {verb} = {100 * noun + verb}")
                return

if __name__ == "__main__":
    main()
