package msg

type FinalChallengeParams struct {
	AppID          DOMString       `json:"appID"`
	Challenge      ServerChallenge `json:"challenge"`
	FacetID        DOMString       `json:"facetID"`
	ChannelBinding ChannelBinding  `json:"channelBinding"`
}
