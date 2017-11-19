package api

type Record interface {
	CI() string
	AssignedTo() string
	Number() string
}
