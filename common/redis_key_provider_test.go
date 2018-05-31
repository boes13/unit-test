package common

import (
	"fmt"
	"testing"
)

func TestGetRedisKeyUser(t *testing.T) {
	input := int64(123123)
	expected := fmt.Sprintf(redisKeyUser, input)

	output := GetRedisKeyUser(input)
	if output != expected {
		t.Errorf("expected output %s, got %s", expected, output)
	}
}
