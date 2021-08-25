package database

import "github.com/cfabrica46/search-engine/structure"

func Search(query string) (products []structure.Product, err error) {

	rows, err := db.Query("SELECT products.id, products.name, products.type, products.brand, products.price, products.stock, products.date FROM products WHERE to_tsvector('english',description) @@ to_tsquery('english',$1)", query)
	if err != nil {
		return
	}

	for rows.Next() {
		var productAux structure.Product
		err = rows.Scan(&productAux.ID, &productAux.Name, &productAux.Type, &productAux.Brand, &productAux.Price, &productAux.Stock, &productAux.Date)
		if err != nil {
			return
		}
		products = append(products, productAux)
	}
	return
}
