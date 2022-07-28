package brand

import (
	"context"

	"jamtangan/domain"
)

func (r *brandRepository) GetByID(ctx context.Context, id int64) (domain.Brand, error) {
	const query = `SELECT id, name, created_at, updated_at, is_deleted, deleted_at
			FROM brand 
			WHERE id = ?`

	rows, err := r.sqlDB.QueryContext(ctx, query, id)
	if err != nil {
		return domain.Brand{}, err
	}
	defer rows.Close()

	brands := make([]domain.Brand, 0)
	for rows.Next() {
		var brand domain.Brand
		err = rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.CreatedAt,
			&brand.UpdatedAt,
			&brand.IsDeleted,
			&brand.DeletedAt,
		)
		if err != nil {
			return domain.Brand{}, err
		}

		brands = append(brands, brand)
	}

	var brand domain.Brand
	if len(brands) > 0 {
		brand = brands[0]
	} else {
		err = domain.ErrNotFound
	}

	return brand, err
}
