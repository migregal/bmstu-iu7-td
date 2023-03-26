package render

type Config struct {
	HTML HTMLOpts
}

type Renderer struct {
	cfg Config
}

func New(cfg Config) Renderer {
	return Renderer{cfg: cfg}
}
