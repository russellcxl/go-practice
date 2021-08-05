package client_controller

import (
	"fmt"
	pb "git.garena.com/russell.chanxl/be-class/assignment_3/protos"
	tcp "git.garena.com/russell.chanxl/be-class/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"net"
)

func GetLeaveApplications(conn net.Conn, userId uint64) ([]*pb.Leave, string) {
	var user pb.User
	user.UserId = userId
	req, _ := proto.Marshal(&user)
	tcp.SendMessage(conn, pb.Command_GET_LEAVE_APPLICATIONS_BY_USER_ID, req, pb.Error_SUCCESS)

	msg := tcp.ReadMessage(conn)
	if msg.Error != pb.Error_SUCCESS {
		return nil, fmt.Sprintf("> [Server error] failed to get leave applications, error code: %v\n", msg.GetError())
	}
	var leaves pb.Leaves
	err := proto.Unmarshal(msg.GetData(), &leaves)
	if err != nil {
		return nil, fmt.Sprintf("> Failed to unmarshal data into User struct: %v\n", err)
	}

	return leaves.GetLeaves(), ""
}
