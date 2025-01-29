package service

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/vagnerlg/deliveryApi/internal/entity"
)

type ProductService struct {
	repo entity.ProductRespository
}

func New(repo entity.ProductRespository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product *entity.Product) {
	s.repo.Create(product)
}

func (s *ProductService) Update(id string, pacth jsonpatch.Patch) entity.Product {
	originaProduct := s.repo.Get(id)
	originalJson, _ := json.Marshal(originaProduct)
	newJson, _ := pacth.Apply(originalJson)
	newProduct := &entity.Product{}
	json.Unmarshal(newJson, newProduct)
	s.repo.Update(newProduct)

	return *newProduct
}

func (s *ProductService) List() []entity.Product {
	return s.repo.List()
}

func (s *ProductService) Get(id string) entity.Product {
	return s.repo.Get(id)
}
