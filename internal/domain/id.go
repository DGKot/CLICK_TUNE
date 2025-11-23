package domain

type ID string

func (id ID) String() string {
	return string(id)
}

func (id ID) IsZero() bool {
	return id == ""
}
