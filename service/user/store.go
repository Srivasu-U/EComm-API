package user

import (
	"database/sql"
	"fmt"

	"github.com/Srivasu-U/EComm-API/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email) // Returns pointers to rows
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() { // Iterate through rows
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil // Placeholder
}

func (s *Store) CreateUser(user types.User) error {
	return nil // Placeholder
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan( // Scan the row values into User struct
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
