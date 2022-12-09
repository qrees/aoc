class Solution:
    def bisect(self, nums, item, low, hight):
        if low == hight:
            return low
        if nums[low] == item:
            return low
        if nums[hight] == item:
            return hight

        median_loc = int((low + hight) / 2)

        if item > nums[median_loc]:
            return self.bisect(nums, item, median_loc + 1, hight)
        elif item < nums[median_loc]:
            return self.bisect(nums, item, low, median_loc)
        return median_loc

    def findMedianSortedArrays(self, nums1, nums2):
        len_num1 = len(nums1)
        len_num2 = len(nums2)
        low1 = 0
        high1 = len(nums1) - 1
        low2 = 0
        high2 = len(nums2) - 1
        while True:
            median1 = (low1 + high1) / 2
            median2 = (low2 + high2) / 2
            if median1 + median2 < (len_num1 - median1) + (len_num2 - median2):
                if nums1[median1] < nums2[median2]:
                    low1 = median1
                else:
                    low2 = median2
            else:
                if nums1[median1] < nums2[median2]:
                    low1 = median1
                else:
                    low2 = median2
                

s = Solution()

print(s.bisect([1,2,3,5,6,7,8], 4, 0, 6))
