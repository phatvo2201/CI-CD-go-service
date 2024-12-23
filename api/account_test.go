package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mockdb "github.com/phatvo2201/simplebank/db/mock"
	db "github.com/phatvo2201/simplebank/db/sqlc"
	"github.com/phatvo2201/simplebank/util"
)

func TestGetAccount(t *testing.T) {
	account := randomAccount()

	//Apply DDD
	testCases := []struct {
		name           string
		accountID      int64
		BuildStubsMock func(store *mockdb.MockStore)
		Check          func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{name: "test_ok",
			accountID: account.ID,
			BuildStubsMock: func(mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).Return(account, nil)

			},
			Check: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchAccount(t, recoder.Body, account)

			},
		},
		{name: "not_found",
			accountID: account.ID,
			BuildStubsMock: func(mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).Return(db.Account{}, sql.ErrNoRows)

			},
			Check: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			contrl := gomock.NewController(t)
			defer contrl.Finish()

			mockStore := mockdb.NewMockStore(contrl)
			tc.BuildStubsMock(mockStore)

			serverTest := NewServer(mockStore)
			recoder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			serverTest.router.ServeHTTP(recoder, request)
			tc.Check(t, recoder)

		})
	}

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
