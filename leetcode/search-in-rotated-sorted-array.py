class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        l, h = 0, n - 1 
        while h > l:
            m = (h + l) // 2
            if nums[m] == target:
                return m

            if target < nums[l] or target > nums[m-1]:
                l = m
            else:
                h = m

        return -1


if __name__ == "__main__":

# [3,4,5,6,7,8,9,10,1,2]
    pass
