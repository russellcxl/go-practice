package server_controller

import (
	pb "github.com/russellcxl/go-practice/assignment_3/protos"
	"github.com/russellcxl/go-practice/assignment_3/server/cache"
	db "github.com/russellcxl/go-practice/assignment_3/server/database"
	"log"
)

// getUserFromCache checks the cache for user data. If not found, retrieves from DB and stores it in the cache
func getUserFromCache(r *cache.RedisRepo, userId uint64) (*pb.User, error) {
	var user *pb.User
	var err error
	user, err = r.GetUser(userId)
	if err != nil {
		log.Printf("Failed to get \"user_%d\" in Redis\n", userId)
		user, err = db.GetUser(userId)
		if err != nil {
			return nil, err
		}
		err = r.SetUser(user)
		if err != nil {
			log.Printf("Failed to set \"user_%d\" in Redis\n", userId)
			return nil, err
		}
	} else {
		log.Printf("Successfully retrieved \"user_%d\" from Redis\n", userId)
	}
	return user, nil
}
