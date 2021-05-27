class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        a = list(enumerate(nums))
        print(a)

        temp = {}
        for index, num in enumerate(nums):
            nn = target - num
            print(nn)
            
            if nn not in temp:
                temp[num] = index
                print(temp)
            else:
                return [temp[nn], index]
