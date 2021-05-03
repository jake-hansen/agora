package domain

// SchemaMigration represents the data provided by the go migrate tool.
type SchemaMigration struct {
	Version int
	Dirty   int
}

// SchemaMigrationRepo stores information about SchemaMigrations.
type SchemaMigrationRepo interface {
	GetSchemaMigrationByVersion(version int) (*SchemaMigration, error)
	GetSchemaMigration() (*SchemaMigration, error)
}
