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
	User     int          `json:"user,omitempty"` // User ID
	Username string       `json:"username,omitempty"`
	Archive  bool         `json:"archive,omitempty"`
	Email    string       `json:"email,omitempty"`
	Password string       `json:"-"` // hashed password, hide in JSON output
	Metadata UserMetadata `json:"metadata,omitempty"`
}

type UserMetadata struct {
	Gender int    `json:"gender,omitempty"` // male 0 | female 1
	Age    int    `json:"age,omitempty"`
	Role   string `json:"role,omitempty"`
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
func (db *Database) CreateUser(user User) (int, error) {
	// Hash the password before storing it.
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	// Marshal the metadata to JSON.
	metadataJSON, err := json.Marshal(user.Metadata)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal metadata to JSON: %w", err)
	}

	stmt, err := db.db.Prepare(`
		INSERT INTO user (username, archive, email, password, metadata) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Archive, user.Email, hashedPassword, string(metadataJSON))
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return 0, fmt.Errorf("username or email already exists: %w", err)
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
	row := db.db.QueryRow(`SELECT user, username, archive, email, password, metadata FROM user WHERE user = ?`, userID)

	if user, err := scanUserRecord(row); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	} else {
		return user, nil
	}
}

// FetchUserByUsername retrieves a user record from the database by username.
func (db *Database) FetchUserByUsername(username string) (*User, error) {
	row := db.db.QueryRow(`SELECT user, username, archive, email, password, metadata FROM user WHERE username = ?`, username)

	if user, err := scanUserRecord(row); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	} else {
		return user, nil
	}
}

// FetchUserByEmail retrieves a user record from the database by email.
func (db *Database) FetchUserByEmail(email string) (*User, error) {
	row := db.db.QueryRow(`SELECT user, username, archive, email, password, metadata FROM user WHERE email = ?`, email)

	if user, err := scanUserRecord(row); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	} else {
		return user, nil
	}
}

func scanUserRecord(row *sql.Row) (*User, error) {
	var user User
	var metadata []byte

	err := row.Scan(&user.User, &user.Username, &user.Archive, &user.Email, &user.Password, &metadata)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if err = json.Unmarshal(metadata, &user.Metadata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}
	return &user, nil
}

// ArchiveUser archives a user by setting the 'archive' flag to true (soft delete)
func (db *Database) ArchiveUser(userID int) error {
	stmt, err := db.db.Prepare(`UPDATE user SET archive = TRUE WHERE user = ?`)
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
	// Hash the password before storing it.
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	metadataJSON, err := json.Marshal(user.Metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata to JSON: %w", err)
	}

	stmt, err := db.db.Prepare(`
		UPDATE user
		SET username = ?, archive = ?, email = ?, password = ?, metadata = ?
		WHERE user = ?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Archive, user.Email, hashedPassword, string(metadataJSON), user.User)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("username or email already exists: %w", err)
		}
		return fmt.Errorf("failed to execute statement: %w", err)
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
