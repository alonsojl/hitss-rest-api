package user

import (
	"context"
	"database/sql"
	"hitss/internal/db/mysql"
	"hitss/internal/model"
	"hitss/pkg/helper/logger"
	"hitss/pkg/helper/password"
)

type store struct {
	storage *sql.DB
}

func New() *store {
	return &store{
		storage: mysql.Open(),
	}
}

func (s *store) GetAll() (model.Users, error) {
	defer s.storage.Close()

	ctx := context.Background()
	var users model.Users
	stmt, err := s.storage.PrepareContext(ctx, "CALL spUsersGetAll()")
	if err != nil {
		logger.Write(err)
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		logger.Write(err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Tag, &user.Active)
		if err != nil {
			logger.Write(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *store) Create(u *model.User) error {
	defer s.storage.Close()

	ctx := context.Background()
	stmt, err := s.storage.PrepareContext(ctx, "CALL spUserCreate(?,?,?,?,?)")
	if err != nil {
		logger.Write(err)
		return err
	}
	defer stmt.Close()

	u.Password, err = password.Encrypt(u.Password)
	if err != nil {
		logger.Write(err)
		return err
	}
	u.Active = 1
	err = stmt.QueryRowContext(ctx, u.Name, u.Email, u.Password, u.Tag, u.Active).Scan(&u.Id)
	if err != nil {
		logger.Write(err)
		return err
	}

	return nil
}

func (s *store) Update(u *model.User) error {
	defer s.storage.Close()

	ctx := context.Background()
	stmt, err := s.storage.PrepareContext(ctx, "CALL spUserUpdate(?,?,?,?,?,?)")
	if err != nil {
		logger.Write(err)
		return err
	}
	defer stmt.Close()

	u.Password, err = password.Encrypt(u.Password)
	if err != nil {
		logger.Write(err)
		return err
	}
	_, err = stmt.ExecContext(ctx, u.Id, u.Name, u.Email, u.Password, u.Tag, u.Active)
	if err != nil {
		logger.Write(err)
		return err
	}

	return nil
}

func (s *store) Delete(id int) error {
	defer s.storage.Close()

	ctx := context.Background()
	stmt, err := s.storage.PrepareContext(ctx, "CALL spUserDelete(?)")
	if err != nil {
		logger.Write(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		logger.Write(err)
		return err
	}

	return nil
}
