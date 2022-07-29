package transaction_test

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

func TestTransactionRepository_GetByID(t *testing.T) {
	t.Parallel()

	const (
		query = `SELECT t.id, t.total_price, t.created_at, t.updated_at, t.is_deleted, t.deleted_at,
       		tp.transaction_id, tp.product_id, tp.quantity, tp.price, tp.created_at, tp.updated_at, tp.is_deleted, tp.deleted_at
			FROM transaction t
			INNER JOIN transaction_product tp on t.id = tp.transaction_id
			WHERE id = \?`
	)

	var (
		now       = time.Now()
		id  int64 = 1552855546665635840

		transaction = domain.Transaction{
			ID:         id,
			TotalPrice: 13597000,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		firstTransactionProduct = domain.TransactionProduct{
			TransactionID: id,
			ProductID:     1552676634333548544,
			Quantity:      1,
			Price:         4799000,
		}
		secondTransactionProduct = domain.TransactionProduct{
			TransactionID: id,
			ProductID:     1552688894279946240,
			Quantity:      2,
			Price:         3999000,
		}
		transactionProducts = []domain.TransactionProduct{
			firstTransactionProduct,
			secondTransactionProduct,
		}

		rowColumns = []string{
			"id", "total_price", "t.created_at", "t.updated_at", "t.is_deleted", "t.deleted_at",
			"transaction_id", "product_id", "quantity", "price",
			"tp.created_at", "tp.updated_at", "tp.is_deleted", "tp.deleted_at",
		}
		firstRowValues = []driver.Value{
			transaction.ID,
			transaction.TotalPrice,
			transaction.CreatedAt,
			transaction.UpdatedAt,
			transaction.IsDeleted,
			transaction.DeletedAt,
			firstTransactionProduct.TransactionID,
			firstTransactionProduct.ProductID,
			firstTransactionProduct.Quantity,
			firstTransactionProduct.Price,
			firstTransactionProduct.CreatedAt,
			firstTransactionProduct.UpdatedAt,
			firstTransactionProduct.IsDeleted,
			firstTransactionProduct.DeletedAt,
		}
		secondRowValues = []driver.Value{
			transaction.ID,
			transaction.TotalPrice,
			transaction.CreatedAt,
			transaction.UpdatedAt,
			transaction.IsDeleted,
			transaction.DeletedAt,
			secondTransactionProduct.TransactionID,
			secondTransactionProduct.ProductID,
			secondTransactionProduct.Quantity,
			secondTransactionProduct.Price,
			secondTransactionProduct.CreatedAt,
			secondTransactionProduct.UpdatedAt,
			secondTransactionProduct.IsDeleted,
			secondTransactionProduct.DeletedAt,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *transactionTest)
	}{
		{
			name: "error",
			test: func(t *GomegaWithT, b *transactionTest) {
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnError(errDatabase)

				_, _, err := b.transactionRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "not found",
			test: func(t *GomegaWithT, b *transactionTest) {
				rows := sqlmock.NewRows(rowColumns)

				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				_, transactionProducts, err := b.transactionRepository.GetByID(context.Background(), id)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(transactionProducts).Should(BeEmpty())
			},
		},
		{
			name: "error scan null",
			test: func(t *GomegaWithT, b *transactionTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow([]driver.Value{
						nil,
						transaction.TotalPrice,
						transaction.CreatedAt,
						transaction.UpdatedAt,
						transaction.IsDeleted,
						transaction.DeletedAt,
						firstTransactionProduct.TransactionID,
						firstTransactionProduct.ProductID,
						firstTransactionProduct.Quantity,
						firstTransactionProduct.Price,
						firstTransactionProduct.CreatedAt,
						firstTransactionProduct.UpdatedAt,
						firstTransactionProduct.IsDeleted,
						firstTransactionProduct.DeletedAt,
					}...)
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				_, _, err := b.transactionRepository.GetByID(context.Background(), id)
				t.Expect(err).Should(HaveOccurred())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *transactionTest) {
				rows := sqlmock.NewRows(rowColumns).
					AddRow(firstRowValues...).
					AddRow(secondRowValues...)
				b.mock.ExpectQuery(query).
					WithArgs(id).
					WillReturnRows(rows)

				transactionByID, transactionProductsByID, err := b.transactionRepository.GetByID(context.Background(), id)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(transactionByID).Should(Equal(transaction))
				t.Expect(transactionProductsByID).Should(Equal(transactionProducts))
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
