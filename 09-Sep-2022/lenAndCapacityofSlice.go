package main
import "fmt"
func main() {
v :=make([]int , 0,10)
fmt.Println(v ,len(v) , cap(v))
for i:=1;i<=10;i++{
    v=append(v,i)
   // fmt.Println(v ,len(v) , cap(v))
}
fmt.Println(v)
v=append(v[:3], append(v,11),v[3:]...)
fmt.Println(v)
}