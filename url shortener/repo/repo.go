package repo

// Repo define interface for storage engines
type Repo interface {
	Shorten(string) (string, error)
	Expand(string) (string, error)
}
