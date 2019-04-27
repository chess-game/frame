package uuid

import (
	"github.com/satori/go.uuid"
)

/**
创建全球唯一的UUID
*/
func CreateUUID() string {
	id, err := uuid.NewV4()

	if err != nil {
		return ""
	}
	if len(id.Bytes()) > 0 {
		return id.String()
	}
	return ""
}
