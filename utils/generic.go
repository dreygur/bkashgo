package utils

func RequireNonEmpty(fields ...string) bool {
	for _, f := range fields {
		if f == "" {
			return false
		}
	}
	return true
}
