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

func TestCustomerUseCase_GetTransactionByID(t *testing.T) {
	t.Parallel()

	const (
		GetByID = "GetByID"
	)

	var (
		id int64 = 1552855546665635840

		transaction = domain.Transaction{
			ID: id,
		}

		firstTransactionProduct = domain.TransactionProduct{
			TransactionID: id,
			ProductID:     1552676634333548544,
			Quantity:      2,
			Price:         4799000,
		}
		secondTransactionProduct = domain.TransactionProduct{
			TransactionID: id,
			ProductID:     1552688894279946240,
			Quantity:      1,
			Price:         3999000,
		}

		transactionProducts = []domain.TransactionProduct{
			firstTransactionProduct,
			secondTransactionProduct,
		}

		transactionDetail = domain.TransactionDetail{
			Transaction: &transaction,
			TransactionProducts: []*domain.TransactionProduct{
				&firstTransactionProduct,
				&secondTransactionProduct,
			},
		}

		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *customerTest)
	}{
		{
			name: "error",
			test: func(t *GomegaWithT, a *customerTest) {
				a.transactionRepository.On(GetByID, mock.Anything, id).
					Return(domain.Transaction{}, nil, errDatabase)

				_, err := a.customerUseCase.GetTransactionByID(context.Background(), id)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "not found",
			test: func(t *GomegaWithT, a *customerTest) {
				a.transactionRepository.On(GetByID, mock.Anything, id).
					Return(domain.Transaction{}, []domain.TransactionProduct{}, nil)

				_, err := a.customerUseCase.GetTransactionByID(context.Background(), id)
				t.Expect(errors.Is(err, domain.ErrNotFound)).Should(BeTrue())
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, a *customerTest) {
				a.transactionRepository.On(GetByID, mock.Anything, id).
					Return(transaction, transactionProducts, nil)

				transactionDetailByID, err := a.customerUseCase.GetTransactionByID(context.Background(), id)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(transactionDetailByID).Should(Equal(transactionDetail))
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
