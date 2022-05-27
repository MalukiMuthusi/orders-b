package postgres

type Postgres struct{}

func New() (*Postgres, error) {
	return &Postgres{}, nil
}
