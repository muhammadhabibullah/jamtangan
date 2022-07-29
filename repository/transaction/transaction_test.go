package transaction_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/gomega"

	"jamtangan/domain"
	mocks "jamtangan/mock"
	"jamtangan/repository/transaction"
)

type transactionTest struct {
	transactionRepository domain.TransactionRepository
	sqlDB                 *sql.DB
	mock                  sqlmock.Sqlmock
	snowflakeMock         *mocks.Snowflake
}

func test(t *testing.T, fn func(*gomega.WithT, *transactionTest)) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	snowflakeMock := new(mocks.Snowflake)
	transactionRepository := transaction.NewRepository(sqlDB, snowflakeMock)

	b := transactionTest{
		transactionRepository: transactionRepository,
		sqlDB:                 sqlDB,
		mock:                  mock,
		snowflakeMock:         snowflakeMock,
	}

	g := gomega.NewGomegaWithT(t)
	fn(g, &b)

	if err = b.mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
