package server_controller

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	pb "github.com/russellcxl/go-practice/assignment_3/protos"
	"github.com/russellcxl/go-practice/assignment_3/server/cache"
	tcp "github.com/russellcxl/go-practice/assignment_3/tcp_manager"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

func CheckLogin(conn net.Conn, r *cache.RedisRepo, data []byte) {
	key := "123456789012345678901234" // this should be stored in a private file
	var u pb.User
	err := proto.Unmarshal(data, &u)
	if err != nil {
		log.Printf("failed to unmarshal into user struct: %v\n", err)
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_CHECK_LOGIN)
		return
	}
	userId := u.GetUserId()
	password := u.GetPassword()

	user, err := getUserFromCache(r, userId)
	if err != nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_FAILED_TO_CHECK_LOGIN)
		return
	}

	// check input credentials against DB credentials
	if user == nil {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_INVALID_USER_ID)
		return
	}
	if encrypt(key, password) != user.GetPassword() {
		tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, nil, pb.Error_LOGIN_CREDENTIALS_INVALID)
		return
	}

	resp, _ := proto.Marshal(user)

	// for simplicity of the protobuf, this assumes that the client knows what to expect
	// ideally, there should be messages for req and res
	tcp.SendMessage(conn, pb.Command_UNKNOWN_METHOD, resp, pb.Error_SUCCESS)
}

//******************* PASSWORD ENCRYPTION *******************//

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encrypt(key, text string) string {

	// create a new cipher block using your key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	// convert the input string into byte array
	plaintext := []byte(text)

	// create an empty bytes array using the length of the above byte array; this will be used as the destination
	ciphertext := make([]byte, len(plaintext))

	// create a stream cipher (too hard to understand)
	stream := cipher.NewCFBEncrypter(block, iv)

	// fills up destination byte[] (ciphertext)
	stream.XORKeyStream(ciphertext, plaintext)
	return encodeB64(ciphertext)
}

func decrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext := decodeB64(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}

// changes bytes array to encoded string
func encodeB64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// decodes encoded string to bytes array
func decodeB64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
