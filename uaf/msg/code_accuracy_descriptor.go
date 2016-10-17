package msg

type CodeAccuracyDescriptor struct {
	Base          uint16 `json:"base"`
	MinLength     uint16 `json:"minLength"`
	MaxRetries    uint16 `json:"maxRetries"`
	BlockSlowdown uint16 `json:"blockSlowdown"`
}
