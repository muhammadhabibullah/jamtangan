package transaction_test

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

func TestTransactionRepository_Create(t *testing.T) {
	t.Parallel()

	const (
		transactionQuery        = `INSERT INTO transaction SET id = \?, total_price = \?`
		transactionProductQuery = `INSERT INTO transaction_product 
			SET transaction_id = \?, product_id = \?, quantity = \?, price = \?`

		Generate = "Generate"
	)

	var (
		transaction = domain.Transaction{
			TotalPrice: 13597000,
		}
		firstTransactionProduct = domain.TransactionProduct{
			ProductID: 1552676634333548544,
			Quantity:  1,
			Price:     4799000,
		}
		secondTransactionProduct = domain.TransactionProduct{
			ProductID: 1552688894279946240,
			Quantity:  2,
			Price:     3999000,
		}

		transactionID = snowflake.ID(1552855546665635840)

		transactionArgs = []driver.Value{
			transactionID, transaction.TotalPrice,
		}
		firstTransactionProductArgs = []driver.Value{
			transactionID, firstTransactionProduct.ProductID, firstTransactionProduct.Quantity, firstTransactionProduct.Price,
		}
		secondTransactionProductArgs = []driver.Value{
			transactionID, secondTransactionProduct.ProductID, secondTransactionProduct.Quantity, secondTransactionProduct.Price,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *transactionTest)
	}{
		{
			name: "begin tx error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin().
					WillReturnError(errDatabase)

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "prepare transaction error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					WillReturnError(errDatabase)
				b.mock.ExpectRollback()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "exec transaction error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnError(errDatabase)
				b.mock.ExpectRollback()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "prepare transaction product error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				b.mock.ExpectPrepare(transactionProductQuery).
					WillReturnError(errDatabase)
				b.mock.ExpectRollback()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "exec first transaction product error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				b.mock.ExpectPrepare(transactionProductQuery).
					ExpectExec().
					WithArgs(firstTransactionProductArgs...).
					WillReturnError(errDatabase)
				b.mock.ExpectRollback()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "exec second transaction product error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep := b.mock.ExpectPrepare(transactionProductQuery)
				prep.ExpectExec().
					WithArgs(firstTransactionProductArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep.ExpectExec().
					WithArgs(secondTransactionProductArgs...).
					WillReturnError(errDatabase)
				b.mock.ExpectRollback()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "exec second transaction product duplicate error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := firstTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep := b.mock.ExpectPrepare(transactionProductQuery)
				prep.ExpectExec().
					WithArgs(firstTransactionProductArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep.ExpectExec().
					WithArgs(firstTransactionProductArgs...).
					WillReturnError(&mysql.MySQLError{
						Number: 1062,
					})
				b.mock.ExpectRollback()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(errors.Is(err, domain.ErrBadRequest)).Should(BeTrue())
			},
		},
		{
			name: "commit error",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep := b.mock.ExpectPrepare(transactionProductQuery)
				prep.ExpectExec().
					WithArgs(firstTransactionProductArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep.ExpectExec().
					WithArgs(secondTransactionProductArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				b.mock.ExpectCommit().
					WillReturnError(errDatabase)

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, b *transactionTest) {
				createTransaction := transaction
				createFirstTransactionProduct := firstTransactionProduct
				createSecondTransactionProduct := secondTransactionProduct
				createTransactionProduct := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				td := domain.TransactionDetail{
					Transaction:         &createTransaction,
					TransactionProducts: createTransactionProduct,
				}

				b.snowflakeMock.On(Generate).
					Return(transactionID)
				b.mock.ExpectBegin()
				b.mock.ExpectPrepare(transactionQuery).
					ExpectExec().
					WithArgs(transactionArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep := b.mock.ExpectPrepare(transactionProductQuery)
				prep.ExpectExec().
					WithArgs(firstTransactionProductArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				prep.ExpectExec().
					WithArgs(secondTransactionProductArgs...).
					WillReturnResult(sqlmock.NewResult(transaction.ID, 1))
				b.mock.ExpectCommit()

				err := b.transactionRepository.Create(context.Background(), td.Transaction, td.TransactionProducts)
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
