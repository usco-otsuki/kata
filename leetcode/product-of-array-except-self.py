from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = [1] * len(nums)
        multiplied = 1
        for i in range(len(nums)-1):
            multiplied *= nums[i]
            result[i+1] *= multiplied

        multiplied = 1
        for i in range(len(nums)-1, 0, -1):
            multiplied *= nums[i]
            result[i-1] *= multiplied
            
        print(result)
        return result


if __name__ == "__main__":
    solver = Solution()
    assert solver.productExceptSelf([1,2,3,4]) == [24,12,8,6]
