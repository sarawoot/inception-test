package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/omise/omise-go"
	"github.com/sarawoot/payment-gateway/handle"
	"github.com/sarawoot/payment-gateway/repo"
	"github.com/sarawoot/payment-gateway/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	OmisePublicKey = "pkey_test_5mci4nahnjn8p8byryi"
	OmiseSecretKey = "skey_test_5mci4naho2s633rwwmd"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./payment.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	omiseClient, err := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if err != nil {
		log.Fatal(err)
	}

	repoTxn := repo.NewPaymentTranctionRepo(db)
	srv := service.New(omiseClient, repoTxn)
	h := handle.NewPaymentHandle(srv)

	route := gin.Default()
	route.POST("/payment/internet-banking-scb", h.PayInternetBanking)
	route.GET("/payment/:id", h.PaymentDetail)
	if err := route.Run(); err != nil {
		log.Fatal(err)
	}

}
