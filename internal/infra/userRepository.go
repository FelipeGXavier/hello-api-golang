package infra

import (
	"context"
	"errors"
	"hello-api-go/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Add(user entity.User) (entity.User, error);
	Update(user entity.User, id int) (entity.User, error);
	Delete(id int) error;
	FindById(id int) (entity.User, error);
}

type PgUserRepository struct {
	db *pgxpool.Pool;
}

func (u PgUserRepository) Add(user entity.User) (entity.User, error) {
	insertStatement := `INSERT INTO users (name, surname, email) VALUES ($1, $2, $3) RETURNING id`
	var userID int
	err := u.db.QueryRow(context.Background(), insertStatement, user.Name, user.Surname, user.Email).Scan(&userID)
	if err != nil {
		return entity.User{}, err;
	}
	user.Id = userID;
	return user, nil;
}

func (u PgUserRepository) Update(user entity.User, id int) (entity.User, error) {
	return entity.User{}, errors.New("Not implemented");
}

func (u PgUserRepository) Delete(id int) (error) {
	return errors.New("Not implemented");
}

func (u PgUserRepository) FindById(id int) (entity.User, error) {
	return entity.User{}, errors.New("Not implemented");
}

func MakeUserRepository(db *pgxpool.Pool) UserRepository {
	return PgUserRepository{db: db};
}