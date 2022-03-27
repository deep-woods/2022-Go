package Users

import database "RESTAPI/pkg/Database"

type service struct {
	db database.DBconnection
}

func (s *service) AddUser(user *User) {
	query := `INSERT INTO users(FirstName, LastName, Eamil) VALUES (?, ?, ?)`
	s.db.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}
