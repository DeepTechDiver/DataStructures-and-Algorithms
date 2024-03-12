package main

// 循环链表 next 表示后驱节点，prev 表示前驱节点，Value 表示值

type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}
func main() {
	r := new(Ring)
	r.init()
}
