package storage

type StorageInterface interface {
    StoreServerDataString(username, serverDataString string)

	GetUsername(serverDataString string) string

	Store(records RegistrationRecord[]) DuplicateKeyError, SystemError

	ReadRegistrationRecord(key string) RegistrationRecord

	Update(records []RegistrationRecord)    
}
