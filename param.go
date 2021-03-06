package baserouter

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) ByName(name string) string {
	for _, p := range ps {
		if p.Key == name {
			return p.Value
		}
	}
	return ""
}

func (p *Params) appendKey(key string) {

	*p = append(*p, Param{Key: key})
}

func (p *Params) setVal(val string) {
	(*p)[len(*p)-1].Value = val
}
