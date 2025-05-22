# pyright: basic
class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        s = set()
        for c in sentence.lower():
            s.add(c)
        return len(s) == (ord("z") - ord("a") + 1)
