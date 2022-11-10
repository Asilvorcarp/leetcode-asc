/*
 * @lc app=leetcode id=121 lang=cpp
 *
 * [121] Best Time to Buy and Sell Stock
 */
#include <bits/stdc++.h>
using namespace std;
// @lc code=start
class Solution {
public: 
    #define MAX 1e4+1
    int maxProfit(vector<int>& prices) {
        // destroys the prices vector to get the flow
        // for (size_t i = prices.size()-1; i > 0; i++)
        // {
        //     prices[i] = prices[i]-prices[i-1];
        // }
        // prices[0]=0;
        int min = MAX, out = 0;
        for (size_t i = 0; i < prices.size(); i++)
        {
            if(prices[i] < min){
                min = prices[i];
            }
            int curProfit = prices[i] - min;
            if(curProfit > out){
                out = curProfit;
            }
        }
        return out;        
    }
};
// @lc code=end

