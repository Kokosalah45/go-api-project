package database

type DBService interface {
	Health() map[string]string
	Close() error
	GetDBName() string
}
