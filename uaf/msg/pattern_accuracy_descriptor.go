package msg

type PatternAccuracyDescriptor struct {
	MinComplexity uint32 `json:"minComplexity"`
	MaxRetries    uint16 `json:"maxRetries"`
	BlockSlowdown uint16 `json:"blockSlowdown"`
}
