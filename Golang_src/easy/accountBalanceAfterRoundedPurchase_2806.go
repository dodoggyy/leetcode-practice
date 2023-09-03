package easy

// Use math:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.92 MB, less than 97.10%
func accountBalanceAfterPurchase(purchaseAmount int) int {
	amount := 100
	mod := 10
	if purchaseAmount%mod >= mod/2 {
		purchaseAmount += mod - purchaseAmount%mod
	} else {
		purchaseAmount -= purchaseAmount % 10
	}

	return amount - purchaseAmount
}
