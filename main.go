package main
import (
	"fmt"
)

type List []int
func (l List)Len()int{
	return len(l)
}
func (l *List) Append(val int){
	*l=append(*l,val)
}
type AppendIr interface{
	Append(int)
}
func CountInto(a AppendIr,start,end int){
	for i:=start;i<=end;i++{
		a.Append(i)
	}
}

