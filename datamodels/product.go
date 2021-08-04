package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"id" imoc:"id"`
	ProductName  string `json:"productname" sql:"productname" imoc:"productname"`
	ProductImage string `json:"productImage" sql:"productImage" imoc:"productImage"`
	ProductUrl   string `json:"producturl" sql:"producturl" imoc:"producturl"`
}
