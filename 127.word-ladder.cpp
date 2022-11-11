/*
 * @lc app=leetcode.cn id=127 lang=cpp
 *
 * [127] Word Ladder
 */

#include <bits/stdc++.h>
using namespace std;

// @lc code=start
class Solution {
public:
    bool isOneDiff(string &a, string &b) {
        // assume a.size() == b.size()
        uint8_t diff = 0;
        for (int i = 0; i < a.size(); i++) {
            if (a[i] != b[i]) {
                diff++;
                if (diff > 1) {
                    return false;
                }
            }
        }
        if (diff == 1) {
            return true;
        } else {
            // diff == 0
            return false;
        }
    }

    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        vector<bool> visited(wordList.size() + 1, false);
        // TODO not sure: maybe need to remove the beginWord first
        wordList.insert(wordList.begin(), beginWord);
        queue<int> q({0});

        auto endIndex = find(wordList.begin(), wordList.end(), endWord);
        if (endIndex == wordList.end()) {
            // if endWord not in wordList, return 0
            return 0;
        } else {
            // else remove the endWord from wordList
            wordList.erase(endIndex);
        }

        int level = 0;
        while (!q.empty()) {
            level++;
            printf("level %d\n", level);
            auto len = q.size();
            for (int i = 0; i < len; i++) {
                auto curIndex = q.front();
                q.pop();
                printf("pop %d\n", curIndex);
                string cur = wordList[curIndex];
                if (isOneDiff(cur, endWord)) {
                    return level + 1;
                }
                // j is nextIndex
                for (int j = 0; j < wordList.size(); j++) {
                    if (j == curIndex) {
                        continue;
                    }
                    if (visited[j]) continue;
                    auto next = wordList[j];
                    if (isOneDiff(cur, next)) {
                        q.push(j);
                        printf("push %d\n", j);
                    }
                }
                visited[curIndex] = true;
            }
        }
        return 0;
    }
};

class Solution_copilot {
public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string> wordSet(wordList.begin(), wordList.end());
        if (wordSet.find(endWord) == wordSet.end()) return 0;
        unordered_set<string> beginSet{beginWord};
        unordered_set<string> endSet{endWord};
        int len = 1;
        int strLen = beginWord.length();
        while (!beginSet.empty() && !endSet.empty()) {
            if (beginSet.size() > endSet.size()) {
                swap(beginSet, endSet);
            }
            unordered_set<string> tempSet;
            for (string word : beginSet) {
                wordSet.erase(word);
                for (int i = 0; i < strLen; i++) {
                    char ch = word[i];
                    for (int j = 'a'; j <= 'z'; j++) {
                        word[i] = j;
                        if (endSet.find(word) != endSet.end()) {
                            return len + 1;
                        }
                        if (wordSet.find(word) != wordSet.end()) {
                            tempSet.insert(word);
                            wordSet.erase(word);
                        }
                    }
                    word[i] = ch;
                }
            }
            swap(beginSet, tempSet);
            len++;
        }
        return 0;
    }
};
// @lc code=end

