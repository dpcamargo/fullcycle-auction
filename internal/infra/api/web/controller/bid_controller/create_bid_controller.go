package bid_controller

import (
	"context"
	"net/http"

	"github.com/dpcamargo/fullcycle-auction/configuration/rest_err"
	"github.com/dpcamargo/fullcycle-auction/internal/infra/api/web/validation"
	"github.com/dpcamargo/fullcycle-auction/internal/usecase/bid_usecase"
	"github.com/gin-gonic/gin"
)

type BidController struct {
	BidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(BidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		BidUseCase: BidUseCase,
	}
}

func (u *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.BidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
