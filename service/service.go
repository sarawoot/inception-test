package service

import (
	"encoding/json"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/sarawoot/payment-gateway/model"
	"github.com/sarawoot/payment-gateway/repo"
)

type PaymentService struct {
	omiseClient *omise.Client
	txnRepo     *repo.PaymentTranctionRepo
}

func New(omiseClient *omise.Client, txnRepo *repo.PaymentTranctionRepo) *PaymentService {
	return &PaymentService{
		omiseClient: omiseClient,
		txnRepo:     txnRepo,
	}
}

func (s *PaymentService) PayInternetBanking(amount float64) (*model.PaymentTransaction, error) {
	scbType := "internet_banking_scb"
	currency := "thb"
	returnURI := "http://www.example.com/"
	amt := int64(amount * 100)

	var source omise.Source
	createSource := operations.CreateSource{
		Type:     scbType,
		Currency: currency,
		Amount:   amt,
	}
	if err := s.omiseClient.Do(&source, &createSource); err != nil {
		return nil, err
	}

	var charge omise.Charge
	createCharge := operations.CreateCharge{
		Source:    source.ID,
		Amount:    source.Amount,
		Currency:  currency,
		ReturnURI: returnURI,
	}

	if err := s.omiseClient.Do(&charge, &createCharge); err != nil {
		return nil, err
	}

	chargeByte, err := json.Marshal(charge)
	if err != nil {
		return nil, err
	}

	return s.txnRepo.Create(charge.ID, source.ID, amt, string(charge.Status), string(chargeByte))
}

func (s *PaymentService) GetPaymentDetail(id uint64) (*model.PaymentTransaction, error) {
	return s.txnRepo.GetByID(id)
}
