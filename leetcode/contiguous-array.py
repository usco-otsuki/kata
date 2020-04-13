from typing import List

class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        d = {0: 0}
        max_len = 0
        cnt = 0
        for i, num in enumerate(nums):
            i += 1
            cnt += (1 if num == 1 else -1)
            if cnt in d:
                max_len = max(max_len, i - d[cnt])
            else:
                d[cnt] = i

        print(max_len)
        return max_len
        

if __name__ == "__main__":
    solver = Solution()
    assert solver.findMaxLength([0, 0, 1, 0, 0, 0, 1, 1]) == 6
    assert solver.findMaxLength([0, 1]) == 2


