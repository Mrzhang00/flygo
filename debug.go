package flygo

//DebugTrace
func (a *App) DebugTrace(fn func()) *App {
	if a.Config.Dev.Debug {
		fn()
	}
	return a
}
