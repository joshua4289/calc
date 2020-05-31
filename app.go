package main
import ("fmt"
"github.com/joshua4289/calc/src/one")

func main() {
	fmt.Println("Main context starts")
	one.One()
	
	ans1,_ := one.Divide(2,3)
	fmt.Println(ans1)
}