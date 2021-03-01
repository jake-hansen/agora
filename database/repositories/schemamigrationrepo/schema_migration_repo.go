package schemamigrationrepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type SchemaMigrationRepo struct {
	DB *gorm.DB
}

func (s *SchemaMigrationRepo) GetSchemaMigrationByVersion(version int) (migration *domain.SchemaMigration, err error) {
	m := new(domain.SchemaMigration)
	if err := s.DB.First(m, version).Error; err != nil {
		return nil, fmt.Errorf("error retrieving schema migration with version %d: %w", version, err)
	}
	return m, nil
}

func (s *SchemaMigrationRepo) GetSchemaMigration() (*domain.SchemaMigration, error) {
	m := new(domain.SchemaMigration)
	if err := s.DB.First(m).Error; err != nil {
		return nil, fmt.Errorf("error retrieving a schema migration")
	}
	return m, nil
}
