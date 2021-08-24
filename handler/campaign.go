package handler

import (
	"net/http"
	"startup/campaign"
	"startup/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service menentukan repository mana yang dipakai
// repository akses ke db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

//api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse(
			"Error getting campaigns",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse(
		"List of campaigns",
		http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
