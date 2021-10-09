package handle

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sarawoot/payment-gateway/service"
)

type payInternetBankingReq struct {
	Amount float64 `json:"amount" binding:"required"`
}

type PaymentHandle struct {
	srv *service.PaymentService
}

func NewPaymentHandle(srv *service.PaymentService) *PaymentHandle {
	return &PaymentHandle{srv: srv}
}

func (h *PaymentHandle) PayInternetBanking(c *gin.Context) {
	var req payInternetBankingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.PayInternetBanking(req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     res.ID,
		"status": res.Status,
	})
}

func (h *PaymentHandle) PaymentDetail(c *gin.Context) {
	id := strings.Trim(c.Param("id"), " ")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txn, err := h.srv.GetPaymentDetail(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if txn == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     txn.ID,
		"amount": float64(txn.Amount) / 100.0,
		"status": txn.Status,
	})
}
