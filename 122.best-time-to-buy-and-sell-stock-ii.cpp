/*
 * @lc app=leetcode id=122 lang=cpp
 *
 * [122] Best Time to Buy and Sell Stock II
 */
#include <bits/stdc++.h>
using namespace std;
// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int out = 0;
        for (size_t i = 1; i < prices.size(); i++)
        {
            int change = prices[i]-prices[i-1];
            if(change>0){
                out += change;
            }
        }
        return out;
    }
};
// @lc code=end

