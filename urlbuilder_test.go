package urlbuilder

import (
	"testing"
)

func TestUrlBuilder(t *testing.T) {
	builder := Init().WithPort(9091).WithPath("/anh/:id").WithPathParam("id", "123")
	builder.WithQueryParam("anh", "a").WithQueryParam("321", "asda").WithQueryParams(map[string]string{
		"432": "asd",
		"ga":  "123",
	})
	url := builder.Build()
	t.Log(url)
}
