#include <vector>
using namespace std;
class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        nums1.resize(m + n);
        std::copy(nums2.begin(), nums2.begin() + n, nums1.begin() + m);
        std::sort(nums1.begin(), nums1.end());  
    }
};