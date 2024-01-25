package tools_test

import (
	"testing"

	"payment/pkg/tools"

	"github.com/stretchr/testify/assert"
)

func TestIsExistSlice(t *testing.T) {
	var (
		payload = []struct {
			want bool
			got  string
		}{
			{true, "test1"},
			{false, "test66"},
		}
		slice = []string{
			"test1", "test2", "test3",
		}
	)
	t.Run("Test on get value is exist in string slice", func(t *testing.T) {
		for _, item := range payload {
			assert.Equal(t, item.want, tools.IsExistSlice(item.got, slice))
		}
	})
}

func TestStringToInt(t *testing.T) {
	var payload = []struct {
		want int
		got  string
	}{
		{10, "10"},
		{0, "fail"},
	}
	t.Run("Test on get value is exist in string slice", func(t *testing.T) {
		for _, item := range payload {
			assert.Equal(t, item.want, tools.StringToInt(item.got))
		}
	})
}

func TestReverseSign(t *testing.T) {
	var payload = []struct {
		want float64
		got  float64
	}{
		{-1, 1},
		{1, -1},
	}
	t.Run("Test on get value is exist in string slice", func(t *testing.T) {
		for _, item := range payload {
			assert.Equal(t, item.want, tools.ReverseSign(item.got))
		}
	})
}
