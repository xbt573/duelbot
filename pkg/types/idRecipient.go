package types

import (
	"strconv"
)

type IdRecipient struct {
	Id int64
}

func (this *IdRecipient) Recipient() string {
	return strconv.FormatInt(this.Id, 10)
}
