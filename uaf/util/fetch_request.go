package util

import (
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/ops"
)

type FetchRequest struct {
	AppID        string
	AllowedAAIDs []string
}

func (fetchRequest *FetchRequest) GetRegistrationRequest(username string) msg.RegistrationRequest {
	notary := NotaryImpl{}
	request := ops.RegistrationRequestGeneration{fetchRequest.AppID, fetchRequest.AllowedAAIDs}

	return request.CreateRegistrationRequest(username, notary)
}

func (fetchRequest *FetchRequest) GetAuthenticationRequest() msg.AuthenticationRequest {
	authReq := ops.AuthenticationRequestGeneration{fetchRequest.AppID, fetchRequest.AllowedAAIDs}
	notary := NotaryImpl{}

	return authReq.CreateAuthenticationRequest(notary)
}
