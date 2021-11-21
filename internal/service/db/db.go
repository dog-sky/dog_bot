package db

type DB interface {
	SetStatus(status string) error
}

type db struct{}

func New() DB {
	return &db{}
}

func (d *db) SetStatus(status string) error {
	return nil
}
