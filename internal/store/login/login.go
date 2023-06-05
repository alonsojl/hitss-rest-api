package login

import (
	"context"
	"database/sql"
	"errors"
	"hitss/internal/db/mysql"
	"hitss/internal/model"
	"hitss/pkg/helper/logger"
	"hitss/pkg/helper/password"
	"hitss/pkg/helper/token"
)

type store struct {
	storage *sql.DB
}

func New() *store {
	return &store{
		storage: mysql.Open(),
	}
}

func (s *store) SignIn(l model.Login) (string, error) {
	defer s.storage.Close()

	ctx := context.Background()
	stmt, err := s.storage.PrepareContext(ctx, "CALL spUserLogin(?)")
	if err != nil {
		logger.Write(err)
		return "", err
	}
	defer stmt.Close()

	u := model.User{}
	err = stmt.QueryRowContext(ctx, l.Email).Scan(&u.Id, &u.Name, &u.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}

	if err != nil {
		logger.Write(err)
		return "", err
	}

	if !password.Compare(u.Password, l.Password) {
		return "", nil
	}

	u.Password = ""
	u.Active = 1

	token, err := token.Generate(u)
	if err != nil {
		logger.Write(err)
		return "", err
	}

	return token, nil
}
