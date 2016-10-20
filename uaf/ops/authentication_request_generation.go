package ops

import (
	"crypto/rand"
	"encoding/base64"
	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/util"
	"strconv"
	"time"
)

func CreateRegistrationRequest(username string, appID string, acceptedAaids []string, notary crypto.Notary) msg.RegistrationRequest {
	challenge := generateChallenge()
	serverDataString := generateServerData(username, challenge, notary)

	return createRegistrationRequest(username, serverDataString, challenge, appID, acceptedAaids)
}

func generateChallenge() string {
	b := make([]byte, 32)
	rand.Read(b)

	return util.ToWebsafeBase64(base64.StdEncoding.EncodeToString(b))
}

func generateServerData(username, challenge string, notary crypto.Notary) string {
	dataToSign := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))) + "." + base64.StdEncoding.EncodeToString([]byte(username)) + "." + base64.StdEncoding.EncodeToString([]byte(challenge))
	signature := notary.Sign(dataToSign)

	return util.ToWebsafeBase64(base64.StdEncoding.EncodeToString([]byte(dataToSign + "." + signature)))
}

func createRegistrationRequest(username, serverData, challenge, appID string, acceptedAaids []string) msg.RegistrationRequest {
	regRequest := msg.RegistrationRequest{}
	header := msg.OperationHeader{}
	header.ServerData = serverData
	regRequest.Header = header
	regRequest.Header.Op = msg.Reg
	regRequest.Header.AppID = appID
	regRequest.Header.UPV = msg.Version{1, 0}
	regRequest.Challenge = msg.ServerChallenge(challenge)
	regRequest.Policy = constructAuthenticationPolicy(acceptedAaids)
	regRequest.Username = msg.DOMString(username)

	return regRequest
}

func constructAuthenticationPolicy(acceptedAaids []string) msg.Policy {
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
