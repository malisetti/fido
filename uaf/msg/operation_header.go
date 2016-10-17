package msg

type OperationHeader struct {
	UPV        Version     `json:"upv"`
	Op         Operation   `json:"op"`
	AppID      string      `json:"appID"`
	ServerData string      `json:"serverData"`
	Exts       []Extension `json:"exts"`
}
