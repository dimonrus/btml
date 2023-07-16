package btml

import (
	"encoding/json"
	"errors"
)

// VLayer layer
type VLayer interface {
	Valid() error
}

// Layer include into each template struct
type Layer struct {
	// Template code
	Name string `json:"_n_"`
	// Text of template
	Text string `json:"_t_"`
}

// Valid is layer valid
func (l *Layer) Valid() error {
	if l.Name == "" {
		return errors.New("layer name is required")
	}
	return nil
}

// NewLayer layer constructor
func NewLayer(arg ...string) *Layer {
	l := &Layer{}
	if len(arg) == 1 {
		l.Name = arg[0]
	} else if len(arg) == 2 {
		l.Name = arg[0]
		l.Text = arg[1]
	}
	return l
}

// Convert layer to ITemplate
func Convert(layer VLayer) ITemplate {
	if err := layer.Valid(); err != nil {
		println(err)
		return nil
	}
	data, err := json.Marshal(layer)
	if err != nil {
		// print but not return
		println(err)
		return nil
	}
	tmpl := NewTemplate()
	err = json.Unmarshal(data, tmpl)
	if err != nil {
		// print but not return
		println(err)
		return nil
	}
	return tmpl
}
