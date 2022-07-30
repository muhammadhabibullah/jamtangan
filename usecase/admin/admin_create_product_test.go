package admin_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"jamtangan/domain"
)

func TestAdminUseCase_CreateProduct(t *testing.T) {
	t.Parallel()

	const (
		Create  = "Create"
		GetByID = "GetByID"
	)

	var (
		productName          = "Garmin Instinct 010-02064-64 Seafoam Digital Dial Blue Rubber Strap"
		productPrice   int64 = 3999000
		productBrandID int64 = 1552664566511439872

		id int64 = 1552688894279946240

		product = domain.Product{
			Name:    productName,
			Price:   productPrice,
			BrandID: productBrandID,
		}
		productWithID = domain.Product{
			ID:      id,
			Name:    productName,
			Price:   productPrice,
			BrandID: productBrandID,
		}

		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *adminTest)
	}{
		{
			name: "create error",
			test: func(t *GomegaWithT, a *adminTest) {
				createProduct := product
				a.productRepository.On(Create, mock.Anything, &createProduct).
					Return(errDatabase)

				_, err := a.adminUseCase.CreateProduct(context.Background(), product)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "get by ID error",
			test: func(t *GomegaWithT, a *adminTest) {
				createProduct := product
				a.productRepository.On(Create, mock.Anything, &createProduct).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*domain.Product)
						arg.ID = id
					}).
					Return(nil)
				a.productRepository.On(GetByID, mock.Anything, id).
					Return(domain.Product{}, errDatabase)

				_, err := a.adminUseCase.CreateProduct(context.Background(), product)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, a *adminTest) {
				createProduct := product
				a.productRepository.On(Create, mock.Anything, &createProduct).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*domain.Product)
						arg.ID = id
					}).
					Return(nil)
				a.productRepository.On(GetByID, mock.Anything, id).
					Return(productWithID, nil)

				newProduct, err := a.adminUseCase.CreateProduct(context.Background(), product)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(newProduct).Should(Equal(productWithID))
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
