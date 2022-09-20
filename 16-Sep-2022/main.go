package main

import ( 
    "fmt"
//    "os"
)

type Queue struct{
    Info [] string
}

func New() *Queue{
    return ( &Queue{
                    Info :[]string{},
    })
}

func (q *Queue) Push(ele string)(*Queue) {
    q.Info=append(q.Info, ele)
    return q
}

func (q *Queue) Pop()(string,error) { // B,C 
    ele:=q.Info[0] // ele -> B // taking backup of 1st element for Queue i.e, FIFO 
    q.Info=q.Info[1:] // q.info = {C} -> starts from 1st element i.e, 0th element is popped out 
    return ele , nil // 1st element 
}

/*func (q *Queue) StackPop()(string ,error){ 

     ele:=q.Info[  len (q.Info)-1  ] // taking backup of last element for Stack i.e, LIFO
     q.Info=q.Info[   :   len (q.Info)-1  ] 
     return ele , nil

}*/ 

func main() {
    head :=New()
    head= head.Push("A")
    head= head.Push("B")
    head= head.Push("C")
    //head= head.Push("D")

    // Push -> A,B,C 
    fmt.Println(head.Pop())
    fmt.Println(head.Pop())
    fmt.Println(head.Pop())
    // Queue -> FIFO, Pop Logic -> "A" popped -> B,C

    //fmt.Println(head.StackPop())
    fmt.Println(*head)

}
