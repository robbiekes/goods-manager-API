package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/pashagolub/pgxmock"
	"github.com/robbiekes/goods-manager-api/pkg/postgres"
	"testing"
)

func TestRepository_ReserveItems(t *testing.T) {
	// ctx := context.Background()

	type TestCase struct {
		Name         string
		Expectations func()
		Result       int
	}

	// mock pool
	mockPool, err := pgxmock.NewPool()
	if err != nil {
		t.Error()
	}
	defer mockPool.Close()

	// mock Postgres
	postgresMock := &postgres.Postgres{
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		Pool:    mockPool,
	}

	_ = postgresMock

	testcases := []TestCase{
		{
			Name: "OK",
			Expectations: func() {
				rows := mockPool.NewRows([]string{"id", "balance"}).
					AddRow(1, 500)

				mockPool.ExpectQuery("select itemid from items where storageid =  and reserved = false").
					WithArgs().
					WillReturnRows(rows)
			},
		},
	}

	_ = testcases

}
