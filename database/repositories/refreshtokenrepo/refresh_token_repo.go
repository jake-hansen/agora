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
	if err := r.DB.Model(token).Updates(domain.RefreshToken{
		Value:           token.Value,
		ExpiresAt:       token.ExpiresAt,
		TokenNonceHash:  token.TokenNonceHash,
		ParentTokenHash: token.ParentTokenHash,
		UserID:          token.UserID,
		Revoked:         token.Revoked,
	}).Error; err != nil {
		return fmt.Errorf("error updating Refresh Token with id %d: %w", token.ID, err)
	}
	return nil
}

func (r *RefreshTokenRepo) Delete(ID uint) error {
	if err := r.DB.Delete(&domain.RefreshToken{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting Refresh Token with id %d: %w", ID, err)
	}
	return nil
}

func (r *RefreshTokenRepo) GetByTokenHash(hash string) (*domain.RefreshToken, error) {
	var foundToken = new(domain.RefreshToken)

	if err := r.DB.Where("token_hash = ? AND deleted_at IS NULL", hash).First(foundToken).Error; err != nil {
		return nil, fmt.Errorf("error retrieving Refresh Token with hash %s: %w", hash, err)
	}
	return foundToken, nil
}

func (r *RefreshTokenRepo) GetByParentTokenHash(hash string) (*domain.RefreshToken, error) {
	var foundToken = new(domain.RefreshToken)

	if err := r.DB.Where("parent_token_hash = ?", hash).First(foundToken).Error; err != nil {
		return nil, fmt.Errorf("error retrieving Refresh Token with parent hash %s: %w", hash, err)
	}
	return foundToken, nil
}

func (r *RefreshTokenRepo) GetByTokenNonceHash(nonceHash string) (*domain.RefreshToken, error) {
	var foundToken = new(domain.RefreshToken)

	if err := r.DB.Where("token_nonce_hash = ?", nonceHash).First(foundToken).Error; err != nil {
		return nil, fmt.Errorf("error retrieving Refresh Token with nonce hash %s: %w", nonceHash, err)
	}
	return foundToken, nil
}
