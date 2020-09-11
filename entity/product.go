package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock hampir habis"
	} else {
		status = "Stock terbatas"
	}
	return status
}
