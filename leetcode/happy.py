class Solution:
    def get_digits(self, n: int):
        return list(map(lambda x: int(x), list(str(n))))

    def isHappy(self, n: int) -> bool:
        cur = n
        for _ in range(100):
            cur = sum(digit**2 for digit in self.get_digits(cur))
            if cur == 1:
                return True

            if cur == n:
                return False

        return False

if __name__ == "__main__":
    solver = Solution()
    assert(solver.isHappy(19) == True)
    assert(solver.isHappy(2) == False)
            
