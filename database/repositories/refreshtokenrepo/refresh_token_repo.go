package refreshtokenrepo

import (
	"fmt"

	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type RefreshTokenRepo struct {
	DB *gorm.DB
}

// Create creates a new RefreshToken in the database.
func (r *RefreshTokenRepo) Create(token domain.RefreshToken) (uint, error) {
	if err := r.DB.Create(&token).Error; err != nil {
		return 0, fmt.Errorf("error creating Refresh Token: %w", err)
	}
	return token.ID, nil
}

// GetAll gets all RefreshTokens in the database.
func (r *RefreshTokenRepo) GetAll() ([]*domain.RefreshToken, error) {
	var refreshTokens []*domain.RefreshToken

	if err := r.DB.Find(&refreshTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all Refresh Tokens: %w", err)
	}
	return refreshTokens, nil
}

// GetByToken gets a RefreshToken from the database where the token_hash field
// matches the provided RefreshToken's hash.
func (r *RefreshTokenRepo) GetByToken(token domain.RefreshToken) (*domain.RefreshToken, error) {
	var foundToken *domain.RefreshToken

	if err := r.DB.Where("token_hash = ?", token.Value).Find(&foundToken).Error; err != nil {
		return nil, fmt.Errorf("error retrieving Refresh Token with hash %s: %w", token.Value, err)
	}
	return foundToken, nil
}

// Update updates the revoked status in the database for the provided RefreshToken.
func (r *RefreshTokenRepo) Update(token *domain.RefreshToken) error {
	if err := r.DB.Model(token).Updates(domain.RefreshToken{
		Revoked: token.Revoked,
	}).Error; err != nil {
		return fmt.Errorf("error updating Refresh Token with id %d: %w", token.ID, err)
	}
	return nil
}

// Delete deletes the RefreshToken from the database with the given ID.
func (r *RefreshTokenRepo) Delete(ID uint) error {
	if err := r.DB.Delete(&domain.RefreshToken{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting Refresh Token with id %d: %w", ID, err)
	}
	return nil
}

// GetByTokenNonceHash gets the first found RefreshToken in the database where the
// token_nonce_hash field matches the provided nonceHash.
func (r *RefreshTokenRepo) GetByTokenNonceHash(nonceHash string) (*domain.RefreshToken, error) {
	var foundToken = new(domain.RefreshToken)

	if err := r.DB.Where("token_nonce_hash = ?", nonceHash).First(foundToken).Error; err != nil {
		return nil, fmt.Errorf("error retrieving Refresh Token with nonce hash %s: %w", nonceHash, err)
	}
	return foundToken, nil
}
