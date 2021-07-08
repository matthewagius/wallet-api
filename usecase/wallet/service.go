package wallet

import (
	"time"

	"github.com/matthewagius/wallet-api/entity"
	"github.com/shopspring/decimal"
)

type Service struct {
	repo WalletRepository
}

func NewService(r WalletRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetWallet(id uint) (*entity.Wallet, error) {
	data, err := s.repo.Get(id)

	if data == nil {
		return nil, entity.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Service) UpdateWallet(walletid uint, transactionType string, amount decimal.Decimal) (*entity.Wallet, error) {
	data, err := s.GetWallet(walletid)

	if err != nil {
		return nil, err
	}

	e, err := entity.Initialise(data.Id, data.UserId, data.CurrencyCode, amount, transactionType, data.CreatedAt)

	if err != nil {
		return nil, err
	}

	switch e.TransactionType {
	case "Credit":
		e.Amount = data.Amount.Add(e.Amount)
	case "Debit":
		err := entity.ValidateDebitTransaction(amount, data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, entity.ErrInvalidTransactionType
	}
	e.UpdatedAt = time.Now()
	return e, s.repo.Update(e)
}
