package customer_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"jamtangan/domain"
)

func TestCustomerUseCase_CreateTransaction(t *testing.T) {
	t.Parallel()

	const (
		Create  = "Create"
		GetByID = "GetByID"
	)

	var (
		firstTransactionProduct = domain.TransactionProduct{
			ProductID: 1552676634333548544,
			Quantity:  2,
		}
		secondTransactionProduct = domain.TransactionProduct{
			ProductID: 1552688894279946240,
			Quantity:  1,
		}

		firstProduct = domain.Product{
			ID:    1552676634333548544,
			Price: 4799000,
		}
		secondProduct = domain.Product{
			ID:    1552688894279946240,
			Price: 3999000,
		}

		totalPrice  int64 = 13597000
		transaction       = domain.Transaction{
			TotalPrice: totalPrice,
		}

		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *customerTest)
	}{
		{
			name: "get first product by ID error",
			test: func(t *GomegaWithT, c *customerTest) {
				firstTP, secondTP := firstTransactionProduct, secondTransactionProduct
				transactionDetail := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTP,
						&secondTP,
					},
				}
				c.productRepository.On(GetByID, mock.Anything, firstTP.ProductID).
					Return(domain.Product{}, errDatabase)

				err := c.customerUseCase.CreateTransaction(context.Background(), &transactionDetail)
				t.Expect(errors.Is(err, errDatabase)).Should(BeTrue())
			},
		},
		{
			name: "get second product by ID error",
			test: func(t *GomegaWithT, c *customerTest) {
				firstTP, secondTP := firstTransactionProduct, secondTransactionProduct
				transactionDetail := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTP,
						&secondTP,
					},
				}
				c.productRepository.On(GetByID, mock.Anything, firstTP.ProductID).
					Return(firstProduct, nil)
				c.productRepository.On(GetByID, mock.Anything, secondTP.ProductID).
					Return(domain.Product{}, errDatabase)

				err := c.customerUseCase.CreateTransaction(context.Background(), &transactionDetail)
				t.Expect(errors.Is(err, errDatabase)).Should(BeTrue())
			},
		},
		{
			name: "create transaction error",
			test: func(t *GomegaWithT, c *customerTest) {
				firstTP, secondTP := firstTransactionProduct, secondTransactionProduct
				transactionDetail := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTP,
						&secondTP,
					},
				}

				c.productRepository.On(GetByID, mock.Anything, firstTP.ProductID).
					Return(firstProduct, nil)
				c.productRepository.On(GetByID, mock.Anything, secondTP.ProductID).
					Return(secondProduct, nil)

				createTransaction := transaction
				createFirstTransactionProduct := firstTP
				createFirstTransactionProduct.Price = firstProduct.Price
				createSecondTransactionProduct := secondTP
				createSecondTransactionProduct.Price = secondProduct.Price
				createTransactionProducts := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				c.transactionRepository.On(Create, mock.Anything, &createTransaction, createTransactionProducts).
					Return(errDatabase)

				err := c.customerUseCase.CreateTransaction(context.Background(), &transactionDetail)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "get by id error",
			test: func(t *GomegaWithT, c *customerTest) {
				firstTP, secondTP := firstTransactionProduct, secondTransactionProduct
				transactionDetail := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTP,
						&secondTP,
					},
				}

				c.productRepository.On(GetByID, mock.Anything, firstTP.ProductID).
					Return(firstProduct, nil)
				c.productRepository.On(GetByID, mock.Anything, secondTP.ProductID).
					Return(secondProduct, nil)

				createTransaction := transaction
				createFirstTransactionProduct := firstTP
				createFirstTransactionProduct.Price = firstProduct.Price
				createSecondTransactionProduct := secondTP
				createSecondTransactionProduct.Price = secondProduct.Price
				createTransactionProducts := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				c.transactionRepository.On(Create, mock.Anything, &createTransaction, createTransactionProducts).
					Return(nil)

				c.transactionRepository.On(GetByID, mock.Anything, transaction.ID).
					Return(domain.Transaction{}, nil, errDatabase)

				err := c.customerUseCase.CreateTransaction(context.Background(), &transactionDetail)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, c *customerTest) {
				firstTP, secondTP := firstTransactionProduct, secondTransactionProduct
				transactionDetail := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTP,
						&secondTP,
					},
				}

				c.productRepository.On(GetByID, mock.Anything, firstTP.ProductID).
					Return(firstProduct, nil)
				c.productRepository.On(GetByID, mock.Anything, secondTP.ProductID).
					Return(secondProduct, nil)

				createTransaction := transaction
				createFirstTransactionProduct := firstTP
				createFirstTransactionProduct.Price = firstProduct.Price
				createSecondTransactionProduct := secondTP
				createSecondTransactionProduct.Price = secondProduct.Price
				createTransactionProducts := []*domain.TransactionProduct{
					&createFirstTransactionProduct,
					&createSecondTransactionProduct,
				}

				c.transactionRepository.On(Create, mock.Anything, &createTransaction, createTransactionProducts).
					Return(nil)
				c.transactionRepository.On(GetByID, mock.Anything, transaction.ID).
					Return(createTransaction, []domain.TransactionProduct{
						createFirstTransactionProduct,
						createSecondTransactionProduct,
					}, nil)

				err := c.customerUseCase.CreateTransaction(context.Background(), &transactionDetail)
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
