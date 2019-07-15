/**
 * 
 */

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MoveZeroes_283 {
    // 2019年7月15日
    // Use Index to record non-zero
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.1 MB, less than 87.12%
    public void moveZeroes(int[] nums) {
        if(nums == null) {
            return;
        }
        
        int mLength = nums.length;
        int mIndexNonZero = 0;

        for (int i = 0; i < mLength; i++) {
            if (nums[i] != 0) {
                swap(nums, i, mIndexNonZero++);
            }
        }
    }

    private void swap(int[] nums, int mIndexNum1, int mIndexNum2) {
        if (mIndexNum1 == mIndexNum2) {
            return;
        }

        nums[mIndexNum1] ^= nums[mIndexNum2];
        nums[mIndexNum2] ^= nums[mIndexNum1];
        nums[mIndexNum1] ^= nums[mIndexNum2];
    }
}
