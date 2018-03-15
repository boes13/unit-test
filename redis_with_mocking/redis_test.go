package redis_with_mocking

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInsertSaddRedis(t *testing.T) {
	command := "SADD"
	key := "SomeKeyAnything"
	value := "SomeValueAnything"
	retval := int64(1)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mock := NewMockRedisConnI(mockController)
	mock.EXPECT().
		Do(command, key, value).
		Times(1).
		Return("", fmt.Errorf("test error"))
	cw.conn = mock
	_, err := InsertSaddRedis(key, value)
	if err == nil {
		t.Error("expected error in InsertSaddRedis, got no error")
	}

	mock.EXPECT().
		Do(command, key, value).
		Times(1).
		Return(retval, nil)
	cw.conn = mock
	count, err := InsertSaddRedis(key, value)
	if err != nil {
		t.Error("expected no error in InsertSaddRedis, got error", err.Error())
	}
	if count != retval {
		t.Errorf("expected InsertSaddRedis output %d, got output %d", retval, count)
	}
}
