package flygo

var Logger = &logger{}

//Define logger struct
type logger struct {
}

func (log *logger) Name() string {
	return "FlygoLogger"
}

func (log *logger) Type() string {
	return "AFTER"
}

func (log *logger) Pattern() string {
	return "/**"
}

func (log *logger) Process() FilterHandler {
	return func(c *FilterContext) {
		//logger
		app.Info("[%s]%s", c.RequestMethod, c.RequestURI)
	}
}
