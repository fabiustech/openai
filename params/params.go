// Package params provides a helper function to simplify setting optional
// parameters in struct literals.
package params

// Optional returns a pointer to |v|.
func Optional[T any](v T) *T {
	return &v
}
