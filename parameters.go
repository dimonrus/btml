package btml

// Parameters for render template
type Parameters map[string]any

// TextEnricher extract text for enrich params
type TextEnricher interface {
	GetText(name string) string
}

// CheckTemplateParams check if params template
func (p Parameters) CheckTemplateParams(name string) (params Parameters, ok bool) {
	if v, k := p[name]; k {
		if _, ok = v.(ITemplate); ok {
			return
		} else if ps, okPS := v.(Parameters); okPS {
			if _, okN := ps[TemplateNameKey]; okN {
				params = ps
				ok = true
			}
		} else if m, okP := v.(map[string]any); okP {
			if _, okN := m[TemplateNameKey]; okN {
				params = m
				ok = true
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
