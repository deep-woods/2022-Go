package Users

// there is no risk of circular dependency here.

type service struct { // this service uses this Repository (repo).
	repo Repository // oracle, sql, doesn't matter.
}

type Repository interface { // this is a repository for service struct, and contains db transaction functionality.
	ExecuteQuery(query string, value ...interface{})
}

type Service interface { // the reason for declaring this interface is because
	AddUser(user *User) // we want to declare the below function.
}

func NewService(repo Repository) Service { // wrapper function?
	return &service{repo}
}

func (s *service) AddUser(user *User) {
	query := `INSERT INTO users(FirstName, LastName, Eamil) VALUES (?, ?, ?)`
	// insert into 'users' table.
	s.repo.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}

/*
2022/03/27 18:41:44 http: panic serving 127.0.0.1:62818: runtime error: invalid memory address or nil pointer dereference
goroutine 18 [running]:
*/
