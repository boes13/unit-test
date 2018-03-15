package redis_with_mocking

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInit(t *testing.T) {
	mockNetwork := "tcp"
	mockAddress := "somehost:3245"

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mock := NewMockRedisI(mockController)

	mock.EXPECT().
		Dial(mockNetwork, mockAddress).
		Times(1).
		Return(nil, fmt.Errorf("test error"))
	rdsWrapper.redis = mock
	network = mockNetwork
	address = mockAddress

	err := Init()
	if err == nil {
		t.Error("expected Init error, got no error")
	}

	mock.EXPECT().
		Dial(mockNetwork, mockAddress).
		Times(1).
		Return(nil, nil)
	rdsWrapper.redis = mock
	network = mockNetwork
	address = mockAddress

	err = Init()
	if err != nil {
		t.Error("expected Init not error, got error", err.Error())
	}
}
