class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0 or (x % 10 == 0 and x is not 0):
            return False
        
        revertN = 0
        while (x > revertN):
            revertN = int(revertN * 10 + x % 10)
            x = int(x / 10)
            
        return x == revertN or x == int(revertN / 10)