package client_controller

import (
	"fmt"
	pb "git.garena.com/russell.chanxl/be-class/assignment_3/protos"
	tcp "git.garena.com/russell.chanxl/be-class/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"net"
	"strconv"
)

func DeleteLeaveApplication(conn net.Conn, user *pb.User, input string) {

	leaveId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		fmt.Println("\n**************************************\n ")
		return
	}

	leave := pb.Leave{LeaveId: uint64(leaveId), UserId: user.GetUserId()}
	req, _ := proto.Marshal(&leave)
	tcp.SendMessage(conn, pb.Command_DELETE_LEAVE_APPLICATION, req, pb.Error_SUCCESS)

	res := tcp.ReadMessage(conn)
	if res.GetError() != 0 {
		fmt.Printf("[Server error] Failed to delete leave application, error code: %v\n", res.GetError())
		fmt.Println("\n**************************************\n ")
		return
	}

	var updatedUser pb.User
	err = proto.Unmarshal(res.GetData(), &updatedUser)
	if err != nil {
		fmt.Printf("Failed to unmarshal data into User struct: %v\n", err)
		fmt.Println("\n**************************************\n ")
		return
	}

	// update locally cached user data
	user.LeaveBalance = updatedUser.GetLeaveBalance()


	fmt.Printf("Successfully deleted leave application! You have %d days left\n", user.GetLeaveBalance())
	fmt.Println("\n**************************************\n ")
}
