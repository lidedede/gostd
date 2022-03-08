package main

import "fmt"

// ListNode 链表的节点
type ListNode struct {
	val  int
	next *ListNode
}

// is_empty判断一个链表是否为空
func (l *ListNode) is_empty() bool {
	return l.next == nil
}

// link_len 检查链表的长度
func (l *ListNode) link_len() int {
	count := 0
	cur := l.next
	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}

// append 从尾部插入节点
func (l *ListNode) append(elm int) {
	node := new(ListNode)
	node.val = elm
	if l.is_empty() {
		l.next = node
	} else {
		cur := l.next
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
}

// add 从头部插入节点
func (l *ListNode) add(elm int) {
	node := new(ListNode)
	node.val = elm
	if l.is_empty() {
		l.next = node
	} else {
		node.next = l.next
		l.next = node
	}
}

// Insert 从指定位置插入节点
func (l *ListNode) Insert(pos int, elm int) {
	node := new(ListNode)
	node.val = elm

	if l.is_empty() {
		l.next = node
	} else if pos >= l.link_len() {
		l.append(elm)
	} else if pos <= 0 {
		l.add(elm)
	} else {
		count := 1
		cur := l.next
		for cur != nil {
			if pos == count {
				node.next = cur.next
				cur.next = node
				break
			} else {
				count++
				cur = cur.next
			}
		}
	}
}

// delete 删除链表中的元素
func (l *ListNode) delete(pos int) {
	if l.is_empty() {
		fmt.Println("链表为空！！！")
		return
	} else if l.link_len() == 1 {
		l.next = nil
	} else if pos <= 0 {
		l.next = l.next.next
	} else if pos >= l.link_len() {
		pre := l.next
		cur := l.next
		for cur.next != nil {
			pre = cur
			cur = cur.next
		}
		pre.next = nil
	} else {
		pre := l.next
		cur := l.next
		count := 0
		for cur.next != nil {
			if pos == count {
				pre.next = cur.next
				break
			}
			count++
			pre = cur
			cur = cur.next
		}
	}
}

// Reverse 链表逆序
func (l *ListNode) Reverse() {
	if l.is_empty() {
		fmt.Println("链表为空！！！")
		return
	} else {
		cur := l.next
		next := cur.next
		pre := new(ListNode)

		for cur.next != nil {
			cur.next = pre.next
			pre.next = cur
			cur = next
			next = cur.next
		}
		cur.next = pre.next
		l.next = cur
	}
}

// RemoveTheSame 从无序链表中移除重复项
func (l *ListNode) RemoveTheSame() {
	if l.is_empty() {
		fmt.Println("链表为空！！！")
		return
	} else {
		cur := l.next
		for cur != nil {
			next := cur.next
			pre := cur
			for next != nil {
				if next.val == cur.val {
					pre.next = next.next
					next = next.next
				} else {
					pre = next
					next = next.next
				}
			}
			cur = cur.next
		}
	}
}

// Traverse 遍历链表
func (l *ListNode) Traverse() {
	cur := l.next
	for cur != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
	fmt.Printf("\n")
}

func main() {
	link := new(ListNode)
	fmt.Println("the link is empty?", link.is_empty())

	fmt.Println("从尾部插入节点--------")
	link.append(1)
	link.append(2)
	link.append(3)
	link.Traverse()
	fmt.Println("--------从尾部插入节点")
}