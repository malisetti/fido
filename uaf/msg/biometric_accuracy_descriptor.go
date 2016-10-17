package msg

type BiometricAccuracyDescriptor struct {
	FAR                  float32 `json:"FAR"`
	FRR                  float32 `json:"FRR"`
	EER                  float32 `json:"EER"`
	FAAR                 float32 `json:"FAAR"`
	MaxReferenceDataSets uint16  `json:"maxReferenceDataSets"`
	MaxRetries           uint16  `json:"maxRetries"`
	BlockSlowdown        uint16  `json:"blockSlowdown"`
}
