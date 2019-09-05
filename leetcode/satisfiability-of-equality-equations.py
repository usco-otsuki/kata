class DisjointSet:
    def __init__(self, vertices, parent):
        self.vertices = vertices
        self.parent = parent

    def find(self, item):
        if self.parent[item] == item:
            return item
        else:
            return self.find(self.parent[item])

    def union(self, set1, set2):
        root1 = self.find(set1)
        root2 = self.find(set2)
        self.parent[root1] = root2

class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        vertices = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
        parent = {}
        for v in vertices:
            parent[v] = v
        
        ds = DisjointSet(vertices, parent)
        for eq in equations:
            if "==" in eq:
                ds.union(eq[0], eq[3])
                
        for eq in equations:
            if "!=" in eq:
                if ds.find(eq[0]) == ds.find(eq[3]):
                    return False
        
        return True
