package rate

import (
	"net/http"

	"github.com/Gealber/nuitee/errors"
	"github.com/Gealber/nuitee/internal/service"
	"github.com/Gealber/nuitee/internal/service/model"
	"github.com/gin-gonic/gin"
)

type controller struct {
	rateProvider service.RateProvider
}

func New(rateProvider service.RateProvider, r *gin.Engine) {
	ctr := controller{rateProvider: rateProvider}

	group := r.Group("/hotels")
	group.GET("", ctr.get)
}

// get cheapest rate availability
func (ctr *controller) get(c *gin.Context) {
	var filters model.RateFilters
	err := c.ShouldBindQuery(&filters)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "unable to process filters provided"})
		return
	}

	rates, err := ctr.rateProvider.Get(filters)
	if err != nil {
		code, err := errors.ParseServiceError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rates)
}
