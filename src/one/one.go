package one 
 import (
 	"fmt"
 	"errors"
 )

 func One(){
 	fmt.Println("Welcome to Package One")
	
}                                                                                                                                                                           

func Divide(p1,p2 float64)(ans float64 ,err error){
 	if p2 == 0 {
 		err = errors.New("Cannot divide by 0 ")	
 	}
 	ans = p1/p2
 	return
// 	// this will return ans 
}