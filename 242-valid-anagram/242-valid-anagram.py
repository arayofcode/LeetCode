class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        first = {}
        second = {}

        for a, b in zip(s, t):
            if a not in first:
                first[a] = 0
            if b not in second:
                second[b] = 0
            first[a] += 1
            second[b] += 1

        for a in first.keys():
            if a not in second:
                return False
            
            if first[a] != second[a]:
                return False
        
        if(set(first.keys()) == set(second.keys())):
            return True
        
        return False