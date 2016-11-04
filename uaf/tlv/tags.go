package tlv

type Tags struct {
	Tags map[int]Tag
}

func (tags *Tags) Add(tag Tag) {
	tags.Tags[tag.ID] = tag
}

func (tags *Tags) AddAll(all Tags) {
	for _, tag := range all.Tags {
		all.Add(tag)
	}
}

func ToUAFV1TLV() string {
	return ""
}

func (tags *Tags) GetTags() map[int]Tag {
	return tags.Tags
}

func (tags *Tags) String() string {
	res := ""
	for _, tag := range tags.Tags {
		res += ", "
		res += tag.String()
	}

	if len(res) > 0 {
		return "{" + res[:1] + "}"
	}

	return "{}"
}
