//不使用任何内建的哈希表库设计一个哈希映射
//
//具体地说，你的设计应该包含以下的功能
//
//    put(key, value)：向哈希映射中插入(键,值)的数值对。如果键对应的值已经存在，更新这个值。
//    get(key)：返回给定的键所对应的值，如果映射中不包含这个键，返回-1。
//    remove(key)：如果映射中存在这个键，删除这个数值对。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/design-hashmap
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//注意：
//
//    所有的值都在 [1, 1000000]的范围内。
//    操作的总数目在[1, 10000]范围内。
//    不要使用内建的哈希库。
package main

func main() {

}

//用数组随便写写好了。。。
type MyHashMap struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	var hm MyHashMap
	hm.data = make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		hm.data[i] = -1
	}
	return hm
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.data[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.data[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
