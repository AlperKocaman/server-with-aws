package app

import (
	"github.com/AlperKocaman/server-with-aws/core/response"
	"github.com/gin-gonic/gin"
	"log"
)

type Router interface {
	Register(group *gin.RouterGroup)
	ListObjects(c *gin.Context)
	SaveObject(c *gin.Context)
	GetObject(c *gin.Context)
}

type router struct {
	controller Controller
}

func NewRouter(controller Controller) Router {
	return router{controller: controller}
}

func NewDefaultRouter() Router {
	return NewRouter(NewDefaultController())
}

func (r router) Register(group *gin.RouterGroup) {
	group.GET("list", r.ListObjects)
	group.POST("put", r.SaveObject)
	group.GET("get", r.GetObject)
}

func (r router) ListObjects(c *gin.Context) {
	log.Println("location: ListObjectsRouter")

	c.JSON(response.Generate(r.controller.ListObjects()))
}

func (r router) SaveObject(c *gin.Context) {
	log.Println("location: SaveObjectRouter")

	var param SaveObjectParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(response.Generate(nil, response.BadRequest))
		return
	}

	c.JSON(response.Generate(r.controller.SaveObject(param)))
}

func (r router) GetObject(c *gin.Context) {
	log.Println("location: GetObjectRouter")

	var param GetObjectParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(response.Generate(nil, response.BadRequest))
		return
	}

	c.JSON(response.Generate(r.controller.GetObject(param)))
}
