package handler

type Handler struct {
	Health *healthHandler
}

func New() *Handler {
	return &Handler{Health: NewHealthHandler()}
}
