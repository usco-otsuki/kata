from typing import List

class Solution:
    def advantageCount(self, orig_a: List[int], orig_b: List[int]) -> List[int]:
        n = len(orig_a)
        a = list(zip(orig_a,[i for i in range(len(orig_a))]))
        b = list(zip(orig_b,[i for i in range(len(orig_b))]))
        sorted_a = sorted(a)
        sorted_b = sorted(b)

        not_used = []
        result = [None] * n 
        ptr_a, ptr_b = 0, 0
        while ptr_a < n:
            orig_a_index, orig_b_index = sorted_a[ptr_a][1], sorted_b[ptr_b][1]
            if sorted_a[ptr_a][0] <= sorted_b[ptr_b][0]:
                ptr_a += 1
                not_used.append(orig_a_index)
                continue

            result[orig_b_index] = sorted_a[ptr_a][0]
            ptr_a += 1
            ptr_b += 1

        ptr = 0
        for i in range(n):
            if result[i] is None:
                result[i] = orig_a[not_used[ptr]]
                ptr += 1

        return result

def main():
    solver = Solution()
    test_cases = [
        ([2,7,11,15], [1,10,4,11], [2,11,7,15]),
        ([12,24,8,32], [13,25,32,11], [24,32,8,12])
    ]

    for test_case in test_cases:
        assert solver.advantageCount(test_case[0], test_case[1]) == test_case[2]

if __name__ == "__main__":
    main()

