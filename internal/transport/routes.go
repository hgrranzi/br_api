package transport

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(braineeHandler *BraineeHandler) *gin.Engine {
	router := gin.Default()

	braineeRoutes := router.Group("/brainees")
	{
		braineeRoutes.POST("", braineeHandler.CreateBrainee)
		braineeRoutes.GET("", braineeHandler.GetAllBrainees)
		braineeRoutes.GET("/:braineeId", braineeHandler.GetBraineeById)
	}

	return router
}
