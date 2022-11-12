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
    // V3, start from two ends (beginWord and endWord)
    // TODO: I think it can be even faster with Go's goroutine
    int ladderLength_V3(string beginWord, string endWord,
                        vector<string> &wordList) {
        unordered_set<string> notVisited(wordList.begin(), wordList.end());
        queue<string> q({beginWord});

        int depth = 0;
        while (!q.empty()) {
            depth++;
            // printf("level %d\n", level);
            auto len = q.size();
            for (int i = 0; i < len; i++) {
                auto cur = q.front();
                q.pop();
                // printf("pop %d\n", curP);
                if (cur == endWord) {
                    return depth;
                }
                notVisited.erase(cur);  // dont search for self
                auto candidates = getCandidates(cur);
                for (auto can : *candidates) {
                    if (notVisited.find(can) != notVisited.end()) {
                        q.push(can);
                        notVisited.erase(can);  // !! key optimize
                    }
                }
                delete candidates;
            }
        }
        return 0;
    }

    // v2, AC but only beat 30%
    // consider that the len<=10, but the wordList.size() is large (<=5000)
    // I'm going to create a condidaate list for each word (learned from cookbook)
    // so that max iteration for each pop is 26*10=260 rather than 5000
    int ladderLength_V2(string beginWord, string endWord,
                        vector<string> &wordList) {
        unordered_set<string> notVisited(wordList.begin(), wordList.end());
        queue<string> q({beginWord});

        int depth = 0;
        while (!q.empty()) {
            depth++;
            // printf("level %d\n", level);
            auto len = q.size();
            for (int i = 0; i < len; i++) {
                auto cur = q.front();
                q.pop();
                // printf("pop %d\n", curP);
                if (cur == endWord) {
                    return depth;
                }
                notVisited.erase(cur);  // dont search for self
                auto candidates = getCandidates(cur);
                for (auto can : *candidates) {
                    if (notVisited.find(can) != notVisited.end()) {
                        q.push(can);
                        notVisited.erase(can);  // !! key optimize
                            // will pop eventually, so dont need to find it anymore
                    }
                }
                delete candidates;
            }
        }
        return 0;
    }

    // get all words that have only one letter differing from s
    vector<string> *getCandidates(string &s) {
        auto ret = new vector<string>();
        for (uint8_t i = 0; i < s.size(); i++) {
            for (char j = 'a'; j <= 'z'; j++) {
                if (s[i] != j) {
                    string can = s;
                    can[i] = j;
                    ret->push_back(can);
                }
            }
        }
        return ret;
    }

    // v1, go through the large list for each pop, thus too much time - beats 10%
    // but mem usage beats 100% (for no storage for candidates)
    int ladderLength_V1(string beginWord, string endWord,
                        vector<string> &wordList) {
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
            // printf("level %d\n", level);
            auto len = q.size();
            for (int i = 0; i < len; i++) {
                auto curIndex = q.front();
                q.pop();
                // printf("pop %d\n", curIndex);
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
                        // printf("push %d\n", j);
                        visited[j] = true;  // key optimize
                    }
                }
                visited[curIndex] = true;
            }
        }
        return 0;
    }

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
};

// to learn - seems to be like V3
// copilot beats both 100%
class Solution_copilot {
   public:
    int ladderLength(string beginWord, string endWord,
                     vector<string> &wordList) {
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
