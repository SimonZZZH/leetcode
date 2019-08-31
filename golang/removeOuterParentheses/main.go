package main

func main() {

}

//有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
//
//如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
//
//给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
//
//对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-outermost-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//用栈获取每一段最小的原语
//然后拆分
func removeOuterParentheses(S string) string {
	var res string
	//原语数组
	primi := make([]string, 0)
	//左右指针
	index := 1
	//stack := make([]byte, 0)
	stack := 0
	for x := 0; x < len(S); x++ {
		switch S[x] {
		case '(':
			stack++
		case ')':
			stack--
		}
		if stack == 0 {
			//合法的原语产生
			tmp := S[index:x]
			primi = append(primi, tmp)
			index = x + 2
		}
	}
	//拼接
	for x := 0; x < len(primi); x++ {
		res += primi[x]
	}
	return res
}
