package media

type Media interface {
}

type impl struct{}

func New() Media {
	return &impl{}
}
