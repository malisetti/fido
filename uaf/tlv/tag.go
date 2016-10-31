package tlv

import (
	"encoding/base64"
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
		ret = ret + " Tag value:" + base64.URLEncoding.EncodeToString(tag.Value)
	}

	return ret
}
