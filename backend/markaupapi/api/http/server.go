package http

type Server interface {
	ListenAndServe() error
}
