package msg

type VerificationMethodDescriptor struct {
	UserVerification uint32                      `json:"userVerification"`
	CaDesc           CodeAccuracyDescriptor      `json:"caDesc"`
	BaDesc           BiometricAccuracyDescriptor `json:"baDesc"`
	PaDesc           PatternAccuracyDescriptor   `json:"paDesc"`
}
