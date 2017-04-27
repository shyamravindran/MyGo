package main

import ( 
	"fmt"
)


func main(){

        respo := map[string]map[string]map[string]string{"a" :{"b":{"c":"d"}}}
	fmt.Println(respo["a"]["b"]["c"])
	
}
