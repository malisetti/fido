package msg

import (
	"errors"
)

//defines all the constants and types for uaf
const (
	FIDO_VERSION = "v1.0"

	//Key Protection Types
	KEY_PROTECTION_SOFTWARE       = 0x01
	KEY_PROTECTION_HARDWARE       = 0x02
	KEY_PROTECTION_TEE            = 0x04
	KEY_PROTECTION_SECURE_ELEMENT = 0x08
	KEY_PROTECTION_REMOTE_HANDLE  = 0x10

	//User Verification Methods
	USER_VERIFY_PRESENCE    = 0x01
	USER_VERIFY_FINGERPRINT = 0x02
	USER_VERIFY_PASSCODE    = 0x04
	USER_VERIFY_VOICEPRINT  = 0x08
	USER_VERIFY_FACEPRINT   = 0x10
	USER_VERIFY_LOCATION    = 0x20
	USER_VERIFY_EYEPRINT    = 0x40
	USER_VERIFY_PATTERN     = 0x80
	USER_VERIFY_HANDPRINT   = 0x100
	USER_VERIFY_NONE        = 0x200
	USER_VERIFY_ALL         = 0x400

	//Matcher Protection Types
	MATCHER_PROTECTION_SOFTWARE = 0x01
	MATCHER_PROTECTION_TEE      = 0x02
	MATCHER_PROTECTION_ON_CHIP  = 0x04

	//Authenticator Attachment Hints
	ATTACHMENT_HINT_INTERNAL    = 0x01
	ATTACHMENT_HINT_EXTERNAL    = 0x02
	ATTACHMENT_HINT_WIRED       = 0x04
	ATTACHMENT_HINT_WIRELESS    = 0x08
	ATTACHMENT_HINT_NFC         = 0x10
	ATTACHMENT_HINT_BLUETOOTH   = 0x20
	ATTACHMENT_HINT_NETWORK     = 0x40
	ATTACHMENT_HINT_READY       = 0x80
	ATTACHMENT_HINT_WIFI_DIRECT = 0x100

	//Transaction Confirmation Display Types
	TRANSACTION_CONFIRMATION_DISPLAY_ANY                 = 0x01
	TRANSACTION_CONFIRMATION_DISPLAY_PRIVILEGED_SOFTWARE = 0x02
	TRANSACTION_CONFIRMATION_DISPLAY_TEE                 = 0x04
	TRANSACTION_CONFIRMATION_DISPLAY_HARDWARE            = 0x08
	TRANSACTION_CONFIRMATION_DISPLAY_REMOTE              = 0x10

	//Tags used for crypto algorithms and types
	UAF_ALG_SIGN_SECP256R1_ECDSA_SHA256_RAW = 0x01
	UAF_ALG_SIGN_SECP256R1_ECDSA_SHA256_DER = 0x02
	UAF_ALG_SIGN_RSASSA_PSS_SHA256_RAW      = 0x03
	UAF_ALG_SIGN_RSASSA_PSS_SHA256_DER      = 0x04
	UAF_ALG_SIGN_SECP256K1_ECDSA_SHA256_RAW = 0x05
	UAF_ALG_SIGN_SECP256K1_ECDSA_SHA256_DER = 0x06

	//Public Key Representation Formats
	UAF_ALG_KEY_ECC_X962_RAW     = 0x100
	UAF_ALG_KEY_ECC_X962_DER     = 0x101
	UAF_ALG_KEY_RSA_2048_PSS_RAW = 0x102
	UAF_ALG_KEY_RSA_2048_PSS_DER = 0x103

	//UAF Status Codes
	OPERATION_COMPLETED              = 1200
	MESSAGE_ACCEPTED                 = 1202
	BAD_REQUEST                      = 1400
	UNAUTHORIZED                     = 1401
	FORBIDDEN                        = 1403
	NOT_FOUND                        = 1404
	REQUEST_TIMEOUT                  = 1408
	UNKOWN_AAID                      = 1480
	UNKOWN_KEYID                     = 1481
	CHANNEL_BINDING_REFUSED          = 1490
	INVALID_REQUEST                  = 1491
	UNACCEPTABLE_AUTHENTICATOR       = 1492
	REVOKED_AUTHENTICATOR            = 1493
	UNACCEPTABLE_KEY                 = 1494
	UNACCEPTABLE_ALGORITHM           = 1495
	UNACCEPTABLE_ATTESTATION         = 1496
	UNACCEPTABLE_CLIENT_CAPABILITIES = 1497
	UNACCEPTABLE_CONTENT             = 1498
	INTERNAL_SERVER_ERROR            = 1500
)

//UAF Status Codes
var statusCodes = map[int]string{
	OPERATION_COMPLETED: `OK. Operation completed`,
	MESSAGE_ACCEPTED: `Accepted. Message accepted, but not completed at this time. The RP may need time to process the attestation, run risk scoring, etc. The
server should not send an authenticationToken with a 1202 response`,
	BAD_REQUEST:     `Bad Request. The server did not understand the message`,
	UNAUTHORIZED:    `Unauthorized. The userid must be authenticated to perform this operation, or this KeyID is not associated with this UserID.`,
	FORBIDDEN:       `Forbidden. The userid is not allowed to perform this operation. Client should not retry`,
	NOT_FOUND:       `Not Found.`,
	REQUEST_TIMEOUT: `Request Timeout.`,
	UNKOWN_AAID:     `Unknown AAID. The server was unable to locate authoritative metadata for the AAID.`,
	UNKOWN_KEYID: `Unknown KeyID. The server was unable to locate a registration for the given UserID and KeyID combination.
This error indicates that there is an invalid registration on the user's device. It is recommended that FIDO UAF Client deletes the key
from local device when this error is received.`,
	CHANNEL_BINDING_REFUSED: `Channel Binding Refused. The server refused to service the request due to a missing or mismatched channel binding(s).`,
	INVALID_REQUEST: `Request Invalid. The server refused to service the request because the request message nonce was unknown, expired or the server has
previously serviced a message with the same nonce and user ID.`,
	UNACCEPTABLE_AUTHENTICATOR: `Unacceptable Authenticator. The authenticator is not acceptable according to the server's policy, for example because the capability
registry used by the server reported different capabilities than client-side discovery.`,
	REVOKED_AUTHENTICATOR: `Revoked Authenticator. The authenticator is considered revoked by the server.`,
	UNACCEPTABLE_KEY:      `Unacceptable Key. The key used is unacceptable. Perhaps it is on a list of known weak keys or uses insecure parameter choices.`,
	UNACCEPTABLE_ALGORITHM: `Unacceptable Algorithm. The server believes the authenticator to be capable of using a stronger mutually-agreeable algorithm than was
presented in the request.`,
	UNACCEPTABLE_ATTESTATION: `Unacceptable Attestation. The attestation(s) provided were not accepted by the server.`,
	UNACCEPTABLE_CLIENT_CAPABILITIES: `Unacceptable Client Capabilities. The server was unable or unwilling to use required capabilities provided supplementally to the
authenticator by the client software.`,
	UNACCEPTABLE_CONTENT:  `Unacceptable Content. There was a problem with the contents of the message and the server was unwilling or unable to process it.`,
	INTERNAL_SERVER_ERROR: `Internal Server Error`,
}

func UAFStatusCode(statusCode int) (string, error) {
	if status, ok := statusCodes[statusCode]; ok {
		return status, nil
	}

	return "", errors.New("Status code does not exist.")
}

type DOMString string

type Operation DOMString

type AAID DOMString

type KeyID DOMString

type ServerChallenge DOMString

type VerificationMethodANDCombinations []VerificationMethodDescriptor
