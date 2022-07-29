package brand_test

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

func TestBrandRepository_Create(t *testing.T) {
	t.Parallel()

	const (
		query = `INSERT INTO brand SET id = \?, name = \?`

		Generate = "Generate"
	)

	var (
		brand = domain.Brand{
			Name: "CASIO",
		}

		id = snowflake.ID(1552655170888798208)

		args = []driver.Value{
			id, brand.Name,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *brandTest)
	}{
		{
			name: "prepare error",
			test: func(t *GomegaWithT, b *brandTest) {
				createBrand := brand

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					WillReturnError(errDatabase)

				err := b.brandRepository.Create(context.Background(), &createBrand)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "error",
			test: func(t *GomegaWithT, b *brandTest) {
				createBrand := brand

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(args...).
					WillReturnError(errDatabase)

				err := b.brandRepository.Create(context.Background(), &createBrand)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "duplicate",
			test: func(t *GomegaWithT, b *brandTest) {
				createBrand := brand

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(args...).
					WillReturnError(&mysql.MySQLError{
						Number: 1062,
					})

				err := b.brandRepository.Create(context.Background(), &createBrand)
				t.Expect(errors.Is(err, domain.ErrDuplicate)).Should(BeTrue())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *brandTest) {
				createBrand := brand

				b.snowflakeMock.On(Generate).
					Return(id)
				b.mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(args...).
					WillReturnResult(sqlmock.NewResult(id.Int64(), 1))

				err := b.brandRepository.Create(context.Background(), &createBrand)
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
