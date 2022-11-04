package body

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type reqbody struct {
	Name string
	Age  int
}

// 解析request中body的构成
func body(rsp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var bodySlc []byte = make([]byte, 1024)

	bodyLen, readErr := req.Body.Read(bodySlc)
	if readErr != nil {
		fmt.Println("read error")
	} else {
		fmt.Println("body has", bodyLen, " bytes")
	}

	fmt.Println("body: ", string(bodySlc))
	fmt.Println("length of body: ", len(bodySlc))

	str := string(bodySlc)
	fmt.Println("length of str: ", len(str))
	var body2 reqbody
	err := json.Unmarshal([]byte(str), &body2)
	if err != nil {
		fmt.Println("unmarshal error")
	} else {
		fmt.Println("name : " + body2.Name)
		fmt.Printf("age : %v\n", body2.Age)
	}
}
