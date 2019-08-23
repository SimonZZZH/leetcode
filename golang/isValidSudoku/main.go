//判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
//
//    数字 1-9 在每一行只能出现一次。
//    数字 1-9 在每一列只能出现一次。
//    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-sudoku
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

func main() {
}

func isValidSudoku(board [][]byte) bool {
	rowMap := make([]map[byte]byte, 9)
	columnMap := make([]map[byte]byte, 9)
	sonMap := make([]map[byte]byte, 9)
	for x := 0; x < 9; x++ {
		rowMap[x] = make(map[byte]byte)
		columnMap[x] = make(map[byte]byte)
		sonMap[x] = make(map[byte]byte)
	}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if rowMap[row][board[row][column]] == 1 || columnMap[column][board[row][column]] == 1 {
				return false
			} else if board[row][column] != '.' {
				rowMap[row][board[row][column]] = 1
				columnMap[column][board[row][column]] = 1
			}
			if row < 3 && column < 3 {
				if sonMap[0][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[0][board[row][column]] = 1
				}
			} else if row < 3 && column < 6 {
				if sonMap[1][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[1][board[row][column]] = 1
				}
			} else if row < 3 && column < 9 {
				if sonMap[2][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[2][board[row][column]] = 1
				}
			} else if row < 6 && column < 3 {
				if sonMap[3][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[3][board[row][column]] = 1
				}
			} else if row < 6 && column < 6 {
				if sonMap[4][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[4][board[row][column]] = 1
				}
			} else if row < 6 && column < 9 {
				if sonMap[5][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[5][board[row][column]] = 1
				}
			} else if row < 9 && column < 3 {
				if sonMap[6][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[6][board[row][column]] = 1
				}
			} else if row < 9 && column < 6 {
				if sonMap[7][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[7][board[row][column]] = 1
				}
			} else if row < 9 && column < 9 {
				if sonMap[8][board[row][column]] == 1 {
					return false
				} else if board[row][column] != '.' {
					sonMap[8][board[row][column]] = 1
				}
			}
		}
	}
	return true
}

//func isValidSudoku(board [][]byte) bool {
//	//增加两个数组保存转换后需要判断的内容
//	trans := make([][]byte, 9)
//	for x := 0; x < 9; x++ {
//		tmp := make([]byte, 9)
//		trans = append(trans, tmp)
//	}
//	trans2 := make([][]byte, 9)
//	for x := 0; x < 9; x++ {
//		tmp := make([]byte, 9)
//		trans2 = append(trans2, tmp)
//	}
//	for x := 0; x < 9; x++ {
//		if containsDuplicate(board[x]) {
//			return false
//		}
//		for y := 0; y < 9; y++ {
//			trans[y][x] = board[x][y]
//			if x < 3 && y < 3 {
//
//			} else if x < 6 && y < 6 {
//
//			} else {
//
//			}
//		}
//	}
//	for x := 0; x < 9; x++ {
//		if containsDuplicate(trans[x]) {
//			return false
//		}
//	}
//	//for x := 0; x < len(board); x++ {
//	//
//	//}
//}
//
////有重复返回true
//func containsDuplicate(nums []byte) bool {
//	if len(nums) < 2 {
//		return false
//	}
//	dataMap := make(map[byte]byte)
//	for x := 0; x < len(nums); x++ {
//		if dataMap[nums[x]] == 1 {
//			return true
//		} else if nums[x] != '.' {
//			dataMap[nums[x]] = 1
//		}
//	}
//	return false
//}
