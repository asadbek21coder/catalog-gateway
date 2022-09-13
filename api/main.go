package api

import (
	v1 "github.com/asadbek21coder/catalog/gateway/api/handlers/v1"
	"github.com/asadbek21coder/catalog/gateway/config"
	"github.com/asadbek21coder/catalog/gateway/pkg/logger"
	"github.com/asadbek21coder/catalog/gateway/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	_ "github.com/asadbek21coder/catalog/gateway/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))
	// router.Use(MaxAllowed(100))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	// books
	apiV1.GET("/books", handlerV1.GetAll)
	apiV1.GET("/books/:id", handlerV1.GetById)
	apiV1.POST("/books", handlerV1.Create)
	apiV1.PUT("/books", handlerV1.Update)
	apiV1.DELETE("/books/:id", handlerV1.Delete)

	//categories
	apiV1.GET("/categories", handlerV1.GetAllCategories)
	apiV1.GET("/categories/:id", handlerV1.GetCategoryById)
	apiV1.POST("/categories", handlerV1.CreateCategory)
	apiV1.PUT("/categories", handlerV1.UpdateCategory)
	apiV1.DELETE("/categories/:id", handlerV1.DeleteCategory)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
