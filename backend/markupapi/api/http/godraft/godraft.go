package godraft

import (
	"fmt"

	"github.com/gothing/draft"
)

type Config struct {
	Address string
}

type Documentation struct {
	draft *draft.APIService
	cfg   Config
}

func Init() {
	draft.SetupDoc(draft.DocConfig{
		FrontURL:    "https://gothing.github.io/draft-front/",
		ActiveGroup: "auth",
		Groups: []draft.DocGroup{
			{ID: "auth", Name: "AUTH", Entries: []string{"http://localhost:2047/godraft:scheme/"}},
		},
	})
}

func New(cfg Config) Documentation {
	doc := Documentation{cfg: cfg}
	doc.draft = draft.Create(draft.Config{
		DevMode: true,
		// MockMode: draft.MockDisable,
	})

	fmt.Printf("Server: http://%s\n", cfg.Address)
	fmt.Printf(" - http://%s/godraft:docs/\n", cfg.Address)
	fmt.Printf(" - http://%s/godraft:scheme/\n\n", cfg.Address)
	fmt.Println("API:")

	return doc
}

func (d *Documentation) ListenAndServe() error {
	return d.draft.ListenAndServe(d.cfg.Address)
}
