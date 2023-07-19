package btml

// Parameters for render template
type Parameters map[string]any

// TextEnricher extract text for enrich params
type TextEnricher interface {
	GetText(name string) string
}

// IsTemplate check if template
func (p Parameters) IsTemplate(name string) (t ITemplate, ok bool) {
	if v, k := p[name]; k {
		if t, ok = v.(ITemplate); ok {
			return
		} else if m, okP := v.(Parameters); okP {
			if n, okN := m[TemplateNameKey]; okN {
				if s, okS := n.(string); okS {
					if Constructor != nil {
						t = Constructor(s)
					}
				}
			}
		}
	}
	return
}

// EnrichText set text for each template in params
func (p Parameters) EnrichText(te TextEnricher, override bool) Parameters {
	for k, v := range p {
		if k == TemplateNameKey {
			text := te.GetText(p[TemplateNameKey].(string))
			if override || p[TemplateTextKey] == "" {
				p[TemplateTextKey] = text
			}
			continue
		}
		if m, okP := v.(map[string]interface{}); okP {
			p[k] = Parameters(m).EnrichText(te, override)
		}
	}
	return p
}
