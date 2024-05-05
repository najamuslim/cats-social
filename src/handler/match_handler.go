package handler

import (
	"cats-social/model/dto"
	"cats-social/src/usecase"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	iMatchUsecase usecase.MatchUsecaseInterface
}

func NewMatchHandler(iMatchUsecase usecase.MatchUsecaseInterface) MatchHandlerInterface {
	return &MatchHandler{iMatchUsecase}
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var request dto.RequestCreateMatch
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	reqUserId,_ := c.Get("user_id")
	err = h.iMatchUsecase.CreateMatch(request, reqUserId.(int))
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully send match request",
	})
}

func (h *MatchHandler) GetMatch(c *gin.Context) {

	userId, _  := c.Get("user_id")
	response, err := h.iMatchUsecase.GetMatch(userId.(int))
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully get match requests",
		"data": response,
	})
}

func (h *MatchHandler) DeleteMatch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
	}

	err = h.iMatchUsecase.GetMatchById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "match not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to check match existence"})
		return
	}

	err = h.iMatchUsecase.DeleteMatch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete match from the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *MatchHandler) ApproveMatch(c *gin.Context) {
	// TODO - add validations
	// matchCatId no longer valid
	
	var request dto.RequestApproveMatch
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = h.iMatchUsecase.GetMatchById(request.MatchId)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "match not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to check match existence"})
		return
	}

	matchCatId, userCatId, err := h.iMatchUsecase.GetCatIdByMatchId(request.MatchId)

	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"status": "not found", "message": "invalid cat id"})
		return
	}

	err = h.iMatchUsecase.ApproveMatch(request.MatchId, matchCatId, userCatId)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully approve match request",
	})
}

func (h *MatchHandler) RejectMatch(c *gin.Context) {
	var request dto.RequestApproveMatch
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = h.iMatchUsecase.GetMatchById(request.MatchId)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "match not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to check match existence"})
		return
	}

	err = h.iMatchUsecase.RejectMatch(request.MatchId)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully reject match request",
	})
}