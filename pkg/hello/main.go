package hello

import (
	"golangstart/pkg/utils"

	"github.com/hashicorp/go-uuid"
)

func Log() string {
	str := utils.TestStr()
	uuid, _ := uuid.GenerateUUID()
	return str + uuid
}
