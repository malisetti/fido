package tlv

type Tags struct {
	InTags map[int]*Tag
}

func (tags *Tags) Add(tag *Tag) {
	tags.InTags[tag.ID] = tag
}

func (tags *Tags) AddAll(all *Tags) {
	for _, tag := range all.InTags {
		all.Add(tag)
	}
}

func ToUAFV1TLV() string {
	return ""
}

func (tags *Tags) GetTags() map[int]*Tag {
	return tags.InTags
}

// func (tags *Tags) String() string {
// 	res := ""
// 	for _, tag := range tags.InTags {
// 		res += ", "
// 		res += tag.String()
// 	}

// 	if len(res) > 0 {
// 		return "{" + res[:1] + "}"
// 	}

// 	return "{}"
// }
