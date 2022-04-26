package model

type Product struct {
	ProductID string `db:"product_id"`
	Name      string `db:"name"`
	Text      string `db:"text"`
	Picture   string `db:"name"`
}
