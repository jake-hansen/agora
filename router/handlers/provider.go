package handlers

func ProvideHandlerManager(handlers *[]Handler) *HandlerManager {
	return NewHandlerManager(handlers)
}