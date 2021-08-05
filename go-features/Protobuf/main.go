package main

import (
	"fmt"
	"git.garena.com/russell.chanxl/personal/Protobuf/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := pb.Person {
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber {
			{Number: "555-4321", Type: pb.Person_HOME}, // pb.Person_HOME = 1
		},
	}


	bs, _ := proto.Marshal(&p)
	fmt.Printf("\nDefault: %v\n\nBinary: %b\n\nHex: %x\n\n", bs, bs, bs)

	pUnmarshalled := pb.Person{}

	err := proto.Unmarshal(bs, &pUnmarshalled)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", &pUnmarshalled)

}
