// brute force

using namespace std;
// @before-stub-for-debug-end

/*
 * @lc app=leetcode id=1 lang=cpp
 *
 * [1] Two Sum
 */
#include <bits/stdc++.h>
using namespace std;
// @lc code=start
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> out;
        for (size_t i = 0; i < nums.size(); i++)
        {
            for (size_t j = i+1; j < nums.size(); j++)
            {
                if(nums[i]+nums[j]==target){
                    out.push_back(i);
                    out.push_back(j);
                    return out;
                }
            }
        }
        return out; //for all path
    }
};
// @lc code=end

