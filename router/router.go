// router package define initial setup of gin-gonic router defining availables routes
// that the microservice serves
package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Gealber/nuitee/client/hotelbeds"
	"github.com/Gealber/nuitee/config"
	rateCtr "github.com/Gealber/nuitee/internal/controller/rate"
	"github.com/Gealber/nuitee/internal/service/rate"
)

func Setup(cfg *config.AppConfig) (*gin.Engine, error) {
	r := gin.Default()

	provider, err := hotelbeds.New(cfg)
	if err != nil {
		return nil, err
	}

	srvs := rate.New(provider)
	// register routes defined in controllers
	rateCtr.New(srvs, r)

	return r, nil
}
