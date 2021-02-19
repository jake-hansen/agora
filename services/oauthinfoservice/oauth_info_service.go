package oauthinfoservice

import (
	"context"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

type OAuthInfoService struct {
	platformService	domain.MeetingPlatformService
	repo 			domain.OAuthInfoRepository
}

func (o *OAuthInfoService) CreateOAuthInfo(ctx context.Context, authorization string, userID uint, platform *domain.MeetingPlatform) error {
	token, err := o.platformService.GetOAuthToken(ctx, authorization, platform)
	if err != nil {
		return err
	}

	oauthInfo := &domain.OAuthInfo{
		UserID:            userID,
		MeetingProviderID: platform.ID,
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		TokenType:  	   token.TokenType,
		Expiry: 		   token.Expiry,
	}

	_, err = o.repo.Create(oauthInfo)
	return err
}

func (o *OAuthInfoService) GetOAuthInfo(userID uint, platform *domain.MeetingPlatform) (*domain.OAuthInfo, error) {
	oauthInfos, err := o.repo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	for _, info := range oauthInfos {
		if info.MeetingProviderID == platform.ID {
			token := restoreToken(info)

			if token.Valid() {
				return info, nil
			} else {
				newToken, err := o.platformService.RefreshOAuthToken(nil, &token, platform)
				if err != nil {
					return nil, err
				} else {
					info.AccessToken = newToken.AccessToken
					info.RefreshToken = newToken.RefreshToken
					info.Expiry = newToken.Expiry
					info.TokenType = newToken.TokenType
					err = o.repo.Update(info)
					if err != nil {
						return nil, err
					} else {
						return info, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("the user with id %d does not have OAuth tokens stored for platform with id %d", userID, platform.ID)
}

func restoreToken(oauthInfo *domain.OAuthInfo) oauth2.Token {
	return oauth2.Token{
		AccessToken: oauthInfo.AccessToken,
		RefreshToken: oauthInfo.RefreshToken,
		Expiry: oauthInfo.Expiry,
		TokenType: oauthInfo.TokenType,
	}
}
