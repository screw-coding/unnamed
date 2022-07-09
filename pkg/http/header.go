package http

type Header map[string][]string

func (h Header) Set(key string, value string) {
	h[key] = append(h[key], value)
}

func (h Header) Del(key string) {
	delete(h, key)
}

func (h Header) get(key string) string {

	v := h[key]
	if len(v) > 0 {
		return v[0]
	}
	return ""

}

func (h Header) Clone() Header {
	if h == nil {
		return nil
	}
	newLength := 0
	for _, s2 := range h {
		newLength += len(s2)
	}

	newValues := make([]string, newLength)
	newHeader := make(Header, len(h))
	for k, s2 := range h {
		n := copy(newValues, s2)
		newHeader[k] = newValues[:n:n]
		newValues = newValues[:n]
	}
	return newHeader
}

func (h Header) Add(key string, value string) {
	h[key] = append(h[key], value)
}
