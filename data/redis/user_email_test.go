package redis

import (
	"fmt"
	"log"
	"testing"

	"github.com/boes13/unit-test/common"
	"github.com/golang/mock/gomock"
)

func TestGetUserEmail(t *testing.T) {
	type testCase struct {
		userID      int64
		emailOutput string
		error       bool
		mockExpect  func()
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDataProvider := common.NewMockDataProviderI(mockController)
	common.SetDataProvider(mockDataProvider)

	mockRedisConn := common.NewMockConn(mockController)

	testCases := make(map[string]testCase)

	userID := int64(123)
	email := "someone@somehost.com"

	testCases["redis conn error"] = testCase{
		userID: userID,
		error:  true,
		mockExpect: func() {
			mockDataProvider.EXPECT().
				GetRedisConnection().
				Times(1).
				Return(mockRedisConn)
			mockRedisConn.EXPECT().
				Do("GET", common.GetRedisKeyUser(userID)).
				Times(1).
				Return(nil, fmt.Errorf("test error"))
		},
	}

	testCases["normal"] = testCase{
		userID:      userID,
		error:       false,
		emailOutput: email,
		mockExpect: func() {
			mockDataProvider.EXPECT().
				GetRedisConnection().
				Times(1).
				Return(mockRedisConn)
			mockRedisConn.EXPECT().
				Do("GET", common.GetRedisKeyUser(userID)).
				Times(1).
				Return(email, nil)
		},
	}

	for k, v := range testCases {
		log.Println("running test case:", k)
		v.mockExpect()
		email, err := GetUserEmail(v.userID)
		log.Println("err=", err)
		if v.error && err == nil {
			t.Error("expected error, got no error")
			continue
		}

		if !v.error && v.emailOutput != email {
			t.Errorf("expected email %s, got %s", v.emailOutput, email)
		}
	}
}
