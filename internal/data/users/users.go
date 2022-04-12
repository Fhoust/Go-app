package repositories

import (
	"context"
	"database/sql"
	repositories "github.com/zerepl/go-app/internal/data"
	"github.com/zerepl/go-app/internal/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	rows, err := u.db.QueryContext(ctx, "SELECT * FROM users")
	defer rows.Close()

	if err != nil {
		log.Println(err) // TODO: Add log level
		return nil, err
	}

	for rows.Next() {
		var user model.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	return users, err
}

func (u userRepository) CreateNewUser(ctx context.Context, user model.User) (groupId int64, err error) {
	row, err := u.db.ExecContext(ctx, `INSERT INTO users(name) VALUES(?);`, user.Name)

	if err != nil {
		log.Println(err) // TODO: Add log level
		return -1, err
	}

	groupId, err = row.LastInsertId()
	if err != nil {
		log.Println(err) // TODO: Add log level
		return -1, err
	}

	return groupId, nil
}

func (u userRepository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user := model.User{ID: id}

	row, err := u.db.QueryContext(ctx, "SELECT name FROM users WHERE id = ?", id)
	defer row.Close()

	if err != nil {
		log.Println(err) // TODO: Add log level
		return nil, err
	}

	row.Next()
	row.Scan(&user.Name)

	return &user, err
}

func (u userRepository) UpdateUser(ctx context.Context, user model.User) error {
	_, err := u.db.ExecContext(ctx, `UPDATE users SET name = ? WHERE id = ?`, user.Name, user.ID)
	if err != nil {
		log.Println(err) // TODO: Add log level
		return err
	}

	return nil
}

func (u userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := u.db.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, id)

	if err != nil {
		log.Println(err) // TODO: Add log level
		return err
	}

	return nil
}
