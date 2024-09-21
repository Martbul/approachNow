
// var logger = hclog.Default()
// var pool *pgxpool.Pool

// func InitDB() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		logger.Error("Unable to load .env file", "error", err)
// 	}

// 	dbConnStr := os.Getenv("DB_NEON_CONN_STR")

// 	var errConn error
// 	pool, errConn = pgxpool.Connect(context.Background(), dbConnStr)
// 	if errConn != nil {
// 		logger.Error("Unable to connect to DB", "error", err)
// 	}

// 	logger.Info("Connected to PostgreSQL with provider NeonDB")
// }

// // Query executes an SQL query using the connection pool.
// func Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
// 	return pool.Query(ctx, sql, args...)
// }

// // Close closes the connection pool when the app is done.
// func Close() {
// 	pool.Close()
// }


// func CreateUser(ctx context.Context, username, email, passwordHash string) error {
// 	_, err := Query(ctx, "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", username, email, passwordHash)
// 	if err != nil {
// 		return fmt.Errorf("failed to create user: %w", err)
// 	}
// 	return nil
// }

// func GetUserByEmail(ctx context.Context, email string) (*User, error) {
// 		log.Println("here2.1")

// 	row := pool.QueryRow(ctx, "SELECT username, password_hash FROM users WHERE email = $1", email)
// 			log.Println("here2.2")

// 	var user User
// 	err := row.Scan(&user.Username, &user.PasswordHash)
// 	if err != nil {

// 		if err == pgx.ErrNoRows {
// 			return nil, nil // No user found
// 		}
		
// 		return nil, err
// 	}
// 	log.Println(user)
// 	return &user, nil
// }
package data

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var logger = hclog.Default()
var pool *pgxpool.Pool

// InitDB initializes the connection pool.
func InitDB() {
    err := godotenv.Load(".env")
    if err != nil {
        logger.Error("Unable to load .env file", "error", err)
        return
    }

    dbConnStr := os.Getenv("DB_NEON_CONN_STR")

    if dbConnStr == "" {
        logger.Error("DB_NEON_CONN_STR is not set in .env")
        return
    }

    var errConn error
    pool, errConn = pgxpool.Connect(context.Background(), dbConnStr)
    if errConn != nil {
        logger.Error("Unable to connect to DB", "error", errConn)
        return
    }

    if pool == nil {
        logger.Error("Database pool is nil after initialization")
        return
    }

    logger.Info("Connected to PostgreSQL with provider NeonDB")
}


// Query executes an SQL query using the connection pool.
func Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if pool == nil {
		//Todo: when sending user loc and the getting the id from the jwt token pool is not initialized,
		//Todo: here i am colling InitDB, but htis is a temporere solution
		InitDB()
		// return nil, fmt.Errorf("database pool is NOT initialized")
		return pool.Query(ctx, sql, args...)
	}
	return pool.Query(ctx, sql, args...)
}

// Close closes the connection pool when the app is done.
func Close() {
	if pool != nil {
		pool.Close()
	}
}

// CreateUser creates a new user in the database.
func CreateUser(ctx context.Context, username, email, passwordHash string) error {
	_, err := Query(ctx, "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", username, email, passwordHash)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByEmail retrieves a user by email.
func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	if pool == nil {
		return nil, fmt.Errorf("database pool is not initialized")
	}

	log.Println("here2.1")

	row := pool.QueryRow(ctx, "SELECT username, password_hash FROM users WHERE email = $1", email)
	log.Println("here2.2")

	var user User
	err := row.Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	log.Println(user)
	return &user, nil
}
