package methods

// Debug
func (b *Bkash) Debug(enable bool) *Bkash {
	b.debug = enable
	return b
}
