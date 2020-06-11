package custom

import "time"

type Holder struct {
	hold    time.Duration
	service Service
}

func New(opts ...Option) *Holder {
	h := &Holder{}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *Holder) Service() Service {
	return h.service
}

func (h *Holder) Hold() {
	if h.hold > 0 {
		time.Sleep(h.hold)
	}
}
