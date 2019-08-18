package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"1"}))
}

//你现在是棒球比赛记录员。
//给定一个字符串列表，每个字符串可以是以下四种类型之一：
//1.整数（一轮的得分）：直接表示您在本轮中获得的积分数。
//2. "+"（一轮的得分）：表示本轮获得的得分是前两轮有效 回合得分的总和。
//3. "D"（一轮的得分）：表示本轮获得的得分是前一轮有效 回合得分的两倍。
//4. "C"（一个操作，这不是一个回合的分数）：表示您获得的最后一个有效 回合的分数是无效的，应该被移除。
//
//每一轮的操作都是永久性的，可能会对前一轮和后一轮产生影响。
//你需要返回你在所有回合中得分的总和。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/baseball-game
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func calPoints(ops []string) int {
	var res int
	score := make([]int, 0)
	//score = []int{1, 2, 3, 4, 5, 7}
	//score = score[:len(score)-1]
	//fmt.Println(score)
	//return 0
	for x := 0; x < len(ops); x++ {
		switch ops[x] {
		case "+":
			//前两轮和
			switch len(score) {
			case 0:
				break
			case 1:
				score = append(score, score[0])
			default:
				score = append(score, score[len(score)-1]+score[len(score)-2])
			}
		case "D":
			//前一轮两倍
			tmp := score[len(score)-1] * 2
			score = append(score, tmp)
		case "C":
			//清除上一轮得分
			score = score[:len(score)-1]
		default:
			//数字
			tmp, _ := strconv.Atoi(ops[x])
			score = append(score, tmp)
		}
	}
	//求和
	for x := 0; x < len(score); x++ {
		res += score[x]
	}
	return res
}
