type MinStack struct {
    data []int
    min []int
}


func Constructor() MinStack {
    return MinStack{[]int{}, []int{}}
}


func (this *MinStack) Push(val int)  {
    this.data = append(this.data, val)

    if len(this.min) == 0 {
        this.min = append(this.min, val)
    } else {
        currentMin := this.min[len(this.min) - 1]
        if (currentMin > val) {
            this.min = append(this.min, val)
        } else {
            this.min = append(this.min, currentMin)
        }
    }
}


func (this *MinStack) Pop()  {
    if len(this.data) == 0 {
        return
    }
    
    this.data = this.data[:len(this.data) - 1]
    this.min = this.min[:len(this.min) - 1]
}


func (this *MinStack) Top() int {
    if len(this.data) == 0 {
        return 0
    }
    return this.data[len(this.data) - 1]
}


func (this *MinStack) GetMin() int {
    if len(this.min) == 0 {
        return 0
    }
    return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */