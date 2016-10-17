package msg

type Transaction struct {
	ContentType                 DOMString                           `json:"contentType"`
	Content                     DOMString                           `json:"content"`
	TcDisplayPNGCharacteristics DisplayPNGCharacteristicsDescriptor `json:"tcDisplayPNGCharacteristics"`
}
