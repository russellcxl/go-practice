package server_controller

import (
	pb "git.garena.com/russell.chanxl/be-class/assignment_3/protos"
	"git.garena.com/russell.chanxl/be-class/assignment_3/server/cache"
	db "git.garena.com/russell.chanxl/be-class/assignment_3/server/database"
	tcp "git.garena.com/russell.chanxl/be-class/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

func SetLeave(conn net.Conn, r *cache.RedisRepo, data []byte) {
	var leave pb.Leave
	err := proto.Unmarshal(data, &leave)
	if err != nil {
		log.Printf("Failed to unmarshal into Leave: %v\n", err)
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}

	leaveId, userId, daysToBeTaken, teamId, startTime, endTime, status := leave.GetLeaveId(), leave.GetUserId(), leave.GetDaysTaken(), leave.GetTeamId(), leave.GetStartTime(), leave.GetEndTime(), leave.GetStatus()

	user, err := getUserFromCache(r, userId)
	if err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}
	if user == nil {
		log.Println("User not found")
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_INVALID_USER_ID)
		return
	}
	currentLeaveBalance := user.GetLeaveBalance()

	if currentLeaveBalance < daysToBeTaken {
		log.Println("User does not have enough leave balance")
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
		return
	}

	// if leave id is not provided, create leave, else update
	switch leaveId {
	case 0:

		log.Printf("%v\n", &leave)

		if userId == 0 || teamId == 0 || startTime == 0 || endTime == 0 || daysToBeTaken == 0 {
			log.Println("Invalid params")
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_INVALID_PARAMS)
			return
		}

		// TODO approverId hardcoded for now
		err = db.CreateLeave(
			userId,
			teamId,
			startTime,
			endTime,
			daysToBeTaken,
			int32(pb.LeaveStatus_LEAVE_STATUS_PENDING),
			1000)

		if err != nil {
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
			return
		}

		// set user's leaves in cache
		leaves, err := db.GetLeavesByUserId(userId)
		if err != nil {
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_GET_LEAVES)
			return
		}
		err = r.SetLeavesByUserId(userId, leaves)
		if err != nil {
			log.Printf("Failed to set in Redis: %v\n", err)
		}

	default:

		// check that leave id exists
		checkLeave, err := db.GetLeaveByLeaveId(leaveId)
		if err != nil {
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_GET_LEAVES)
			return
		}
		if checkLeave == nil {
			log.Println("LeaveId does not exist")
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
			return
		}
		if err = db.UpdateLeaveStatus(leaveId, status); err != nil {
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_LEAVE)
			return
		}
	}

	// update user in DB and cache
	newLeaveBalance := currentLeaveBalance - daysToBeTaken
	if err = db.SetUserLeaveBalance(userId, newLeaveBalance); err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_SET_USER)
		return
	}
	user.LeaveBalance = newLeaveBalance
	err = r.SetUser(user)
	if err != nil {
		log.Printf("Failed to set \"user_%d\" in Redis\n", userId)
	}

	// return user data to client
	userBytes, _ := proto.Marshal(user)
	tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, userBytes, pb.Error_SUCCESS)
}
