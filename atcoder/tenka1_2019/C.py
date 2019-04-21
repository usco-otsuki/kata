n = int(input())
s = list(input())

w_sum = [0]
b_sum = [0]
for i in range(len(s)):
    if s[i] == '.':
        w_sum.append(w_sum[-1] + 1) 
        b_sum.append(b_sum[-1])
    else:
        b_sum.append(b_sum[-1] + 1)
        w_sum.append(w_sum[-1])

assert len(w_sum) == len(b_sum)
assert len(w_sum) == len(s) + 1

minSum = 10 ** 5
for i in range(len(s)):
    minSum = min(minSum, b_sum[i] + (w_sum[-1] - w_sum[i+1]))
    
print(minSum)
