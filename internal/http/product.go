package httpApi

import (
	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/deliveryApi/internal/entity"
)

type http struct {
	repo entity.ProductRespository
}

func New(repo entity.ProductRespository) *http {
	return &http{repo: repo}
}

func (h *http) Create(ctx *gin.Context) {
	var product = &entity.Product{}
	error := ctx.ShouldBindBodyWithJSON(&product)
	if error != nil {
		ctx.JSON(400, gin.H{"msg": error.Error()})
		return
	}

	h.repo.Create(product)

	ctx.JSON(200, product)
}

func (h *http) List(ctx *gin.Context) {
	ctx.JSON(200, h.repo.List())
}

func (h *http) Get(ctx *gin.Context) {
	prod := h.repo.Get(ctx.Param("id"))

	ctx.JSON(200, gin.H{"data": prod})
}
