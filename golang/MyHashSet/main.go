//不使用任何内建的哈希表库设计一个哈希集合
//
//具体地说，你的设计应该包含以下的功能
//
//    add(value)：向哈希集合中插入一个值。
//    contains(value) ：返回哈希集合中是否存在这个值。
//    remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/design-hashset
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

func main() {

}

type MyHashSet struct {
	set map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	var set MyHashSet
	set.set = make(map[int]int)
	return set
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = key
}

func (this *MyHashSet) Remove(key int) {
	delete(this.set, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.set[key]; ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
