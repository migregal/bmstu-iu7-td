package godraft

import (
	"fmt"
	"markup2/markupapi/api/http/v1/auth"

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
		Projects: []draft.DocProject{
			{
				ID:      "auth",
				Name:    "Auth",
				Host:    "markup2.ru",
				HostDEV: "localhost:2047",
			},
		},
	})
}

func New(cfg Config) *Documentation {
	doc := Documentation{cfg: cfg}
	doc.draft = draft.Create(draft.Config{
		DevMode:      true,
		ClientConfig: draft.ClientConfig{SkipVerifyCert: true},
		MockMode:     draft.MockEnable,
	})

	doc.draft.Add(auth.Service)

	fmt.Printf("Server: http://%s\n", cfg.Address)
	fmt.Printf(" - http://%s/godraft:docs/\n", cfg.Address)
	fmt.Printf(" - http://%s/godraft:scheme/\n\n", cfg.Address)
	fmt.Println("API:")

	for _, u := range doc.draft.URLs() {
		fmt.Printf(" --> http://%s%s\n", cfg.Address, u)
	}

	fmt.Println("")

	return &doc
}

func (d *Documentation) ListenAndServe() error {
	return d.draft.ListenAndServe(d.cfg.Address)
}
