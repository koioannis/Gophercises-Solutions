package factory

type Parser interface {
	Parse([]byte) (map[string]string, error)
}