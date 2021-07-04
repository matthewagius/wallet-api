package wallet

import (
	"github.com/matthewagius/wallet-api/entity"
	"github.com/shopspring/decimal"
)

type Reader interface {
	Get(id uint) (*entity.Wallet, error)
}

type Writer interface {
	Update(e *entity.Wallet) error
}

type WalletRepository interface {
	Reader
	Writer
}

type WalletUseCase interface {
	GetWallet(id uint) (*entity.Wallet, error)
	UpdateWallet(walletid uint, transactionType string, amount decimal.Decimal) (*entity.Wallet, error)
}
