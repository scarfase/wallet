package wallet

import (
	"testing"

	"github.com/scarfase/wallet/pkg/types"
)

func TestService_FindAccountById_notFound(t *testing.T) {
	accounts := &Service{
		accounts: []*types.Account{
			{
				ID: 123,
			},
		},
	}
	_, err := accounts.FindAccountById(1234)
	if err != nil {
		t.Error(err)
	}
}

func TestService_FindAccountById_Found(t *testing.T) {
	accounts := &Service{
		accounts: []*types.Account{
			{
				ID: 123,
			},
		},
	}
	account, err := accounts.FindAccountById(123)
	if err != nil {
		t.Error(account)
	}
}
