package domain

type SchemaMigration struct {
	Version int
	Dirty   int
}

type SchemaMigrationRepo interface {
	GetSchemaMigrationByVersion(version int) (*SchemaMigration, error)
	GetSchemaMigration() (*SchemaMigration, error)
}
