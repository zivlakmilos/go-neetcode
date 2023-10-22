package stack

type MinStack struct {
	top *MinStackNode
}

type MinStackNode struct {
	Prev  *MinStackNode
	Value int
	Min   int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := &MinStackNode{
		Value: val,
		Min:   val,
	}

	if this.top != nil {
		node.Prev = this.top

		if this.top.Min < node.Min {
			node.Min = this.top.Min
		}
	}

	this.top = node
}

func (this *MinStack) Pop() {
	if this.top != nil {
		this.top = this.top.Prev
	}
}

func (this *MinStack) Top() int {
	if this.top != nil {
		return this.top.Value
	}

	return 0
}

func (this *MinStack) GetMin() int {
	if this.top != nil {
		return this.top.Min
	}

	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
