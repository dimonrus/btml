package btml

import (
	"bytes"
	"encoding/json"
	"text/template"
)

const (
	// TemplateNameKey name of template
	TemplateNameKey = "_n_"
	// TemplateTextKey text of template
	TemplateTextKey = "_t_"
)

// ITemplate bytes template interface based on json representation
type ITemplate interface {
	// GetName return template name
	GetName() string
	// GetText return template text
	GetText() string
	// GetParam return template param
	GetParam(name string) any
	// SetParam set template param
	SetParam(name string, param any) ITemplate
	// GetParams return template params
	GetParams() Parameters
	// SetParams set params
	SetParams(params Parameters) ITemplate
	// Render template with params
	Render() ([]byte, error)
}

// Template Base template struct
type Template struct {
	// custom parameters
	params Parameters
}

// GetName return template name
func (t *Template) GetName() string {
	name, ok := t.params[TemplateNameKey]
	if ok {
		return name.(string)
	}
	return ""
}

// GetText return template text
func (t *Template) GetText() string {
	text, ok := t.params[TemplateTextKey]
	if ok {
		return text.(string)
	}
	return ""
}

// GetParams return template params
func (t *Template) GetParams() Parameters {
	return t.params
}

// SetParams set params
func (t *Template) SetParams(params Parameters) ITemplate {
	t.params = params
	return t
}

// GetParam return template param
func (t *Template) GetParam(name string) any {
	return t.params[name]
}

// SetParam set template param
func (t *Template) SetParam(name string, param any) ITemplate {
	if t.params == nil {
		t.params = make(Parameters)
	}
	t.params[name] = param
	return t
}

// Render render template
func (t *Template) Render() ([]byte, error) {
	for k, v := range t.params {
		if m, ok := v.(map[string]any); ok {
			name, okName := m[TemplateNameKey]
			_, okText := m[TemplateTextKey]
			if okName && okText {
				tmpl := NewTemplate().SetParams(m)
				data, err := tmpl.Render()
				if err != nil {
					return nil, err
				}
				if k != name {
					delete(t.params, k)
				}
				t.params[name.(string)] = string(data)
			}
		}
	}
	tmpl := template.New(t.GetName())
	tmpl, err := tmpl.Parse(t.GetText())
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	err = tmpl.Execute(buf, t.GetParams())
	return buf.Bytes(), err
}

// UnmarshalJSON unmarshaller
func (t *Template) UnmarshalJSON(data []byte) error {
	t.params = make(map[string]any)
	return json.Unmarshal(data, &t.params)
}

// MarshalJSON marshaller
func (t *Template) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.GetParams())
}

// NewTemplate constructor
func NewTemplate(arg ...string) *Template {
	t := &Template{}
	if len(arg) == 1 {
		t.SetParam(TemplateNameKey, arg[0])
	} else if len(arg) == 2 {
		t.SetParam(TemplateNameKey, arg[0])
		t.SetParam(TemplateTextKey, arg[1])
	}
	return t
}
