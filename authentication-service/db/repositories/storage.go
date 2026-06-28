package db

type Storage struct{ //dependency injection
	UserRepository UserRepository
}

func NewStorage() *Storage{
	return &Storage{
		UserRepository:&UsserRepositoryImpl{},
	}
}