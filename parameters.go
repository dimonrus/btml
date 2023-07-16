package btml

// Parameters for render template
type Parameters map[string]any

// IsTemplate check if template
func (p Parameters) IsTemplate(name string) (t ITemplate, ok bool) {
	if v, k := p[name]; k {
		if t, ok = v.(ITemplate); ok {
			return
		} else if m, okP := v.(Parameters); okP {
			if n, okN := m["_n_"]; okN {
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
