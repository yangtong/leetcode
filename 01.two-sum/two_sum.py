class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        m = {}
        n = 0
        for i in nums:
            if (target - i) in m:
                return m[target - i], n
            m[i] = n
            n = n + 1

