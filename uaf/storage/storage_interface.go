package storage

import (
    "gitlab.pramati.com/seshachalamm/fido/uaf/storage"
)

type StorageInterface interface {
    StoreServerDataString(username, serverDataString string)

	GetUsername(serverDataString string) string

	Store(records storage.RegistrationRecord[]) DuplicateKeyException, SystemErrorException

	ReadRegistrationRecord(key string) RegistrationRecord

	Update(records []RegistrationRecord)    
}
