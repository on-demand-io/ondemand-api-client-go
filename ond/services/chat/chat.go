package chat

type Chat interface {
}

type impl struct{}

func New() Chat {
	return &impl{}
}
