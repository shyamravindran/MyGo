package main

import (
    "encoding/json"
    "fmt"
)

type Data struct {
    Name string
    Id   int
    Json *json.RawMessage
}
type Data2 struct {
    Name string
    Id   int
}


func main() {

    tmp := Data2{"World", 2}
    dict1, _:= json.Marshal(tmp) //string(dict1) will be {"Name":"World","Id":2}
    var j_dict1 *json.RawMessage
    json.Unmarshal(dict1,&j_dict1)
    
    tmp2 := Data{"World2",4,j_dict1}
    dict2, _ := json.Marshal(tmp2)
    fmt.Println("dict2 is ", string(dict2))
    
    var final Data
    json.Unmarshal(dict2,&final)
    fmt.Println("Finaly:",final)
    

}
