package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/MoneyFlowX/models"
	"github.com/icoder-new/MoneyFlowX/pkg/dto"
	"github.com/icoder-new/MoneyFlowX/utils"
	"net/http"
)

func (h *Handler) GetTransactions(c *gin.Context) {
	query := &dto.TransactionRequestQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("get transaction failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	query = dto.FormatQuery(query)

	user := c.MustGet("user").(*models.User)
	transactions, err := h.service.Transaction.GetTransactions(user.ID, query)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("get transactions failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	totalTransactions, err := h.service.Transaction.CountTransaction(user.ID)
	if err != nil {
		response := utils.ErrorResponse("get transactions failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedTransactions := dto.FormatTransactions(transactions)
	metadata := utils.Metadata{
		Resource: "transactions",
		TotalAll: int(totalTransactions),
		TotalNow: len(transactions),
		Page:     query.Page,
		Limit:    query.Limit,
		Sort:     query.Sort,
	}
	response := utils.ResponseWithPagination("get transactions success", http.StatusOK, formattedTransactions, metadata)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) TopUp(c *gin.Context) {
	input := &dto.TopUpRequestBody{}

	if err := c.ShouldBindJSON(input); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("top up failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("user").(*models.User)
	input.User = user

	transaction, err := h.service.Transaction.TopUp(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("top up failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedTransaction := dto.FormatTopUp(transaction)
	response := utils.SuccessResponse("top up success", http.StatusOK, formattedTransaction)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Transfer(c *gin.Context) {
	input := &dto.TransferRequestBody{}

	if err := c.ShouldBindJSON(input); err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("transfer failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("user").(*models.User)
	input.User = user

	transaction, err := h.service.Transaction.Transfer(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("transfer failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedTransaction := dto.FormatTransfer(transaction)
	response := utils.SuccessResponse("transfer success", http.StatusOK, formattedTransaction)
	c.JSON(http.StatusOK, response)
}
