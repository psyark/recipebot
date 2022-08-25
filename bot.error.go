package recipebot

type FancyError struct {
	err error
}

func (e *FancyError) Error() string {
	return e.err.Error()
}
