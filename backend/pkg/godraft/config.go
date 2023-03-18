package godraft

import (
	"fmt"

	"github.com/gothing/draft"
)

type Config struct {
	Address string
}

// doc.draft.Add(auth.Service)
type Documentation struct {
	draft *draft.APIService
	cfg   Config
}

func Init() {
	draft.SetupDoc(draft.DocConfig{
		FrontURL:    "https://gothing.github.io/draft-front/",
		ActiveGroup: "auth",
		Groups: []draft.DocGroup{
			{ID: "auth", Name: "AUTH", Entries: []string{"https://localhost/godraft:scheme/"}},
		},
		Projects: []draft.DocProject{
			{
				ID:      "auth",
				Name:    "Auth",
				Host:    "markup2.ru",
				HostRC:  "host.docker.internal:443",
				HostDEV: "0.0.0.0:2047",
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

func (d *Documentation) Add(g draft.Group, groupHandlers ...draft.GroupHandler) {
	d.draft.Add(g, groupHandlers...)
}

func (d *Documentation) ListenAndServe() error {
	return d.draft.ListenAndServe(d.cfg.Address)
}
