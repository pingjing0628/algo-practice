class Solution:
    def reverse(self, x: int) -> int:
        reve = 0
        a = abs(x)
        while a != 0:
            pop = int(a % 10)
            a = int(a / 10)
            reve = reve * 10 + pop

        if x < 0:
            reve = -reve

        if int(reve) >= 2**31-1 or int(reve) <= -2**31:
            return 0

        return reve