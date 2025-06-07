class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        parens = {
            ")": "(",
            "}": "{",
            "]": "[",
        }
        for c in s:
            if stack and c in parens and stack[-1] == parens[c]:
                stack.pop()
            else:
                stack.append(c)
        return len(stack) == 0
