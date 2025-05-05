package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// User struct represents a user in the database.
type User struct {
	Id             int    `json:"id,omitempty"`
	Archive        bool   `json:"archive,omitempty"`
	Public         bool   `json:"public,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"-"` // hashed password, hide in JSON output
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	Gender         int    `json:"gender,omitempty"` // male 1 | female 2 | other 3
	Age            int    `json:"age,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	About          string `json:"about,omitempty"`
	ProfilePicture int    `json:"profilePicture,omitempty"`
}

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided password matches the hashed password.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CreateUser creates a new user record in the database.
func (db *Database) CreateUser(u User) (int, error) {
	// Hash the password before storing it.
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	stmt, err := db.db.Prepare(`
		INSERT INTO user (archive, public, email, password, first_name, last_name, gender, age, nickname, about, profile_picture) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Archive, u.Public, u.Email, hashedPassword, u.FirstName, u.LastName, u.Gender, u.Age, u.Nickname, u.About, u.ProfilePicture)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return 0, fmt.Errorf("email already exists: %w", err)
		}
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	return int(userID), nil
}

// FetchUser retrieves a user record from the database by user ID.
func (db *Database) FetchUser(userID int) (*User, error) {
	row := db.db.QueryRow(`SELECT id, archive, public, email, password, first_name, last_name, gender, age, nickname, about, profile_picture FROM user WHERE id = ?`, userID)

	if user, err := scanUserRecord(row); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	} else {
		return user, nil
	}
}

// FetchUserByEmail retrieves a user record from the database by email.
func (db *Database) FetchUserByEmail(email string) (*User, error) {
	row := db.db.QueryRow(`SELECT id, archive, public, email, password, first_name, last_name, gender, age, nickname, about, profile_picture FROM user WHERE email = ?`, email)

	if user, err := scanUserRecord(row); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	} else {
		return user, nil
	}
}

func scanUserRecord(row *sql.Row) (*User, error) {
	var u User

	err := row.Scan(&u.Id, &u.Archive, &u.Public, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.Gender, &u.Age, &u.Nickname, &u.About, &u.ProfilePicture)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &u, nil
}

// ArchiveUser archives a user by setting the 'archive' flag to true (soft delete)
func (db *Database) ArchiveUser(userID int) error {
	stmt, err := db.db.Prepare(`UPDATE user SET archive = TRUE WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}

// UpdateUser updates an existing user record.
func (db *Database) UpdateUser(user User) error {

	stmt, err := db.db.Prepare(`
		UPDATE user
		SET public = ?,
			first_name = ?,
			last_name = ?,
			age = ?,
			nickname = ?,
			about = ?,
			gender = ?,
			profile_picture = ?
		WHERE id = ?`)

	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %w", err)
	}
	defer stmt.Close()

	// Execute the statement, passing arguments in the correct order
	_, err = stmt.Exec(
		user.Public,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Nickname,
		user.About,
		user.Gender,
		user.ProfilePicture,
		user.Id,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") && strings.Contains(err.Error(), "user.email") {
			return fmt.Errorf("email '%s' already exists: %w", user.Email, err)
		}
		return fmt.Errorf("failed to execute update statement for user id %d: %w", user.Id, err)
	}

	return nil
}

func (u *User) Marshal() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user to JSON: %w", err)
	}

	return string(userJSON), nil
}
