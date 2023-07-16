# btml
It's template rendered. Based on template string and params. Include nested templates

## How to use

- Define template and nested templates if it's needed. Example with nested templates
```go
type Header struct {
    *btml.Layer
	Date string `json:"date"`
}

func NewHeader() *Header {
	return &Header{Layer: &btml.Layer{
		Name: "header",
		Text: `Today is {{ .date }}`,
	}}
}

type Body struct {
    *btml.Layer
}

func NewBody() *Body {
	return &Body{Layer: &btml.Layer{
		Name: "body",
		Text: `Love is gonna save us`,
	}}
}

type Footer struct {
    *btml.Layer
	Copyrights string `json:"copyrights"`
}

func NewFooter() *Footer {
	return &Footer{Layer: &btml.Layer{
		Name: "footer",
		Text: `Copyrights @{{ .copyrights }}`,
	}}
}

type Layout struct {
	*btml.Layer
	Header *Header `json:"header"`
	Body   *Body   `json:"body"`
	Footer *Footer `json:"footer"`
}

func NewLayout(header *Header, body *Body, footer *Footer) *Layout {
	return &Layout{
		Layer: &btml.Layer{
			Name: "layout",
			Text: `{{ .header }}. {{ .body }}. {{ .footer }}.`,
		},
		Header: header,
		Body:   body,
		Footer: footer,
	}
}
```
The Layout is main template in current case

- Add dynamic parameters for template render
```go
now := time.Date(2001, 12, 13, 23, 59, 59, 0, time.Local)
header := NewHeader()
header.Date = now.Format(time.DateOnly)
body := NewBody()
footer := NewFooter()
footer.Copyrights = "Company"
layout := NewLayout(header, body, footer)
```
- Render converted template. Use always Convert method for correct converting layer struct into template.
```go
output, err := btml.Convert(layout).Render()
if err != nil {
    panic(err)
}
```
- For resending template marshal them to json
```go
resend, err := json.Marshal(layout)
if err != nil {
	panic(err)
}
common := NewTemplate()
err = json.Unmarshal(resend, common)
if err != nil {
    panic(err)
}
output, err = common.Render()
```

#### If you find this package useful or want to support the author, you can send tokens to any of these wallets
- Bitcoin: bc1qgx5c3n7q26qv0tngculjz0g78u6mzavy2vg3tf
- Ethereum: 0x62812cb089E0df31347ca32A1610019537bbFe0D
- Dogecoin: DET7fbNzZftp4sGRrBehfVRoi97RiPKajV
