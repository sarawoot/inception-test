package repo

import (
	"errors"

	"github.com/sarawoot/payment-gateway/model"
	"gorm.io/gorm"
)

type PaymentTranctionRepo struct {
	db *gorm.DB
}

func NewPaymentTranctionRepo(db *gorm.DB) *PaymentTranctionRepo {
	return &PaymentTranctionRepo{
		db: db,
	}
}

func (r *PaymentTranctionRepo) Create(chargeID, sourceID string, amount int64, status, RawRespone string) (*model.PaymentTransaction, error) {
	txn := model.PaymentTransaction{
		ChargeID:   chargeID,
		SourceID:   sourceID,
		Amount:     amount,
		RawRespone: RawRespone,
		Status:     status,
	}
	res := r.db.Create(&txn)

	return &txn, res.Error
}

func (r *PaymentTranctionRepo) GetByID(id uint64) (*model.PaymentTransaction, error) {
	var txn model.PaymentTransaction
	err := r.db.First(&txn, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &txn, err
}
