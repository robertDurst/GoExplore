package GoExplore

type Token interface {
	GetType() string
	PrettyFormat() string
}
