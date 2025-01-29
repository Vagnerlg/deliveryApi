package httpApi

import (
	"io"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/deliveryApi/internal/entity"
	"github.com/vagnerlg/deliveryApi/internal/service"
)

type http struct {
	service service.ProductService
}

func New(service service.ProductService) *http {
	return &http{service: service}
}

func (h *http) Create(ctx *gin.Context) {
	var product = &entity.Product{}
	error := ctx.ShouldBindBodyWithJSON(&product)
	if error != nil {
		ctx.JSON(400, gin.H{"msg": error.Error()})
		return
	}

	h.service.Create(product)

	ctx.JSON(200, product)
}

func (h *http) List(ctx *gin.Context) {
	ctx.JSON(200, h.service.List())
}

func (h *http) Get(ctx *gin.Context) {
	prod := h.service.Get(ctx.Param("id"))

	ctx.JSON(200, gin.H{"data": prod})
}

func (h *http) Patch(ctx *gin.Context) {
	body := ctx.Request.Body
	data, _ := io.ReadAll(body)
	pacth, _ := jsonpatch.DecodePatch(data)

	ctx.JSON(200, gin.H{"data": h.service.Update(ctx.Param("id"), pacth)})
}
