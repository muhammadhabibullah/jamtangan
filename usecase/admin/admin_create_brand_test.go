package admin_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"jamtangan/domain"
)

func TestAdminUseCase_CreateBrand(t *testing.T) {
	t.Parallel()

	const (
		Create  = "Create"
		GetByID = "GetByID"
	)

	var (
		brandName            = "casio"
		brandNameUpper       = "CASIO"
		id             int64 = 1552655170888798208

		brand = domain.Brand{
			Name: brandNameUpper,
		}
		brandWithID = domain.Brand{
			ID:   id,
			Name: brandNameUpper,
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
				createBrand := brand
				a.brandRepository.On(Create, mock.Anything, &createBrand).
					Return(errDatabase)

				_, err := a.adminUseCase.CreateBrand(context.Background(), brandName)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "get by ID error",
			test: func(t *GomegaWithT, a *adminTest) {
				createBrand := brand
				a.brandRepository.On(Create, mock.Anything, &createBrand).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*domain.Brand)
						arg.ID = id
					}).
					Return(nil)
				a.brandRepository.On(GetByID, mock.Anything, id).
					Return(domain.Brand{}, errDatabase)

				_, err := a.adminUseCase.CreateBrand(context.Background(), brandName)
				t.Expect(err).Should(Equal(errDatabase))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, a *adminTest) {
				createBrand := brand
				a.brandRepository.On(Create, mock.Anything, &createBrand).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*domain.Brand)
						arg.ID = id
					}).
					Return(nil)
				a.brandRepository.On(GetByID, mock.Anything, id).
					Return(brandWithID, nil)

				newBrand, err := a.adminUseCase.CreateBrand(context.Background(), brandName)
				t.Expect(err).ShouldNot(HaveOccurred())
				t.Expect(newBrand).Should(Equal(brandWithID))
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
