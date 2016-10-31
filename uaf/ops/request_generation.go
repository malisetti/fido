package ops

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
	"time"

	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
)

func GenerateChallenge() string {
	b := make([]byte, 32)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}

func GenerateServerData(username, challenge string, notary crypto.Notary) string {
	dataToSign := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))) + "." + base64.StdEncoding.EncodeToString([]byte(username)) + "." + base64.StdEncoding.EncodeToString([]byte(challenge))
	signature := notary.Sign(dataToSign)

	return base64.URLEncoding.EncodeToString([]byte(dataToSign + "." + signature))
}

func ConstructAuthenticationPolicy(acceptedAaids []string) msg.Policy {
	p := msg.Policy{}
	accepted := make([][]msg.MatchCriteria, len(acceptedAaids))
	for i := range accepted {
		accepted[i] = make([]msg.MatchCriteria, 1)
	}

	for i, a := range accepted {
		matchCriteria := msg.MatchCriteria{}
		matchCriteria.AAID = make([]msg.AAID, 1)
		matchCriteria.AAID[0] = msg.AAID(acceptedAaids[i])
		a[0] = matchCriteria
		accepted[i] = a
	}
	p.Accepted = accepted

	return p
}
