package btml

import (
	"testing"
)

type Item struct {
	Name string
	Text string
}

type Enricher []Item

func (e Enricher) GetText(name string) string {
	for _, item := range e {
		if item.Name == name {
			return item.Text
		}
	}
	return ""
}

var testEnricher = Enricher{
	{Name: "header", Text: "You need this"},
	{Name: "footer", Text: "Leave it to us"},
}

func TestParameters_EnrichText(t *testing.T) {
	t.Run("enrich", func(t *testing.T) {
		header := NewHeader()
		header.Text = ""
		body := NewBody()
		footer := NewFooter()
		footer.Text = ""
		footer.Copyrights = "Company"
		l := NewLayout(header, body, footer)

		tml := Convert(l)
		params := tml.GetParams().EnrichText(testEnricher, false)
		if params["header"].(Parameters)[TemplateTextKey] != "You need this" {
			t.Fatal("wrong enricher fore header")
		}
		if params["footer"].(Parameters)[TemplateTextKey] != "Leave it to us" {
			t.Fatal("wrong enricher for footer")
		}
		if params["body"].(Parameters)[TemplateTextKey] != "Love is gonna save us" {
			t.Fatal("wrong enricher for body")
		}
		tml.SetParams(params)
		data, err := tml.Render()
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "You need this. Love is gonna save us. Leave it to us." {
			t.Fatal("wrong render")
		}
	})
}
