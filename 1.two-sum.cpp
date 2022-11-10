// Best - One pass hash

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
        map<int,int> hashmap;
        for (size_t i = 0; i < nums.size(); i++)
        {
            int complement = target - nums[i];
            if(hashmap.find(complement)!=hashmap.end()&&hashmap.find(complement)->second!=i){
                out.push_back(i);
                out.push_back(hashmap.find(complement)->second);
                return out;
            }
            hashmap.insert(pair<int,int>(nums[i],i));
        }
        return out; //for all path return value
    }
};
// @lc code=end

