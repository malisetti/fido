package msg

type Extension struct {
	ID            string `json:"id"`
	Data          string `json:"data"`
	FailIfUnknown bool   `json:"fail_if_unknown"`
}
