package helpers

import (
	"github.com/sony/sonyflake"
)

func GenerateUID() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	return id
}
