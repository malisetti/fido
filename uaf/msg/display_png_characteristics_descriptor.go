package msg

type DisplayPNGCharacteristicsDescriptor struct {
	Width       uint32             `json:"width"`
	Height      uint32             `json:"height"`
	BitDepth    []byte             `json:"bitDepth"`
	ColorType   []byte             `json:"colorType"`
	Compression []byte             `json:"compression"`
	Filter      []byte             `json:"filter"`
	Interlace   []byte             `json:"interlace"`
	Plte        []RgbPalletteEntry `json:"plte"`
}
