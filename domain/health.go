package domain

type Health struct {
	Healthy	bool
	Reason	string
}

type SchemaMigration struct {
	Version	int
	Dirty	int
}

type HealthService interface {
	GetHealth() (*Health, error)
}

type SchemaMigrationRepo interface {
	GetSchemaMigrationByVersion(version int) (*SchemaMigration, error)
}
