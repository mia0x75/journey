package structure

// User TODO
type User struct {
	ID       int64
	Name     []byte
	Slug     string
	Email    []byte
	Image    []byte
	Cover    []byte
	Bio      []byte
	Website  []byte
	Location []byte
	Twitter  []byte
	Facebook []byte
	Role     int //1 = Administrator, 2 = Editor, 3 = Author, 4 = Owner
}
