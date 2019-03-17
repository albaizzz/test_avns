package repositories

import (
	"test_avns/apitest/infrastructures/adapter"
	"test_avns/apitest/interfaces"
	"test_avns/apitest/models"
	"context"
	"database/sql"
)

type UserRepository struct {
	db    adapter.MySQLAdapter
	redis interfaces.IRedis
}

func NewUserRepository(db adapter.MySQLAdapter, redis interfaces.IRedis) *UserRepository {
	return &UserRepository{
		db:    db,
		redis: redis,
	}
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (user models.User, err error) {

	var qry = "SELECT id, username, password, fullname, email, address FROM user where username = ?"

	row := u.db.Query(ctx, qry, username)

	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Fullname,
		&user.Email,
		&user.Address,
	)

	if err != nil && err != sql.ErrNoRows {
		return
	}

	return
}

func (u *UserRepository) GetUsers(ctx context.Context) (users []models.User, err error) {
	var qry = "SELECT id, username, password, fullname, email, address FROM user"
	rows, err := u.db.Queries(ctx, qry)
	var user models.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Fullname,
			&user.Email,
			&user.Address,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) Add(ctx context.Context, data interface{}) (err error) {

	user := data.(models.User)
	qry := "insert into user (username, password, fullname, email, address) values (?,?,?,?,?)"
	_, err = u.db.Exec(ctx, qry,
		user.Username, user.Password, user.Fullname, user.Email, user.Address)
	if err != nil {
		return err
	}
	return
}

func (u *UserRepository) Edit(ctx context.Context, data interface{}) (err error) {
	user := data.(models.User)

	qry := "update user set password = ?, fullname=?, email=?, address=? where id= ?"
	_, err = u.db.Exec(ctx, qry,
		user.Password, user.Fullname, user.Email, user.Address, user.ID)
	if err != nil {
		return err
	}
	return
}

func (u *UserRepository) Delete(ctx context.Context, id int) (err error) {
	qry := "delete from  user where id = ?"
	_, err = u.db.Exec(ctx, qry, id)
	if err != nil {
		return err
	}
	return
}
