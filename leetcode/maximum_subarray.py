def max_sub_array(nums):
    dp = [0] * len(nums)
    dp[0] = nums[0]
    for i in range(1, len(nums)):
        if dp[i-1] < 0:
            dp[i] = nums[i]
        else:
            dp[i] = nums[i] + dp[i-1]

    return max(dp)

print(max_sub_array([-2,1,-3,4,-1,2,1,-5,4]))
    
