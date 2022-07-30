package admin_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"jamtangan/domain"
)

func TestAdminUseCase_GetProductByID(t *testing.T) {
	t.Parallel()

	const (
		GetByID = "GetByID"
	)

	var (
		id int64 = 1552688894279946240

		product = domain.Product{
			ID:      id,
			Name:    "Garmin Instinct 010-02064-64 Seafoam Digital Dial Blue Rubber Strap",
			Price:   3999000,
			BrandID: 1552664566511439872,
		}

		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *adminTest)
	}{
		{
			name: "error",
			test: func(t *GomegaWithT, a *adminTest) {
				a.productRepository.On(GetByID, mock.Anything, id).
					Return(domain.Product{}, errDatabase)

				_, err := a.adminUseCase.GetProductByID(context.Background(), id)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, a *adminTest) {
				a.productRepository.On(GetByID, mock.Anything, id).
					Return(product, nil)

				productByID, err := a.adminUseCase.GetProductByID(context.Background(), id)
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
