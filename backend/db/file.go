package db

import (
	"fmt"
)

func (db *Database) UploadImage(filename string, mimetype string, imageData []byte) (int, error) {

	if len(imageData) == 0 {
		return 0, fmt.Errorf("image data cannot be empty")
	}
	if filename == "" {
		return 0, fmt.Errorf("filename cannot be empty")
	}
	if mimetype == "" {
		return 0, fmt.Errorf("mimetype cannot be empty")
	}

	result, err := db.db.Exec(`INSERT INTO file (data, filename, mimetype) VALUES (?, ?, ?)`,
		imageData, filename, mimetype)

	if err != nil {
		return 0, fmt.Errorf("failed to execute insert statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	return int(id), nil
}
