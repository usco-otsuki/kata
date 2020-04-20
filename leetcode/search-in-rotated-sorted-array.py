from typing import List

class Solution:
    def find_pivot(self, nums: List[int], low: int, high: int) -> int:
        if high < low:
            return -1;

        if high == low:
            return low;

        mid = (low + high) // 2
        if mid < high and nums[mid] > nums[mid + 1]:
            return mid;

        if mid > low and nums[mid] < nums[mid - 1]:
            return mid-1

        if nums[low] >= nums[mid]:
            return self.find_pivot(nums, low, mid-1);

        return self.find_pivot(nums, mid + 1, high);

    def binary_search(self, nums: List[int], low: int, high: int, target: int) -> int:
        if high < low:
            return -1

        mid = (low + high) // 2
        if nums[mid] == target:
            return mid

        if target > nums[mid]:
            return self.binary_search(nums, mid + 1, high, target)

        return self.binary_search(nums, low, mid - 1, target)


    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        pivot = self.find_pivot(nums, 0, n-1)

        if pivot == -1:
            return self.binary_search(nums, 0, n-1, target)

        if nums[pivot] == target:
            return pivot

        if nums[0] <= target:
            return self.binary_search(nums, 0, pivot-1, target)

        return self.binary_search(nums, pivot+1, n-1, target)


if __name__ == "__main__":
    solver = Solution()
    assert solver.search([3,4,5,6,7,8,9,10,1,2], 10) == 7
    assert solver.search([3,4,5,6,7,8,9,10,1,2], 3) ==  0
    assert solver.search([3,4,5,6,7,8,9,10,1,2], 2) ==  9
    assert solver.search([1,2,3], 1) ==  0
    assert solver.search([4,5,6,7,0,1,2], 0) ==  4, solver.search([4,5,6,7,0,1,2], 0)
    assert solver.search([4,5,6,7,0,1,2], 3) ==  -1
