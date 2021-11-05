package wallet

import (
	"testing"

	"github.com/scarfase/wallet/pkg/types"
)

func TestService_FindAccountByID_notFound(t *testing.T) {
	accounts := &Service{
		accounts: []*types.Account{
			{
				ID: 123,
			},
		},
	}
	_, err := accounts.FindAccountByID(1234)
	if err == nil {
		t.Error(err)
	}
}

func TestService_FindAccountByID_Found(t *testing.T) {
	accounts := &Service{
		accounts: []*types.Account{
			{
				ID: 123,
			},
		},
	}
	account, err := accounts.FindAccountByID(123)
	if err != nil {
		t.Error(account)
	}
}
