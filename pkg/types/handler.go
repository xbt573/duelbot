package types

type Handler[T any] struct {
	next    *Handler[T]
	Handler func(arg T) bool
}

func (this *Handler[T]) SetNext(handler *Handler[T]) {
	this.next = handler
}

func (this *Handler[T]) Handle(arg T) {
	status := this.Handler(arg)
	if status {
		return
	}

	if this.next != nil {
		this.next.Handle(arg)
	}
}
