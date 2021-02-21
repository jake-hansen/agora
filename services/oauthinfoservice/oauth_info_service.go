package oauthinfoservice

import (
	"context"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

type OAuthInfoService struct {
	platformService	domain.MeetingPlatformService
	repo 			domain.OAuthInfoRepository
}

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

func (o *OAuthInfoService) GetOAuthInfo(userID uint, platform *domain.MeetingPlatform) (*domain.OAuthInfo, error) {
	oauthInfo, err := o.repo.GetByUserIDAndMeetingPlatformID(userID, platform.ID)
	if err != nil {
		return nil, err
	}

	token := restoreToken(oauthInfo)

	if token.Valid() {
		return oauthInfo, nil
	} else {
		newToken, err := o.platformService.RefreshOAuthToken(nil, &token, platform)
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

func restoreToken(oauthInfo *domain.OAuthInfo) oauth2.Token {
	return oauth2.Token{
		AccessToken: oauthInfo.AccessToken,
		RefreshToken: oauthInfo.RefreshToken,
		Expiry: oauthInfo.Expiry,
		TokenType: oauthInfo.TokenType,
	}
}
