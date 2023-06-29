package domain

type RecognizeRequest struct {
	Face []byte
}

type RecognizeResponse struct {
	PersonName   string
	FaceLocation Location
}

type Location struct {
	Left   uint32
	Top    uint32
	Right  uint32
	Bottom uint32
}
