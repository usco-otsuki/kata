n = int(input())
a = list(map(int, input().split()))
b = list(map(int, input().split()))

def solve(n, a, b):
    num_zeros = sum(1 for el in a if el == 0)
    a = [el for el in a if el != 0]
    is_sorted = all(b[i] <= b[i+1] for i in range(len(b)-1))
    if is_sorted:
        if num_zeros == n:
            return 0
        
        return n - num_zeros
    
    dist = [0] * len(b)
    for i in range(len(b)):
        dist[b[i]-1] = i + 2

    max_val = 0
    max_index = -1
    for i in range(len(b)):
        if max_val < dist[b[i]-1] + n - (i + 1):
            max_val = dist[b[i]-1] + n - (i + 1)
            max_index = i

    max_index = 


11
0 0 0 5 0 0 0 4 0 0 11
9 2 6 0 8 1 7 0 3 0 10
1 2 3   5 6 7   9   11

1 2 3 4 5 6 7 8 9 10 11
-----------------------
6 2 9 x x 3 7 5 1 11 x
7 3 10          2 12

                  
8 9 10 11 12 13 14 15 16 17 18
1 2  3 4   5  6  7  8  9 10 11
   
1 2 0
0 3 0

