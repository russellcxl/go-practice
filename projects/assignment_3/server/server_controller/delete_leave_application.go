package server_controller

import (
	pb "github.com/russellcxl/go-practice/assignment_3/protos"
	"github.com/russellcxl/go-practice/assignment_3/server/cache"
	db "github.com/russellcxl/go-practice/assignment_3/server/database"
	tcp "github.com/russellcxl/go-practice/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

func DeleteLeaveApplication(conn net.Conn, r *cache.RedisRepo, data []byte) {
	var leave pb.Leave
	err := proto.Unmarshal(data, &leave)
	if err != nil {
		log.Printf("Failed to unmarshal into Leave: %v\n", err)
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}

	leaveId := leave.GetLeaveId()
	userId := leave.GetUserId()

	// check if leave exists in cache; and if not there, DB
	var userLeaves *pb.Leaves
	userLeaves, err = r.GetLeavesByUserId(userId)
	if err != nil {
		log.Printf("Failed to get from Redis\n")
	} else  {
		log.Printf("Successfully retrieved from Redis\n")
	}
	var leaves []*pb.Leave
	var leaveToBeDeleted *pb.Leave
	if userLeaves != nil {
		leaves = userLeaves.GetLeaves()
	}
	if len(leaves) != 0 {
		for _, l := range leaves {
			if l.GetLeaveId() == leaveId {
				leaveToBeDeleted = l
				break
			}
		}
	}
	if leaveToBeDeleted == nil {
		leaveToBeDeleted, err = db.GetLeaveByLeaveId(leaveId)
		if err != nil {
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_GET_LEAVES)
			return
		}
	}
	if leaveToBeDeleted == nil {
		log.Println("Leave ID not found")
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}

	if err = db.DeleteLeave(leaveId); err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}

	// update user data in DB and cache
	daysRefunded := leaveToBeDeleted.GetDaysTaken()

	user, err := getUserFromCache(r, userId)
	if err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}
	currentLeaveBalance := user.GetLeaveBalance()
	newLeaveBalance := currentLeaveBalance + daysRefunded

	if err = db.SetUserLeaveBalance(userId, newLeaveBalance); err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_USER)
		return
	}

	user.LeaveBalance = newLeaveBalance
	err = r.SetUser(user)
	if err != nil {
		log.Printf("Failed to set in Redis: %v\n", err)
	}

	// update user_leaves cache
	newSetOfLeaves, err := db.GetLeavesByUserId(userId)
	if err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_GET_LEAVES)
		return
	}
	err = r.SetLeavesByUserId(userId, newSetOfLeaves)
	if err != nil {
		log.Printf("Failed to set in Redis: %v\n", err)
	}

	// return user data to client
	user.LeaveBalance = newLeaveBalance
	resByte, _ := proto.Marshal(user)
	tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, resByte, pb.Error_SUCCESS)

}
