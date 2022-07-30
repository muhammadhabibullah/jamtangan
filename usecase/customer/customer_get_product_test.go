package customer_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"jamtangan/domain"
)

func TestCustomerUseCase_FetchProductByBrandID(t *testing.T) {
	t.Parallel()

	const (
		FetchByBrandID = "FetchByBrandID"
	)

	var (
		brandID int64 = 1552664566511439872

		products = []domain.Product{
			{
				ID:      1552688894279946240,
				Name:    "Garmin Instinct 010-02064-64 Seafoam Digital Dial Blue Rubber Strap",
				Price:   3999000,
				BrandID: brandID,
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
				a.productRepository.On(FetchByBrandID, mock.Anything, brandID).
					Return(nil, errDatabase)

				_, err := a.customerUseCase.FetchProductByBrandID(context.Background(), brandID)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, a *customerTest) {
				a.productRepository.On(FetchByBrandID, mock.Anything, brandID).
					Return(products, nil)

				productByID, err := a.customerUseCase.FetchProductByBrandID(context.Background(), brandID)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(productByID).Should(Equal(products))
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
