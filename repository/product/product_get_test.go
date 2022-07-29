package product_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/gomega"

	"jamtangan/domain"
)

func TestProductRepository_GetByID(t *testing.T) {
	t.Parallel()

	const (
		query = `SELECT id, name, price, brand_id, created_at, updated_at, is_deleted, deleted_at
			FROM product 
			WHERE id = \?`
	)

	var (
		now       = time.Now()
		id  int64 = 1552676634333548544

		product = domain.Product{
			ID:        id,
			Name:      "Garmin Instinct 010-02064-94 Tactical Edition Digital Dial Coyote Tan Rubber Strap",
			Price:     4799000,
			BrandID:   1552664566511439872,
			CreatedAt: now,
			UpdatedAt: now,
		}

		rowColumns = []string{
			"id", "name", "price", "brand_id", "created_at", "updated_at", "is_deleted", "deleted_at",
		}
		rowValues = []driver.Value{
			product.ID,
			product.Name,
			product.Price,
			product.BrandID,
			product.CreatedAt,
			product.UpdatedAt,
			product.IsDeleted,
			product.DeletedAt,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *productTest)
	}{
		{
			name: "error",
			test: func(t *GomegaWithT, b *productTest) {
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnError(errDatabase)

				_, err := b.productRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "not found",
			test: func(t *GomegaWithT, b *productTest) {
				rows := sqlmock.NewRows(rowColumns)

				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				_, err := b.productRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(Equal(domain.ErrNotFound))
			},
		},
		{
			name: "error scan null",
			test: func(t *GomegaWithT, b *productTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow([]driver.Value{
						nil,
						product.Name,
						product.Price,
						product.BrandID,
						product.CreatedAt,
						product.UpdatedAt,
						product.IsDeleted,
						product.DeletedAt,
					}...)
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				_, err := b.productRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(HaveOccurred())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *productTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow(rowValues...)
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				productByID, err := b.productRepository.GetByID(context.Background(), id)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(productByID).Should(Equal(product))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test(t, tt.test)
		})
	}
}

func TestProductRepository_FetchByBrandID(t *testing.T) {
	t.Parallel()

	const (
		query = `SELECT id, name, price, brand_id, created_at, updated_at, is_deleted, deleted_at
			FROM product 
			WHERE brand_id = \?`
	)

	var (
		now           = time.Now()
		brandID int64 = 1552664566511439872

		product = domain.Product{
			ID:        brandID,
			Name:      "Garmin Instinct 010-02064-94 Tactical Edition Digital Dial Coyote Tan Rubber Strap",
			Price:     4799000,
			BrandID:   1552664566511439872,
			CreatedAt: now,
			UpdatedAt: now,
		}

		rowColumns = []string{
			"id", "name", "price", "brand_id", "created_at", "updated_at", "is_deleted", "deleted_at",
		}
		rowValues = []driver.Value{
			product.ID,
			product.Name,
			product.Price,
			product.BrandID,
			product.CreatedAt,
			product.UpdatedAt,
			product.IsDeleted,
			product.DeletedAt,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *productTest)
	}{
		{
			name: "error",
			test: func(t *GomegaWithT, b *productTest) {
				b.mock.ExpectQuery(query).
					WithArgs(brandID).
					WillReturnError(errDatabase)

				_, err := b.productRepository.FetchByBrandID(context.Background(), brandID)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "not found",
			test: func(t *GomegaWithT, b *productTest) {
				rows := sqlmock.NewRows(rowColumns)

				b.mock.ExpectQuery(query).
					WithArgs(brandID).
					WillReturnRows(rows)

				productsByBrandID, err := b.productRepository.FetchByBrandID(context.Background(), brandID)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(productsByBrandID).Should(BeEmpty())
			},
		},
		{
			name: "error scan null",
			test: func(t *GomegaWithT, b *productTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow([]driver.Value{
						nil,
						product.Name,
						product.Price,
						product.BrandID,
						product.CreatedAt,
						product.UpdatedAt,
						product.IsDeleted,
						product.DeletedAt,
					}...)
				b.mock.ExpectQuery(query).
					WithArgs(brandID).
					WillReturnRows(rows)

				_, err := b.productRepository.FetchByBrandID(context.Background(), brandID)
				t.Expect(err).Should(HaveOccurred())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *productTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow(rowValues...)
				b.mock.ExpectQuery(query).
					WithArgs(brandID).
					WillReturnRows(rows)

				productsByBrandID, err := b.productRepository.FetchByBrandID(context.Background(), brandID)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(productsByBrandID).Should(Equal([]domain.Product{product}))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test(t, tt.test)
		})
	}
}
