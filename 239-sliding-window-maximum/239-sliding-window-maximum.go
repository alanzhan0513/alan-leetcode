type MyQueue struct{
    queue []int
}

func NewMyQueue() *MyQueue {
    return &MyQueue{
        queue: make([]int, 0),
    }
}

func (m *MyQueue) Push(k int){
    if len(m.queue)>0 && k>m.queue[0]{
        m.queue=[]int{}
    }
    for len(m.queue)>0 && m.queue[len(m.queue)-1]<k{
        m.queue=m.queue[:len(m.queue)-1]
    }
    m.queue=append(m.queue,k)
}

func (m *MyQueue) Pop(t int){  
    if m.queue[0]==t{
        m.queue=m.queue[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
   queue:=NewMyQueue()
    arr:=make([]int,0)
    left,right:=0,k-1

    for i:=left;i<=right;i++{
        queue.Push(nums[i])
    }
    arr=append(arr,queue.queue[0])

    right++ 
    for right<=len(nums)-1{
         queue.Pop(nums[left])
        queue.Push(nums[right])
        arr=append(arr,queue.queue[0])
        left++
        right++
    }

    return arr

}