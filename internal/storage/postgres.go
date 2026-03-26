package storage

import (
	"database/sql"
	"fmt"
	"github.com/codeVend/pigeon/internal/auth"
	"context"
)

type PostgresStore struct{
	DB *sql.DB

}
func OpenPostgres(databaseURL string) (*sql.DB, error){
	db,err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	} 
	if err := db.Ping(); err != nil{
		return nil, fmt.Errorf("Ping postgret %w", err)
	}
	return db, nil
}

func (s *PostgresStore) CreateUser(ctx context.Context, user *auth.User) error {
    const query = `
        INSERT INTO users (username, email, password, about, photo, number)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `

    return s.DB.QueryRowContext(
        ctx,
        query,
        user.Username,
        user.Email,
        user.Password,
        user.About,
        user.Photo,
        user.Number,
    ).Scan(&user.ID)
}