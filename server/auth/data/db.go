package data
import (
	"context"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var logger = hclog.Default()
var pool *pgxpool.Pool

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Unable to load .env file", "error", err)
	}

	dbConnStr := os.Getenv("DB_NEON_CONN_STR")

	var errConn error
	pool, errConn = pgxpool.Connect(context.Background(), dbConnStr)
	if errConn != nil {
		logger.Error("Unable to connect to DB", "error", err)
	}

	logger.Info("Connected to PostgreSQL with provider NeonDB")
}

// Query executes an SQL query using the connection pool.
func Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return pool.Query(ctx, sql, args...)
}

// Close closes the connection pool when the app is done.
func Close() {
	pool.Close()
}
