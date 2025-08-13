package methods

// Debug
func (b *Bkash) Debug(enable bool) *Bkash {
	b.debug = enable
	return b
}

func requireNonEmpty(fields ...string) bool {
	for _, f := range fields {
		if f == "" {
			return false
		}
	}
	return true
}
