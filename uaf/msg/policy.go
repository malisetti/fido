package msg

type Policy struct {
	Accepted   [][]MatchCriteria `json:"accepted"`
	Disallowed []MatchCriteria   `json:"disallowed"`
}
