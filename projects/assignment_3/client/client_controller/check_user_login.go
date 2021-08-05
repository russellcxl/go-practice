package client_controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/projects/assignment_3/client/util"
	pb "github.com/russellcxl/go-practice/assignment_3/protos"
	tcp "github.com/russellcxl/go-practice/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strconv"
	"strings"
)

func CheckUserLoginDetails(conn net.Conn, credentials string) (*pb.User, bool) {

	// check formatting
	if util.ValidateLoginFormat(credentials) == false {
		fmt.Printf("> Wrong format. Please follow [user id] [password]\n")
		return nil, false
	}

	c := strings.Split(credentials, " ")
	userId, _ := strconv.Atoi(c[0])
	password := c[1]

	reqBytes, err := proto.Marshal(&pb.User{UserId: uint64(userId), Password: password})
	if err != nil {
		log.Println("Failed to marshal credentials: ", err)
	}

	tcp.SendMessage(conn, pb.Command_CHECK_USER_LOGIN, reqBytes, pb.Error_SUCCESS)

	// get response from server
	msg := tcp.ReadMessage(conn)
	if msg.Error != pb.Error_SUCCESS {
		return nil, false
	}

	var user pb.User
	err = proto.Unmarshal(msg.Data, &user)
	if err != nil {
		log.Printf("Failed to unmarshal into Login struct: %v\n", err)
	}

	return &user, true

}
