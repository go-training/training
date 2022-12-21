package schedule

type Engine struct {
	canceler *canceler
}

func New() *Engine {
	return &Engine{
		canceler: newCanceler(),
	}
}
