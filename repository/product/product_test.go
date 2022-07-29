package product_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/gomega"

	"jamtangan/domain"
	mocks "jamtangan/mock"
	"jamtangan/repository/product"
)

type productTest struct {
	productRepository domain.ProductRepository
	sqlDB             *sql.DB
	mock              sqlmock.Sqlmock
	snowflakeMock     *mocks.Snowflake
}

func test(t *testing.T, fn func(*gomega.WithT, *productTest)) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	snowflakeMock := new(mocks.Snowflake)
	productRepository := product.NewRepository(sqlDB, snowflakeMock)

	b := productTest{
		productRepository: productRepository,
		sqlDB:             sqlDB,
		mock:              mock,
		snowflakeMock:     snowflakeMock,
	}

	g := gomega.NewGomegaWithT(t)
	fn(g, &b)

	if err = b.mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
