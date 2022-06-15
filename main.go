package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"mwt.com/grpc/demo/service"
)

func main() {
	a := new(string)
	user := &service.User{
		Username:  "老朱",
		Age:       20,
		Height:    18,
		StuNumber: 1608114121,
		Password:  a,
		Addresses: []string{"郑州", "上海", "合肥"},
	}
	//转换为protobuf
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.Addresses[0])
	fmt.Println(newUser.String())
}
