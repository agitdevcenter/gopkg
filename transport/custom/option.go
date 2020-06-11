package custom

import "time"

type Option func(*Holder)

func OptionService(service Service) Option {
	return func(h *Holder) {
		h.service = service
	}
}
func OptionHold(hold time.Duration) Option {
	return func(h *Holder) {
		h.hold = hold
	}
}
