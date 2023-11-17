package template

type resolver interface {
	GetFromPath(pattern string) (matches []string, err error)
	GetPagesPath() (matches []string, err error)
	GetLayoutPath() (matches []string, err error)
}
