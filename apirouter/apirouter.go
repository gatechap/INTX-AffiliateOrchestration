package apirouter

import (
	"net/http"

	"th.truecorp.it.dsm.intcom/affiliateorchestration/apicontrollers"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "th.truecorp.it.dsm.intcom/affiliateorchestration/docs"
)

func SetupAPIRouter(appConfig *config.Config) *gin.Engine {

	apicontrollers.AppConfig = appConfig
	// router := gin.Default()

	router := gin.New()
	router.Use(
		//disable GIN logs from index.html
		gin.LoggerWithWriter(gin.DefaultWriter, "/index.html"),
		gin.Recovery(),
	)

	router.LoadHTMLGlob("web/*.html")

	// redis-int group: root page
	rootpage := router.Group(PATH_ROOT)
	{
		rootpage.GET(PATH_ROOT_INDEX, func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	}

	cpprofile := router.Group(PATH_CONTROLLER_CPEMPLOYEE)
	{
		cpprofile.POST(PATH_MOETHOD_CPPROFILE_PRIMRESOURCE, apicontrollers.CpProfileByPrimResource)
	}

	// swager
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.GET("/swagger-ui", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
