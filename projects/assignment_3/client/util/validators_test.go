package util

import (
	"fmt"
	"testing"
)

func Test_validateCredentials(t *testing.T) {
	got := ValidateLoginFormat("123 pass")
	fmt.Println(got)
}
