from typing import List

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        ptr = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[ptr] = nums[i]
                ptr += 1

        while ptr < len(nums):
            nums[ptr] = 0
            ptr +=1


if __name__ == "__main__":
    solver = Solution()
    a = [0,1,0,3,12]
    solver.moveZeroes(a)
    assert a == [1,3,12,0,0
