class Solution:
    def checkValidString(self, s: str) -> bool:
        l, h = 0, 0
        for c in s:
            l += 1 if c == "(" else -1
            h += 1 if c != ")" else -1
            if h < 0:
                return False

            l = max(l, 0)

        return l == 0


if __name__ == "__main__":
    solver = Solution()
    assert solver.checkValidString("()")
    assert solver.checkValidString("(*)")
    assert solver.checkValidString("(*))")
    assert solver.checkValidString("(*")
    assert solver.checkValidString("***")
    assert not solver.checkValidString("*)(")
    assert not solver.checkValidString("())")
