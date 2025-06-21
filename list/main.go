package main

import (
    "fmt"
)

// ListNode 链表节点数据结构定义
type ListNode struct {
    prev, next *ListNode // 前驱节点、后继节点
    val        any       // 数据
}

// List 双向链表数据结构定义
type List struct {
    head, tail *ListNode // 头节点、尾节点
    len        int       // 链表长度
}

// Next ListNode Next 获取下一个节点，如果不存在，则返回nil
func (l *ListNode) Next() *ListNode {
    if l != nil && l.next != nil {
        return l.next
    }
    return nil
}

// Prev ListNode Prev 获取前一个节点，如果不存在，则返回nil
func (l *ListNode) Prev() *ListNode {
    if l != nil && l.prev != nil {
        return l.prev
    }
    return nil
}

// NewList List 创建一个链表
// 使用虚拟结点写法，创建虚拟头结点、尾结点，避免后续各类边界判断
func NewList() *List {
    dummyHead := &ListNode{val: nil}
    dummyTail := &ListNode{val: nil}

    dummyHead.next = dummyTail
    dummyTail.prev = dummyHead

    list := &List{
        head: dummyHead,
        tail: dummyTail,
        len:  0,
    }
    return list
}

// PushBack 尾插
func (l *List) PushBack(v any) {
    node := &ListNode{
        val: v,
    }
    node.next = l.tail
    node.prev = l.tail.prev
    l.tail.prev.next = node
    l.tail.prev = node
}

// PushFront 头插
func (l *List) PushFront(v any) {
    node := &ListNode{
        val: v,
    }

    node.prev = l.head
    node.next = l.head.next
    l.head.next.prev = node
    l.head.next = node
    l.len++
}

// InsertAfter 在 at 节点后面插入节点
// at 不存在，则直接插入末尾
func (l *List) InsertAfter(at *ListNode, v any) {
    if at == nil {
        l.PushBack(v)
        return
    }

    node := &ListNode{
        val: v,
    }

    node.prev = at
    node.next = at.next
    node.prev.next = node
    node.next.prev = node
    l.len++
}

// InsertBefore 在 at 节点前面插入节点
// at 不存在则直接插入开头
func (l *List) InsertBefore(at *ListNode, v any) {
    if at == nil {
        l.PushFront(v)
        return
    }

    node := &ListNode{
        val: v,
    }

    node.prev = at
    node.next = at.next
    node.prev.next = node
    node.next.prev = node
    l.len++
}

// Find 查找节点
// 仅返回第一个匹配的节点
func (l *List) Find(v any) *ListNode {
    // 跳过虚拟头结点、跳过虚拟尾结点
    for node := l.head.next; node.Next() != nil; node = node.next {
        if node.val == v {
            return node
        }
    }
    return nil
}

// Remove 删除节点
func (l *List) Remove(v any) {
    node := l.Find(v)
    if node != nil {
        node.prev.next = node.next
        node.next.prev = node.prev
        node.next = nil
        node.prev = nil
        l.len--
    }
}

// GetLen 获取长度
func (l *List) GetLen() int {
    return l.len
}

// ForEach 正向遍历
func (l *List) ForEach() []any {
    res := make([]any, 0)
    for node := l.head.next; node.Next() != nil; node = node.next {
        res = append(res, node.val)
    }
    return res
}

// ForReverse 反向遍历
func (l *List) ForReverse() []any {
    res := make([]any, 0)
    for node := l.tail.prev; node.Prev() != nil; node = node.prev {
        res = append(res, node.val)
    }
    return res
}

// Print 打印
func (l *List) Print() {
    res := make([]any, 0)
    // 跳过虚拟头结点、跳过虚拟尾结点
    for node := l.head.next; node.Next() != nil; node = node.next {
        res = append(res, node.val)
    }
    fmt.Println(res)
}

func main() {
    l := NewList()

    // 尾插
    l.PushBack(4)
    l.PushBack(5)
    l.PushBack(6)
    l.PushBack(8)
    l.Print() // [4 5 6 8]

    // 头插
    l.PushFront(2)
    l.PushFront(1)
    l.PushFront(0)
    l.Print() // [0 1 2 4 5 6 8]

    // 在 at 后插入
    at := l.Find(2)
    l.InsertAfter(at, 3)
    l.Print() // [0 1 2 3 4 5 6 8]

    // 在 at 前插入
    at = l.Find(6)
    l.InsertBefore(at, 7)
    l.Print() // [0 1 2 3 4 5 6 7 8]

    // 删除头
    l.Remove(0)
    l.Print() // [1 2 3 4 5 6 7 8]

    // 删除尾
    l.Remove(8)
    l.Print() // [1 2 3 4 5 6 7]

    // 删除中间节点
    l.Remove(4)
    l.Print() // [1 2 3 5 6 7]

    fmt.Println(l.ForEach())    // [1 2 3 5 6 7]
    fmt.Println(l.ForReverse()) // [7 6 5 3 2 1]
}
