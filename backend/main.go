package main

import (
	"backend/api"
	"backend/db"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

// Configurable constants, can be loaded from environment variables
var (
	port      = ":8080"
	staticDir = "static"
)

func init() {
	api.GenOrLoadKey()

	// Get environment variables, using default values if not set
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = ":" + envPort
	}

	envStaticDir := os.Getenv("STATIC_DIR")
	if envStaticDir != "" {
		staticDir = envStaticDir
	}

	// this will be used later on in docker build, in dev it's not used
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		if err := os.Mkdir(staticDir, 0755); err != nil {
			log.Fatalf("Failed to create static directory: %v", err)
		}
	}

	genDevToken()

	if err := db.Connection.Open("db.sqlite3"); err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
}

func main() {

	testDB()

	// File server
	fs := http.FileServer(http.Dir(staticDir))

	// Handlers
	http.Handle("/", http.StripPrefix("/", fs)) // Serves files from the static directory
	http.HandleFunc("/api", api.Router)

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func genDevToken() {
	c := api.Claims{
		Username: "admin",
		UserId:   1,
	}

	sig, _ := c.GetBearer()
	log.Printf("Authorization: %s\n", *sig)
}

func testDB() {
	newUser := db.User{
		Username: "testuser",
		Archive:  false,
		Email:    "test@example.com",
		Password: "password123",
		Metadata: db.UserMetadata{
			Age:    12,
			Gender: 0,
		},
	}

	connection := db.Connection

	userID, err := connection.CreateUser(newUser)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
	} else {
		log.Printf("User created with ID: %d", userID)
	}

	// Fetch the created user
	if userID > 0 {
		fetchedUser, err := connection.FetchUser(userID)
		if err != nil {
			log.Printf("Failed to fetch user: %v", err)
		} else {
			fmt.Printf("Fetched User: %+v\n", fetchedUser)
		}

		// Example: Update the user
		if fetchedUser != nil {
			fetchedUser.Metadata.Gender = 1
			err = connection.UpdateUser(*fetchedUser)
			if err != nil {
				log.Printf("Failed to update user: %v", err)
			} else {
				fmt.Println("User updated successfully.")
			}

			// Fetch the updated user
			updatedUser, err := connection.FetchUser(userID)
			if err != nil {
				log.Printf("Failed to fetch updated user: %v", err)
			} else {
				fmt.Printf("Updated User: %+v\n", updatedUser)
			}
		}

		// Archive the user
		if fetchedUser != nil {
			err = connection.ArchiveUser(userID)
			if err != nil {
				log.Printf("Failed to archive user: %v", err)
			} else {
				fmt.Println("User archived successfully.")
			}

			// Fetch the updated user (check archive status)
			archivedUser, err := connection.FetchUser(userID)
			if err != nil {
				log.Printf("Failed to fetch archived user: %v", err)
			} else {
				fmt.Printf("Archived User: %+v\n", archivedUser)
			}
		}
	}

	// Example: Fetch user by username
	userByName, err := connection.FetchUserByUsername("testuser")
	if err != nil {
		log.Printf("Failed to fetch user by username: %v", err)
	} else {
		fmt.Printf("User by username: %+v\n", userByName)
	}
}
