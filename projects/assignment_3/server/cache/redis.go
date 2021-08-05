package cache

import (
	"context"
	"fmt"
	pb "git.garena.com/russell.chanxl/be-class/assignment_3/protos"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"log"
	"strconv"
	"time"
)

var (
	ctx        = context.Background()
	expiration = time.Minute
	port       = "6379"
	password   = "" // no password set
	DB         = 0  // default DB
)

type RedisRepo struct {
	cli *redis.Client
}

func NewRedisRepo() *RedisRepo {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: password,
		DB:       DB,
	})
	log.Printf("Redis running on port:%s\n", port)
	return &RedisRepo{cli: cli}
}

func (r *RedisRepo) SetUser(user *pb.User) error {
	key := createUserKey(user.GetUserId())
	userBytes, _ := proto.Marshal(user)

	log.Printf("Setting key=\"%s\" in Redis...\n", key)

	err := r.cli.Set(ctx, key, userBytes, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) GetUser(userId uint64) (*pb.User, error) {
	key := createUserKey(userId)

	log.Printf("Getting key=\"%s\" in Redis...", key)

	userBytes, err := r.cli.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	var user pb.User
	err = proto.Unmarshal(userBytes, &user)
	return &user, nil
}

func (r *RedisRepo) SetLeavesByUserId(userId uint64, leaves *pb.Leaves) error {
	key := createUserLeavesKey(userId)
	leavesBytes, _ := proto.Marshal(leaves)
	log.Printf("Setting key=\"%s\" in Redis...\n", key)
	err := r.cli.Set(ctx, key, leavesBytes, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) GetLeavesByUserId(userId uint64) (*pb.Leaves, error) {
	key := createUserLeavesKey(userId)
	log.Printf("Getting key=\"%s\" in Redis...", key)
	leavesBytes, err := r.cli.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	var leaves pb.Leaves
	err = proto.Unmarshal(leavesBytes, &leaves)
	return &leaves, nil
}

func createUserKey(userId uint64) string {
	IdString := strconv.Itoa(int(userId))
	return fmt.Sprintf("user_%s", IdString)
}

func createUserLeavesKey(userId uint64) string {
	IdString := strconv.Itoa(int(userId))
	return fmt.Sprintf("leaves_for_user_%s", IdString)
}
