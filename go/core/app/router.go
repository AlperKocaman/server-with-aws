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

func NewRouter(dController Controller) Router {
	return router{controller: dController}
}

func NewDefaultRouter() Router {
	return NewRouter(NewDefaultController())
}

func (r router) Register(group *gin.RouterGroup) {
	group.GET("list", r.ListObjects)
	group.POST("put", r.SaveObject)
	group.GET("get/:key", r.GetObject)
}

func (r router) ListObjects(c *gin.Context) {
	log.Println("location: ListObjects")

	c.JSON(response.Generate(r.controller.ListObjects()))
}

func (r router) SaveObject(c *gin.Context) {
	log.Println("location: SaveObject")

	var param SaveObjectParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(response.Generate(nil, response.BadRequest))
		return
	}

	c.JSON(response.Generate(r.controller.SaveObject(param)))
}

func (r router) GetObject(c *gin.Context) {
	log.Println("location: GetObject")

	var param GetObjectParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(response.Generate(nil, response.BadRequest))
		return
	}

	c.JSON(response.Generate(r.controller.GetObject(param)))
}
