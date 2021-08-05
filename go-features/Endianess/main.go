package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	bytes1 := []byte("hello")

	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(bytes1)))

	fmt.Println(bytes1)
	fmt.Println(length)
	fmt.Println("******")

	length = append(length, bytes1...)
	fmt.Println(length)
	fmt.Println("******")


	receiver := binary.BigEndian.Uint64(length)
	fmt.Println(receiver)

	bs := convertIntToBytes(360287970189639680)
	fmt.Println(bs)

}

func convertIntToBytes(number uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, number)
	return bs
}
