class Solution:
    def longestValidParentheses(self, s: str) -> int:
        stack = [(")", -1)]
        latest_max = 0
        for index, letter in enumerate(s):
            if letter == "(":
                stack.append((letter, index))
            else:
                if stack[-1][0] == ")":
                    stack.append((letter, index))
                else:
                    stack.pop()
                    latest_max = max(index - stack[-1][1], latest_max)
        return latest_max


s = Solution()
print(s.longestValidParentheses(")()())"))