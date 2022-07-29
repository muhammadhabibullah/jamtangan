package brand_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/gomega"

	"jamtangan/domain"
	mocks "jamtangan/mock"
	"jamtangan/repository/brand"
)

type brandTest struct {
	brandRepository domain.BrandRepository
	sqlDB           *sql.DB
	mock            sqlmock.Sqlmock
	snowflakeMock   *mocks.Snowflake
}

func test(t *testing.T, fn func(*gomega.WithT, *brandTest)) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	snowflakeMock := new(mocks.Snowflake)
	brandRepository := brand.NewRepository(sqlDB, snowflakeMock)

	b := brandTest{
		brandRepository: brandRepository,
		sqlDB:           sqlDB,
		mock:            mock,
		snowflakeMock:   snowflakeMock,
	}

	g := gomega.NewGomegaWithT(t)
	fn(g, &b)

	if err = b.mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
