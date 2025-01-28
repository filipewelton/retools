package valueobjects

type UserRole struct {
	value string
}

func (u *UserRole) Assign(role string) {
	switch role {
	case "admin":
		u.value = role
	case "user":
		fallthrough
	default:
		u.value = role
	}
}

func (u *UserRole) GetValue() string {
	return u.value
}
