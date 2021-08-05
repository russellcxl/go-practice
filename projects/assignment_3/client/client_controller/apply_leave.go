package client_controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/projects/assignment_3/client/util"
	pb "github.com/russellcxl/go-practice/projects/assignment_3/protos"
	tcp "github.com/russellcxl/go-practice/projects/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"net"
	"time"
)

func ApplyLeave(conn net.Conn, user *pb.User, startDate, endDate string) (isSuccess bool, errMsg string) {
	// check input format is correct
	startFormat := util.CheckDateFormat(startDate)
	if startFormat == false {
		return false, "Incorrect start date format"
	}
	endFormat := util.CheckDateFormat(endDate)
	if endFormat == false {
		return false, "Incorrect end date format"
	}

	// check now < start date < end date
	startCorrect := util.CheckStartDate(startDate)
	if startCorrect == false {
		return false, "Start date should be later than now"
	}
	endCorrect := util.CheckEndDate(startDate, endDate)
	if endCorrect == false {
		return false, "End date should be later than start date"
	}

	// check if user has enough leaves. This will also be done on server side
	hasEnough, daysTaken := util.CheckSufficientLeavesForCurrentUser(startDate, endDate, int(user.GetLeaveBalance()))
	if hasEnough == false {
		return false, "You don't have enough leave days left"
	}

	// create leave request
	s, _ := time.Parse("02/01/06", startDate)
	e, _ := time.Parse("02/01/06", endDate)
	sUnix := uint32(s.Unix())
	eUnix := uint32(e.Unix())

	leave := &pb.Leave{
		UserId:     user.GetUserId(),
		TeamId:     user.GetTeamId(),
		StartTime:  sUnix,
		EndTime:    eUnix,
		DaysTaken:  uint32(daysTaken),
		Status:     int32(pb.LeaveStatus_LEAVE_STATUS_PENDING),
	}

	// send and receive message from server
	req, _ := proto.Marshal(leave)
	var res *pb.Message

	tcp.SendMessage(conn, pb.Command_SET_LEAVE_APPLICATION, req, pb.Error_SUCCESS)

	res = tcp.ReadMessage(conn)
	if res.GetError() != pb.Error_SUCCESS {
		return false, fmt.Sprintf("[Server error] Failed to apply leave, error code: %v\n", res.GetError())
	}

	var updatedUser pb.User
	err := proto.Unmarshal(res.GetData(), &updatedUser)
	if err != nil {
		return false, fmt.Sprintf("Failed to unmarshal data into User struct: %v\n", err)
	}

	// update locally cached user data
	user.LeaveBalance = updatedUser.GetLeaveBalance()

	return true, ""
}
