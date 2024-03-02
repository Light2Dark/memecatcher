package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDBConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to postgres db:", version)

	return db, nil
}

// I can't create a table through the UI/CLI, so I have to do it through code
func CreateTables(db *sql.DB) error {
	queryOne := `
	DROP TABLE IF EXISTS memes;
	DROP TABLE IF EXISTS users;
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	queryTwo := `
	CREATE TABLE IF NOT EXISTS memes (
		id SERIAL PRIMARY KEY,
		user_id TEXT NOT NULL,
		image_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err := db.Exec(queryOne)
	if err != nil {
		return err
	}
	fmt.Println("Created users table")

	_, err = db.Exec(queryTwo)
	if err != nil {
		return err
	}
	fmt.Println("Created memes table")

	return nil
}

func CreateUser(db *sql.DB, userID string) error {
	query := `
	INSERT INTO users (id)
	VALUES ($1);
	`

	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

func DoesUserExist(db *sql.DB, userID string) (bool, error) {
	query := `
	SELECT id
	FROM users
	WHERE id = $1;
	`

	row := db.QueryRow(query, userID)
	if row.Err() == sql.ErrNoRows {
		return false, fmt.Errorf("user not found")
	}

	return true, nil
}

func InsertMeme(db *sql.DB, userID string, imageUrl string) error {
	query := `
	INSERT INTO memes (user_id, image_url)
	VALUES ($1, $2);
	`

	_, err := db.Exec(query, userID, imageUrl)
	if err != nil {
		return err
	}

	return nil
}

// Pass in <= 0 for no limit
func GetAllMemes(db *sql.DB, userID string, limit int) ([]Meme, error) {
	query := `
	SELECT user_id, image_url
	FROM memes
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	query += ";"

	var memes []Meme
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var meme Meme
		err = rows.Scan(&meme.UserID, &meme.ImageURL)
		if err != nil {
			return nil, err
		}
		memes = append(memes, meme)
	}

	return memes, nil
}

func SelectAllUsers(db *sql.DB) ([]User, error) {
	query := `
	SELECT id
	FROM users;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
