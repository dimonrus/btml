package btml

import (
	"testing"
)

type testLayer struct {
	*Layer
	Nested testNestedData `json:"nested"`
}

type testNestedData struct {
	*Layer
	Data string `json:"data"`
}

func TestLayer_Convert(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		d := &testLayer{
			Layer: &Layer{
				Name: "some",
				Text: "some {{ .nested }}",
			},
			Nested: testNestedData{
				Layer: &Layer{
					Name: "nested",
					Text: "whatever {{ .data }}",
				},
				Data: "bitcoin",
			},
		}
		it := Convert(d)
		data, err := it.Render()
		if err != nil {
			t.Fatal(err)
		}
		if "some whatever bitcoin" != string(data) {
			t.Fatal("wrong render")
		}
	})
}
