package got_test

import (
	"testing"

	"github.com/xenvoid404/got"
)

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required"`
}

func TestStruct(t *testing.T) {
	v := got.New()
	if err := v.Struct(User{}); err != nil {
		t.Fatal(err)
	}
}
