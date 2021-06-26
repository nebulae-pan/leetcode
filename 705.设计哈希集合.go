/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

// @lc code=start
type MyHashSet struct {
	containers []bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]bool, 10000001)}
}


func (this *MyHashSet) Add(key int)  {
	this.containers[key] = true
}


func (this *MyHashSet) Remove(key int)  {
	this.containers[key] = false
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.containers[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

