package main
import "fmt"

func insert(orig[] int , index int , value int )([] int ){
    fmt.Println(orig  ,"len", len(orig))
    //orig=append(orig[  : index+1]   , orig[index :  ]  ...)
    orig=append(orig[ : len(orig)]   , orig[len(orig)-1 : ]  ...)
    fmt.Println(orig  ,"len", len(orig))
    orig[index]=value
    return orig
}
func main() {
v :=make([]int , 0,10)
fmt.Println(v ,len(v) , cap(v))
for i:=1;i<=10;i++{
    v=append(v,i)
   // fmt.Println(v ,len(v) , cap(v))
}
fmt.Println(v)



fmt.Println( insert(v, 4,-11) )
}
