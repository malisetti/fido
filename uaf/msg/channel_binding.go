package msg

type ChannelBinding struct {
	ServerEndPoint       DOMString `json:"serverEndPoint"`
	TlsServerCertificate DOMString `json:"tlsServerCertificate"`
	TlsUnique            DOMString `json:"tlsUnique"`
	CID_Pubkey           DOMString `json:"cid_pubkey"`
}
