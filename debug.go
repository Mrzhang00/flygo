package flygo

//DebugTrace
func (a *App) DebugTrace(fn func()) *App {
	if a.Config.Debug {
		fn()
	}
	return a
}
