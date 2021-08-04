package repositorys

import (
	"database/sql"
	"study/datamodels"

	_ "github.com/go-sql-driver/mysql"
)

type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Update(*datamodels.Product) error
	Delete(int64) bool
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProduceManager struct {
	table    string
	mydbconn *sql.DB
}

func NewProduceManager(table string, dbconn *sql.DB) IProduct {
	return &ProduceManager{
		table:    table,
		mydbconn: dbconn,
	}

}

func (p *ProduceManager) Conn() (err error) {
	return
}

func (p *ProduceManager) Insert(*datamodels.Product) (id int64, err error) {
	return
}

func (p *ProduceManager) Update(*datamodels.Product) (err error) {
	return
}

func (p *ProduceManager) Delete(int64) bool {
	return true
}

func (p *ProduceManager) SelectByKey(int64) (item *datamodels.Product, err error) {
	return
}

func (p *ProduceManager) SelectAll() (items []*datamodels.Product, err error) {
	return
}
