package handlers

// ProvideHandlerManger returns a HandlerManager configured with the provided *[]Handler.
func ProvideHandlerManager(handlers *[]Handler) *HandlerManager {
	return NewHandlerManager(handlers)
}
