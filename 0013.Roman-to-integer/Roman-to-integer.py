class Solution:
    def romanToInt(self, s: str) -> int:
        output = 0
        roman_table = {"IV": 4, "IX": 9, "I": 1, "V": 5,
                       "XL": 40, "XC": 90, "X": 10, "L":50, 
                       "CD": 400, "CM": 900, "C": 100, "D": 500, "M": 1000}
        
        for symbol, num in roman_table.items():
            if symbol in s:
                output += s.count(symbol) * num 
                s = s.replace(symbol, '')
        
        # print(output)
        return output
