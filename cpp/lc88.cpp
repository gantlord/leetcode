#include <vector>
using namespace std;
class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int p1=0, p2=0;
        vector<int> nums3;
        while(p1<m && p2<n){
            if(nums1[p1]<nums2[p2]){
                nums3.push_back(nums1[p1]);
                p1++;
            }
            else{
                nums3.push_back(nums2[p2]);
                p2++;
            }
        }
        while(p1<m){
            nums3.push_back(nums1[p1]);
            p1++;
        }
        while(p2<n){
            nums3.push_back(nums2[p2]);
            p2++;
        }
        nums1 = nums3;    
    }
};