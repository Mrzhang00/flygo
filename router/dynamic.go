package router

// Dynamic struct
type Dynamic struct {
	// Pos for params
	Pos map[int]string
	// Simple contains
	*Simple
}
