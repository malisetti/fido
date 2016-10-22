package ops

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
	"time"

	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/util"
)

type AuthenticationRequestGeneration struct {
	AppID         string
	AcceptedAAIDs []string
}

func (authReqGen *AuthenticationRequestGeneration) CreateAuthenticationRequest(notary crypto.Notary) msg.AuthenticationRequest {
	authRequest := msg.AuthenticationRequest{}
	header := msg.OperationHeader{}
	challenge := GenerateChallenge()
	authRequest.Challenge = msg.ServerChallenge(challenge)
	header.ServerData = generateServerData(challenge, notary)
	authRequest.Header = header
	//authRequest.Header.Op = Operation.Auth
	authRequest.Header.AppID = authReqGen.AppID
	authRequest.Header.UPV = msg.Version{Major: 1, Minor: 0}

	authRequest.Policy = constructAuthenticationPolicy(authReqGen.AcceptedAAIDs)

	return authRequest
}

func generateChallenge() string {
	b := make([]byte, 32)
	rand.Read(b)

	return util.ToWebsafeBase64(base64.StdEncoding.EncodeToString(b))
}

func generateServerData(challenge string, notary crypto.Notary) string {
	dataToSign := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))) + "." + base64.StdEncoding.EncodeToString([]byte(challenge))
	signature := notary.Sign(dataToSign)

	return util.ToWebsafeBase64(base64.StdEncoding.EncodeToString([]byte(dataToSign + "." + signature)))
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
