package server

import (
	pb "github.com/russellcxl/go-practice/assignment_3/protos"
	"github.com/russellcxl/go-practice/assignment_3/server/cache"
	"github.com/russellcxl/go-practice/assignment_3/server/server_controller"
	tcp "github.com/russellcxl/go-practice/assignment_3/tcp_manager"
	"log"
	"net"
)

type Server struct {
	redisRepo *cache.RedisRepo
	port     string
}

func NewServer(redisRepo *cache.RedisRepo, port string) *Server {
	return &Server{redisRepo: redisRepo, port: port}
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}

	log.Printf("Server running on port%s\n", s.port)

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v\n", err)
		}
		go handleConn(conn, s.redisRepo)
	}
}

func handleConn(c net.Conn, r *cache.RedisRepo) {
	log.Printf("Now serving %s\n", c.RemoteAddr().String())

	defer c.Close()

	for {

		// reads message from client
		msg := tcp.ReadMessage(c)
		if msg == nil || msg.GetCommand() == pb.Command_CLIENT_EXIT {
			log.Printf("No response from client on %s -- closing connection\n", c.RemoteAddr().String())
			break
		}

		// processes the message based on command and sends a response to the client
		cmd := msg.GetCommand()
		data := msg.GetData()

		log.Printf("--> CMD: %v\n", cmd)

		switch cmd {
		case pb.Command_CHECK_USER_LOGIN:
			server_controller.CheckLogin(c, r, data)

		case pb.Command_SET_LEAVE_APPLICATION:
			server_controller.SetLeave(c, r, data)

		case pb.Command_GET_LEAVE_APPLICATIONS_BY_USER_ID:
			server_controller.GetLeavesByUserId(c, r, data)

		case pb.Command_DELETE_LEAVE_APPLICATION:
			server_controller.DeleteLeaveApplication(c, r, data)

		default:
			log.Printf("Invalid command\n")
			tcp.SendMessage(c, pb.Command_UNKNOWN_METHOD, nil, pb.Error_INVALID_COMMAND)
		}
	}
}
