package main

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
 */

func searchMatrixByBinary(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	column := len(matrix[0])
	left := 0
	right := row*column - 1

	for left <= right {
		mid := (left + right) / 2
		num := matrix[mid/column][mid%column]
		if num == target {
			return true
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	row, column := rows-1, 0

	for row >= 0 && column <= columns-1 {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			row--
		} else {
			column++
		}
	}
	return false
}
