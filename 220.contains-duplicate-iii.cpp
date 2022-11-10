// brute force maybe
//   Time Limit Exceeded

/*
 * @lc app=leetcode id=220 lang=cpp
 *
 * [220] Contains Duplicate III
 */
#include <bits/stdc++.h>
using namespace std;
#define ABS(x) (((x)>=0)?(x):(-1*(x)))
// @lc code=start
class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) {
        for (long long distance = 1; distance <= k; distance++)
        {
            for (long long i = 0; i < nums.size()-distance; i++)
            {
                if(ABS((long long)nums[i] - (long long)nums[i+distance]) <= (long long)t)
                { 
                // why are (long long) needed??? 
                // because (vector<int>) don't actually store int?
                    return true;
                }
            }
        }
        return false;
    }
};
// @lc code=end

