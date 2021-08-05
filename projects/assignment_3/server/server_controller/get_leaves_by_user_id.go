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

func GetLeavesByUserId(conn net.Conn, r *cache.RedisRepo, data []byte) {
	var user pb.User
	err := proto.Unmarshal(data, &user)
	if err != nil {
		log.Printf("Failed to unmarshal into User: %v\n", err)
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_GET_LEAVES)
		return
	}

	userId := user.GetUserId()

	// check cache; if not there, get from DB and update cache
	leaves, err := r.GetLeavesByUserId(userId)
	if err != nil {
		log.Printf("Failed to get from Redis\n")
		leaves, err = db.GetLeavesByUserId(userId)
		if err != nil {
			tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_GET_LEAVES)
			return
		}
		err = r.SetLeavesByUserId(userId, leaves)
		if err != nil {
			log.Printf("Failed to set in Redis: %v\n", err)
		}
	} else {
		log.Println("Successfully retrieved from Redis")
	}

	// return leaves to the client
	leavesBytes, _ := proto.Marshal(leaves)
	tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, leavesBytes, pb.Error_SUCCESS)
}
