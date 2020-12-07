package flygo

//Debug trace
func (a *App) DebugTrace(fn func()) *App {
	if a.Config.Flygo.Dev.Debug {
		fn()
	}
	return a
}
