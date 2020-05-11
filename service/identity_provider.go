package service

import (
	sdk "bitbucket.org/accezz-io/api-documentation/go/sdk"
	"context"
	"github.com/antihax/optional"
	"github.com/pkg/errors"
)

type IdentityProviderAPI struct {
	cli *sdk.APIClient
}

func NewIdentityProviderAPI(client *sdk.APIClient) *IdentityProviderAPI {
	return &IdentityProviderAPI{
		cli: client,
	}
}

func (u *IdentityProviderAPI) GetIdentityProviderId(identityProviderName string) (string, error) {
	directoryProviders, _, err := u.cli.IdentityProvidersApi.ListIdentityProviders(context.Background(), &sdk.ListIdentityProvidersOpts{IncludeLocal: optional.NewBool(true)})
	if err != nil {
		return "", err
	}

	for _, directoryProvider := range directoryProviders {
		if directoryProvider.Name == identityProviderName {
			return directoryProvider.Id, nil
		}
	}

	return "", errors.Errorf("can't find identity provider with name '%s'", identityProviderName)
}

func (u *IdentityProviderAPI) GetIdentityProviderTypeById(identityProviderId string) (string, error) {
	directoryProviders, _, err := u.cli.IdentityProvidersApi.ListIdentityProviders(context.Background(), &sdk.ListIdentityProvidersOpts{IncludeLocal: optional.NewBool(true)})
	if err != nil {
		return "", err
	}

	for _, directoryProvider := range directoryProviders {
		if directoryProvider.Id == identityProviderId {
			return directoryProvider.Provider, nil
		}
	}

	return "", errors.Errorf("can't find identity provider with id '%s'", identityProviderId)
}