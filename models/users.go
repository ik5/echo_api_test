package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ik5/echo_api_test/structs"
	"github.com/ik5/echo_api_test/utils/runtimeutils"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	// Basic methods
	Model      `json:"-"`
	ID         int       `json:"id,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	MiddleName string    `json:"middle_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Username   string    `json:"username,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

const (
	sqlInsertUser string = `
insert into users(first_name, middle_name,last_name, username)
values($1, COALESCE($2, ''), $3, $4)
returning *
`
	sqlFindUserByID string = `select * from users where id=$1`
)

func (user *Users) Insert(userData *structs.User) error {
	funcName := runtimeutils.GetCallerFunctionName()

	trx, err := user.db.Begin(context.Background())
	if err != nil {
		user.logger.Debug("Cannot start transaction", "err", err)

		return fmt.Errorf("[%s] %w", funcName, err)
	}

	row := trx.QueryRow(
		context.Background(), sqlInsertUser,
		userData.FirstName, userData.MiddleName, userData.LastName, userData.Username,
	)

	err = row.Scan(
		&user.ID, &user.FirstName, &user.MiddleName, &user.LastName,
		&user.Username, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		user.logger.Debug("unable to scan result", "err", err)

		_ = trx.Rollback(context.Background())

		return fmt.Errorf("[%s] %w", funcName, err)
	}

	err = trx.Commit(context.Background())
	if err != nil {
		user.logger.Debug("Unable to commit result", "err", err)

		return fmt.Errorf("[%s] %w", funcName, err)
	}

	return nil
}

func (user *Users) FindByID(userID int) error {
	err := user.db.QueryRow(
		context.Background(), sqlFindUserByID, userID,
	).Scan(
		&user.ID, &user.FirstName, &user.MiddleName, &user.LastName,
		&user.Username, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		user.logger.Debug("Unable to scan row", "err", err)

		if errors.Is(err, pgx.ErrNoRows) {
			user.logger.Info("User was not found", "id", userID)

			return nil
		}

		funcName := runtimeutils.GetCallerFunctionName()

		return fmt.Errorf("[%s] %w", funcName, err)
	}

	return nil
}
