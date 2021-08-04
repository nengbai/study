package services

import (
	"study/datamodels"
	"study/repositorys"
)

type ProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeletProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
}

type ProduceService struct {
	produceRepository repositorys.IProduct
}

func NewProduceService(repository repositorys.IProduct) ProduceService {
	return ProduceService{produceRepository: repository}
}

func (s *ProduceService) GetProductByID(prdid int64) (*datamodels.Product, error) {
	return s.produceRepository.SelectByKey(prdid)
}

func (s *ProduceService) GetAllProduct() (items []*datamodels.Product, err error) {
	return s.produceRepository.SelectAll()
}

func (s *ProduceService) DeletProductByID(id int64) bool {
	return s.produceRepository.Delete(id)
}
func (s *ProduceService) InsertProduct(product *datamodels.Product) (int64, error) {
	return s.produceRepository.Insert(product)
}

func (s *ProduceService) UpdateProduct(product *datamodels.Product) (err error) {
	return s.produceRepository.Update(product)
}
