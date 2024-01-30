import { merge } from './lc88';

describe('merge', () => {
  it('should merge two sorted arrays into the first array', () => {
    const nums1 = [1, 2, 3, 0, 0, 0];
    const m = 3;
    const nums2 = [2, 5, 6];
    const n = 3;

    merge(nums1, m, nums2, n);

    expect(nums1).toEqual([1, 2, 2, 3, 5, 6]);
  });

  it('should merge two sorted arrays with empty nums2', () => {
    const nums1 = [1, 2, 3, 4, 5];
    const m = 5;
    const nums2: number[] = [];
    const n = 0;

    merge(nums1, m, nums2, n);

    expect(nums1).toEqual([1, 2, 3, 4, 5]);
  });

  it('should merge two sorted arrays with empty nums1', () => {
    const nums1: number[] = [];
    const m = 0;
    const nums2 = [1, 2, 3, 4, 5];
    const n = 5;

    merge(nums1, m, nums2, n);

    expect(nums1).toEqual([1, 2, 3, 4, 5]);
  });

  it('should merge two sorted arrays with duplicate elements', () => {
    const nums1 = [1, 2, 3, 3, 4, 5];
    const m = 6;
    const nums2 = [2, 3, 4, 5, 6];
    const n = 5;

    merge(nums1, m, nums2, n);

    expect(nums1).toEqual([1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 6]);
  });
});
