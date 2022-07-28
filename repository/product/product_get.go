package product

import (
	"context"

	"jamtangan/domain"
)

func (r *productRepository) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	const query = `SELECT id, name, price, brand_id, created_at, updated_at, is_deleted, deleted_at
			FROM product 
			WHERE id = ?`

	rows, err := r.sqlDB.QueryContext(ctx, query, id)
	if err != nil {
		return domain.Product{}, err
	}
	defer rows.Close()

	products := make([]domain.Product, 0)
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.BrandID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.IsDeleted,
			&product.DeletedAt,
		)
		if err != nil {
			return domain.Product{}, err
		}

		products = append(products, product)
	}

	var product domain.Product
	if len(products) > 0 {
		product = products[0]
	} else {
		err = domain.ErrNotFound
	}

	return product, err
}

func (r *productRepository) FetchByBrandID(ctx context.Context, brandID int64) ([]domain.Product, error) {
	const query = `SELECT id, name, price, brand_id, created_at, updated_at, is_deleted, deleted_at
			FROM product 
			WHERE brand_id = ?`

	rows, err := r.sqlDB.QueryContext(ctx, query, brandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]domain.Product, 0)
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.BrandID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.IsDeleted,
			&product.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, err
}
