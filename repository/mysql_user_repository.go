package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"userapi/models"
)

type mysqlUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository creates a new MySQL user repository
func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (name, age, phone_number, email) VALUES (?, ?, ?, ?)`
	log.Printf("Creating user with email: %s", user.Email)

	result, err := r.db.ExecContext(ctx, query, user.Name, user.Age, user.PhoneNumber, user.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	user.ID = id
	log.Printf("Successfully created user with ID: %d", id)
	return nil
}

func (r *mysqlUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT id, name, age, phone_number, email FROM users WHERE id = ?`
	log.Printf("Fetching user with ID: %d", id)

	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Age, &user.PhoneNumber, &user.Email)
	if err == sql.ErrNoRows {
		log.Printf("User not found with ID: %d", id)
		return nil, nil
	}
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", id, err)
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Printf("Successfully fetched user with ID: %d", id)
	return user, nil
}

func (r *mysqlUserRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET name = ?, age = ?, phone_number = ?, email = ? WHERE id = ?`
	log.Printf("Updating user with ID: %d", user.ID)

	result, err := r.db.ExecContext(ctx, query, user.Name, user.Age, user.PhoneNumber, user.Email, user.ID)
	if err != nil {
		log.Printf("Error updating user with ID %d: %v", user.ID, err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected for user update ID %d: %v", user.ID, err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		log.Printf("No user found to update with ID: %d", user.ID)
		return fmt.Errorf("user not found with ID: %d", user.ID)
	}

	log.Printf("Successfully updated user with ID: %d", user.ID)
	return nil
}

func (r *mysqlUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	log.Printf("Deleting user with ID: %d", id)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting user with ID %d: %v", id, err)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected for user deletion ID %d: %v", id, err)
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		log.Printf("No user found to delete with ID: %d", id)
		return fmt.Errorf("user not found with ID: %d", id)
	}

	log.Printf("Successfully deleted user with ID: %d", id)
	return nil
}

func (r *mysqlUserRepository) List(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, name, age, phone_number, email FROM users`
	log.Println("Fetching all users")

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.PhoneNumber, &user.Email); err != nil {
			log.Printf("Error scanning user row: %v", err)
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating user rows: %v", err)
		return nil, fmt.Errorf("failed to iterate users: %w", err)
	}

	log.Printf("Successfully fetched %d users", len(users))
	return users, nil
} 