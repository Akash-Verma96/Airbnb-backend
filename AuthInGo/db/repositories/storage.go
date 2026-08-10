package db

// containes object of all the repositories
// facilitates dependecncy Injection for repositories
type Storage struct {
	UserRepository UserRepository 
}

// constructur of each repositories
func NewStorage() *Storage{
	return &Storage{
		UserRepository: &UserRepositoryImpl{},  // creates a new instance of a struct It is called a struct literal
	}
}