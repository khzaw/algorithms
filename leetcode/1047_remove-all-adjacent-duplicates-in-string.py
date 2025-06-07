# This solves all cases where duplicates mean any number of consecutive number of same letters
# class Solution:
#     def removeDuplicates(self, s: str) -> str:
#         stack = []
#         for idx, c in enumerate(s):
#             if stack and stack[-1] == c:
#                 if (idx + 1) < len(s) and s[idx + 1] != stack[-1]:
#                     stack.pop()
#                 if idx + 1 == len(s) and stack[-1] == c:
#                     stack.pop()
#             else:
#                 stack.append(c)
#         return "".join(stack)


class Solution:
    def removeDuplicates(self, s: str) -> str:
        stack = []
        for c in s:
            if stack and stack[-1] == c:
                stack.pop()
            else:
                stack.append(c)
        return "".join(stack)
