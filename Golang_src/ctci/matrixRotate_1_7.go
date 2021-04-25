package ctci

// Use iterative:
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// same as LC.48
func MatrixRotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}

	size := len(matrix)
	for layer := 0; layer < size/2; layer++ {
		last := size - 1 - layer
		for i := layer; i < last; i++ {
			offset := i - layer
			tmp := matrix[layer][i]                                // keep left-top
			matrix[layer][i] = matrix[last-offset][layer]          // left-bot -> left-top
			matrix[last-offset][layer] = matrix[last][last-offset] // right-bot -> left-bot
			matrix[last][last-offset] = matrix[i][last]            // right-top -> right-bot
			matrix[i][last] = tmp                                  // left-top -> right-top
		}
	}
}
