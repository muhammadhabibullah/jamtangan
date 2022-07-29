package brand_test

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

func TestBrandRepository_GetByID(t *testing.T) {
	t.Parallel()

	const (
		query = `SELECT id, name, created_at, updated_at, is_deleted, deleted_at
			FROM brand 
			WHERE id = \?`
	)

	var (
		now       = time.Now()
		id  int64 = 1552655170888798208

		brand = domain.Brand{
			ID:        id,
			Name:      "CASIO",
			CreatedAt: now,
			UpdatedAt: now,
		}

		rowColumns = []string{
			"id", "name", "created_at", "updated_at", "is_deleted", "deleted_at",
		}
		rowValues = []driver.Value{
			brand.ID,
			brand.Name,
			brand.CreatedAt,
			brand.UpdatedAt,
			brand.IsDeleted,
			brand.DeletedAt,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *brandTest)
	}{
		{
			name: "error",
			test: func(t *GomegaWithT, b *brandTest) {
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnError(errDatabase)

				_, err := b.brandRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "not found",
			test: func(t *GomegaWithT, b *brandTest) {
				rows := sqlmock.NewRows(rowColumns)

				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				_, err := b.brandRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(Equal(domain.ErrNotFound))
			},
		},
		{
			name: "error scan null",
			test: func(t *GomegaWithT, b *brandTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow([]driver.Value{
						nil,
						brand.Name,
						brand.CreatedAt,
						brand.UpdatedAt,
						brand.IsDeleted,
						brand.DeletedAt,
					}...)
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				_, err := b.brandRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(HaveOccurred())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *brandTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow(rowValues...)
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				brandByID, err := b.brandRepository.GetByID(context.Background(), id)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(brandByID).Should(Equal(brand))
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
