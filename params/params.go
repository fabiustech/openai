package params

// Optional returns a pointer to |v|.
func Optional[T any](v T) *T {
	return &v
}
