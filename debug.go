package flygo

//DebugTrace
func (a *App) DebugTrace(call func()) *App {
	if a.Config.Dev.Debug {
		call()
	}
	return a
}
