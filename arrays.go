package main

import ( 
	"fmt"
)


func main(){

        var arr [3]interface{}
 	arr[0] = "shyam"
	arr[1] = 1
	arr[2] = 1.24
	fmt.Println(arr)
	
	arr1 := []byte{3,3,3,4}
	fmt.Println(arr1)
	
	arr2 := []interface{}{"shyam",26}
	fmt.Println(arr2)
	
	arr3 := make([]int,6)
	fmt.Println(arr3)
	
	fmt.Println(len(arr3))
	fmt.Println(cap(arr2))
	
	//
	s := make([]byte, 3, 4)  //4 is capacity
	fmt.Println(s)
	
}
