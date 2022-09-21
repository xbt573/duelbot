package user

import (
	"github.com/xbt573/duelbot/pkg/types"
)

var DefaultClient = types.NewUserDB()

func Set(key int64, value int) {
	DefaultClient.Set(key, value)
}

func Get(key int64) int {
	return DefaultClient.Get(key)
}
