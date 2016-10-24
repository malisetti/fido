package tlv

import (
	"encoding/base64"
	"gitlab.pramati.com/seshachalamm/fido/uaf/util"
	"strconv"
)

type Tag struct {
	StatusID int
	ID       int
	Length   int
	Value    []byte
}

func (tag *Tag) String() string {
	ret := "Tag id:" + strconv.Itoa(tag.ID)
	tagsEnum := TagsEnum{}
	tagID, err := tagsEnum.Get(tag.ID)
	if err == nil {
		ret = ret + " Tag name: " + strconv.Itoa(tagID)
	}
	if tag.Value != nil {
		ret = ret + " Tag value:" + util.ToWebsafeBase64(base64.StdEncoding.EncodeToString(tag.Value))
	}

	return ret
}
