from collections import defaultdict, Counter
from typing import List

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)     
        for s in strs:
            key = tuple(sorted(list(Counter(s).items())))
            groups[key].append(s)

        return list(groups.values())

if __name__ == "__main__":
    solver = Solution()
    print(solver.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))
