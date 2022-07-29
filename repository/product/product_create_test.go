package product_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bwmarrin/snowflake"
	"github.com/go-sql-driver/mysql"
	. "github.com/onsi/gomega"

	"jamtangan/domain"
)

func TestProductRepository_Create(t *testing.T) {
	t.Parallel()

	const (
		query = `INSERT INTO product SET id = \?, name = \?, price = \?, brand_id = \?`

		Generate = "Generate"
	)

	var (
		product = domain.Product{
			Name:    "Garmin Instinct 010-02064-94 Tactical Edition Digital Dial Coyote Tan Rubber Strap",
			Price:   4799000,
			BrandID: 1552664566511439872,
		}

		id = snowflake.ID(1552676634333548544)

		args = []driver.Value{
			id, product.Name, product.Price, product.BrandID,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *productTest)
	}{
		{
			name: "prepare error",
			test: func(t *GomegaWithT, b *productTest) {
				createProduct := product

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					WillReturnError(errDatabase)

				err := b.productRepository.Create(context.Background(), &createProduct)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "error",
			test: func(t *GomegaWithT, b *productTest) {
				createProduct := product

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(args...).
					WillReturnError(errDatabase)

				err := b.productRepository.Create(context.Background(), &createProduct)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "duplicate",
			test: func(t *GomegaWithT, b *productTest) {
				createProduct := product

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(args...).
					WillReturnError(&mysql.MySQLError{
						Number: 1452,
					})

				err := b.productRepository.Create(context.Background(), &createProduct)
				t.Expect(errors.Is(err, domain.ErrBadRequest)).Should(BeTrue())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *productTest) {
				createProduct := product

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(args...).
					WillReturnResult(sqlmock.NewResult(id.Int64(), 1))

				err := b.productRepository.Create(context.Background(), &createProduct)
				t.Expect(err).ShouldNot(HaveOccurred())
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
