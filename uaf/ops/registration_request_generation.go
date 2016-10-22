package ops

import (
	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
)

type RegistrationRequestGeneration struct {
	AppID         string
	AcceptedAAIDs []string
}

func (regRequestGen *RegistrationRequestGeneration) CreateRegistrationRequest(username string, notary crypto.Notary) msg.RegistrationRequest {
	challenge := GenerateChallenge()
	serverDataString := GenerateServerData(username, challenge, notary)

	return createRegistrationRequest(username, serverDataString, challenge, regRequestGen.AppID, regRequestGen.AcceptedAAIDs)
}

func createRegistrationRequest(username, serverData, challenge, appID string, acceptedAaids []string) msg.RegistrationRequest {
	regRequest := msg.RegistrationRequest{}
	header := msg.OperationHeader{}
	header.ServerData = serverData
	regRequest.Header = header
	regRequest.Header.Op = msg.Reg
	regRequest.Header.AppID = appID
	regRequest.Header.UPV = msg.Version{Major: 1, Minor: 0}
	regRequest.Challenge = msg.ServerChallenge(challenge)
	regRequest.Policy = ConstructAuthenticationPolicy(acceptedAaids)
	regRequest.Username = msg.DOMString(username)

	return regRequest
}
