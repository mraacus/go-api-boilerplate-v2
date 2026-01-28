package handler

type UniqueClient interface {
	// Add external client interfaces here
}

type Handler struct {
	SomeClient 		UniqueClient
}

type NewHandlerParam struct {
	SomeClient   	UniqueClient
}

func New(param NewHandlerParam) (*Handler, error) {
	return &Handler{
		SomeClient:  param.SomeClient,
	}, nil
}