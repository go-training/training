package schedule

type Engine struct {
	*canceler
}

func New() *Engine {
	return &Engine{
		canceler: newCanceler(),
	}
}
