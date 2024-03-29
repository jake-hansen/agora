package oauthinfoservice

import (
	"context"

	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

// OAuthInfoService processes information about OAuthInfos.
type OAuthInfoService struct {
	platformService domain.MeetingPlatformService
	repo            domain.OAuthInfoRepository
}

// CreateOAuthInfo creates an OAuthInfo in the repository by exchanging the given authorization with the MeetingPlatform.
func (o *OAuthInfoService) CreateOAuthInfo(ctx context.Context, authorization string, userID uint, platform *domain.MeetingPlatform) error {
	_, err := o.repo.GetByUserIDAndMeetingPlatformID(userID, platform.ID)
	if err == nil {
		return &domain.TokenExistsError{
			UserID:   userID,
			Platform: platform.Name,
		}
	}

	token, err := o.platformService.GetOAuthToken(ctx, authorization, platform)
	if err != nil {
		return err
	}

	oauthInfo := &domain.OAuthInfo{
		UserID:            userID,
		MeetingPlatformID: platform.ID,
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		TokenType:         token.TokenType,
		Expiry:            token.Expiry,
	}

	_, err = o.repo.Create(oauthInfo)
	return err
}

// GetOAuthInfo retrieves the OAuthInfo from the repository for the given User and MeetingPlatform combination. If
// the access tokens contained withing the OAuthInfo is expired, it is automatically refreshed.
func (o *OAuthInfoService) GetOAuthInfo(userID uint, platform *domain.MeetingPlatform) (*domain.OAuthInfo, error) {
	oauthInfo, err := o.repo.GetByUserIDAndMeetingPlatformID(userID, platform.ID)
	if err != nil {
		return nil, err
	}

	token := restoreToken(oauthInfo)

	if token.Valid() {
		return oauthInfo, nil
	} else {
		newToken, err := o.platformService.RefreshOAuthToken(context.Background(), &token, platform)
		if err != nil {
			return nil, err
		} else {
			oauthInfo.AccessToken = newToken.AccessToken
			oauthInfo.RefreshToken = newToken.RefreshToken
			oauthInfo.Expiry = newToken.Expiry
			oauthInfo.TokenType = newToken.TokenType
			err = o.repo.Update(oauthInfo)
			if err != nil {
				return nil, err
			} else {
				return oauthInfo, nil
			}
		}
	}
}

// GetAllAuthenticatedPlatforms gets all MeetingPlatforms a User with the provided userID has authenticated
// with.
func (o *OAuthInfoService) GetAllAuthenticatedPlatforms(userID uint) ([]*domain.MeetingPlatform, error) {
	oauthInfos, err := o.repo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var returnPlatforms []*domain.MeetingPlatform

	for _, oauthInfo := range oauthInfos {
		platform, err := o.platformService.GetByID(oauthInfo.MeetingPlatformID)
		if err != nil {
			return nil, err
		}

		returnPlatforms = append(returnPlatforms, platform)
	}

	return returnPlatforms, nil
}

func restoreToken(oauthInfo *domain.OAuthInfo) oauth2.Token {
	return oauth2.Token{
		AccessToken:  oauthInfo.AccessToken,
		RefreshToken: oauthInfo.RefreshToken,
		Expiry:       oauthInfo.Expiry,
		TokenType:    oauthInfo.TokenType,
	}
}
