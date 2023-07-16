package btml

import (
	"encoding/json"
	"testing"
	"time"
)

type Header struct {
	*Layer
	Date string `json:"date"`
}

func NewHeader() *Header {
	return &Header{Layer: &Layer{
		Name: "header",
		Text: `Today is {{ .date }}`,
	}}
}

type Body struct {
	*Layer
}

func NewBody() *Body {
	return &Body{Layer: &Layer{
		Name: "body",
		Text: `Love is gonna save us`,
	}}
}

type Footer struct {
	*Layer
	Copyrights string `json:"copyrights"`
}

func NewFooter() *Footer {
	return &Footer{Layer: &Layer{
		Name: "footer",
		Text: `Copyrights @{{ .copyrights }}`,
	}}
}

type Layout struct {
	*Layer
	Header *Header `json:"header"`
	Body   *Body   `json:"body"`
	Footer *Footer `json:"footer"`
}

func NewLayout(header *Header, body *Body, footer *Footer) *Layout {
	return &Layout{
		Layer: &Layer{
			Name: "layout",
			Text: `{{ .header }}. {{ .body }}. {{ .footer }}.`,
		},
		Header: header,
		Body:   body,
		Footer: footer,
	}
}

func TestCommonCase(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		now := time.Date(2001, 12, 13, 23, 59, 59, 0, time.Local)
		header := NewHeader()
		header.Date = now.Format(time.DateOnly)
		body := NewBody()
		footer := NewFooter()
		footer.Copyrights = "Company"
		l := NewLayout(header, body, footer)
		output, err := Convert(l).Render()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(output))
		if string(output) != "Today is 2001-12-13. Love is gonna save us. Copyrights @Company." {
			t.Fatal("wrong render")
		}
		data, err := json.Marshal(l)
		if err != nil {
			t.Fatal(err)
		}
		layout := NewLayout(NewHeader(), NewBody(), NewFooter())
		err = json.Unmarshal(data, layout)
		if err != nil {
			t.Fatal(err)
		}
		if layout.Footer.Copyrights != "Company" {
			t.Fatal("wrong unmarshal")
		}
		resend, err := json.Marshal(layout)
		if err != nil {
			t.Fatal(err)
		}
		l1 := NewTemplate()
		err = json.Unmarshal(resend, l1)
		if err != nil {
			t.Fatal(err)
		}
		output, err = l1.Render()
		if err != nil {
			t.Fatal(err)
		}
		if string(output) != "Today is 2001-12-13. Love is gonna save us. Copyrights @Company." {
			t.Fatal("wrong render")
		}
	})
}

func TestMarshal(t *testing.T) {
	t.Run("forward", func(t *testing.T) {
		tmpl := NewTemplate("new", "some text")
		layer := NewLayer("forward", "forward template text")
		tmpl.SetParam("some_nested", layer)
		data, err := json.Marshal(tmpl)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(data))
		if string(data) != `{"_n_":"new","_t_":"some text","some_nested":{"_n_":"forward","_t_":"forward template text"}}` {
			t.Fatal("wrong marshal")
		}
	})
	t.Run("backward", func(t *testing.T) {
		data := []byte(`{"_n_":"new","_t_":"some text","some_nested":{"_n_":"forward","_t_":"forward template text"}}`)
		tmp := NewTemplate()
		err := json.Unmarshal(data, &tmp)
		if err != nil {
			t.Fatal(err)
		}
		if tmp.GetName() != "new" || len(tmp.GetParams()) != 3 {
			t.Fatal("wrong unmarshal")
		}
	})
}

func TestRenderTemplate(t *testing.T) {
	t.Run("same_unmarshal", func(t *testing.T) {
		now := time.Now()
		header := NewHeader()
		header.Date = now.Format(time.DateOnly)
		body := NewBody()
		footer := NewFooter()
		footer.Copyrights = "Dmitry"
		l := NewLayout(header, body, footer)
		data, err := json.Marshal(l)
		if err != nil {
			t.Fatal(err)
		}
		render := NewLayout(NewHeader(), NewBody(), NewFooter())
		err = json.Unmarshal(data, render)
		if err != nil {
			t.Fatal(err)
		}
		if render.Footer.Copyrights != "Dmitry" {
			t.Fatal("wrong unmarshal")
		}
	})
}
