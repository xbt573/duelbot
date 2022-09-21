package types

import (
	"strconv"
)

type IdRecipient struct {
	Id int64
}

func (i *IdRecipient) Recipient() string {
	return strconv.FormatInt(i.Id, 10)
}
