package handler

import (
	"net/http"
	"startup/helper"
	"startup/transaction"
	"startup/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Error getting campaign's transactions",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get campaign's transactions",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Campaign's transaction",
		http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get user's transactions",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User's transaction",
		http.StatusOK, "success", transaction.FormatUsersTransaction(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Create transaction failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse(
			"Create transaction failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := transaction.FormatTransaction(newTransaction)
	response := helper.APIResponse(
		"Success create transaction",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
