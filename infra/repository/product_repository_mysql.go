package repository

import (
	"database/sql"
	"github.com/gokafkamysql/entity"
)

type ProductMySqlRepository struct {
	Db *sql.DB
}

func (c ProductMySqlRepository) Insert(product entity.Product) error {

	stmt, err := c.Db.Prepare("Insert into product(id,name,quantidade) values(?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		product.ID,
		product.Name,
		product.Quantidade,
	)
	if err != nil {
		return err
	}

	return nil
}
