package refreshtokenrepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type RefreshTokenRepo struct {
	DB *gorm.DB
}

func (r *RefreshTokenRepo) Create(token domain.RefreshToken) (uint, error) {
	if err := r.DB.Create(&token).Error; err != nil {
		return 0, fmt.Errorf("error creating Refresh Token: %w", err)
	}
	return token.ID, nil
}

func (r *RefreshTokenRepo) GetAll() ([]*domain.RefreshToken, error) {
	var refreshTokens []*domain.RefreshToken

	if err := r.DB.Find(&refreshTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all Refresh Tokens: %w", err)
	}
	return refreshTokens, nil
}

func (r *RefreshTokenRepo) GetByToken(token domain.RefreshToken) (*domain.RefreshToken, error) {
	hash, _ := token.Value.Value()
	var foundToken *domain.RefreshToken

	if err := r.DB.Where("token_hash = ?", hash).Find(&foundToken).Error; err != nil {
		return nil, fmt.Errorf("error retrieving Refresh Token with hash %s: %w", hash, err)
	}
	return foundToken, nil
}

func (r *RefreshTokenRepo) Update(token *domain.RefreshToken) error {
	panic("implement me")
}

func (r *RefreshTokenRepo) Delete(ID uint) error {
	if err := r.DB.Delete(&domain.RefreshToken{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting Refresh Token with id %d: %w", ID, err)
	}
	return nil
}
