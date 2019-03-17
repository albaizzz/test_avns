package repositories

import (
	"test_avns/apitest/infrastructures/adapter"
	"test_avns/apitest/interfaces"
	"test_avns/apitest/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	config "github.com/spf13/viper"
)

type AuthRepo struct {
	db    adapter.MySQLAdapter
	redis interfaces.IRedis
}

func NewAuthRepo(db adapter.MySQLAdapter, redis interfaces.IRedis) *AuthRepo {
	return &AuthRepo{
		db:    db,
		redis: redis,
	}
}

func (u *AuthRepo) SaveSession(ctx context.Context, userSession models.UserSession) (err error) {
	qry := "insert into user_session (user_id, session, expired_date) values (?,?,?)"
	_, err = u.db.Exec(ctx, qry,
		userSession.UserId, userSession.Session, userSession.ExpiredDate)
	if err != nil {
		return err
	}

	//save data into redis
	redisKey := fmt.Sprintf("user_session:%s", userSession.Session)
	ttl := time.Duration(config.GetInt("redis.cache_duration_session")) * time.Minute
	userData, err := json.Marshal(userSession)
	err = u.redis.Set(redisKey, string(userData), ttl)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthRepo) GetSession(ctx context.Context, session string) (userSession models.UserSession, err error) {
	redisKey := fmt.Sprintf("user_session:%s", session)
	dataSession, err := u.redis.Get(redisKey)
	if err != nil {
		return userSession, err
	}
	if dataSession != nil {
		byteRes := []byte(dataSession.(string))
		err := json.Unmarshal(byteRes, &userSession.User)
		if err != nil {
			return userSession, err
		}
		return userSession, nil
	}

	qry := `select us.user_id, us.session, us.created_date, us.expired_date, u.id, u.username, u.fullname, u.email, u.address  from user_session us 
	inner join user u on us.user_id = u.id 
	where us.session = ?  and UNIX_TIMESTAMP(expired_date) >= UNIX_TIMESTAMP(now())`

	row := u.db.Query(ctx, qry, session)

	err = row.Scan(
		&userSession.UserId,
		&userSession.Session,
		&userSession.CreatedDate,
		&userSession.ExpiredDate,
		&userSession.User.ID,
		&userSession.User.Username,
		&userSession.User.Fullname,
		&userSession.User.Email,
		&userSession.User.Address,
	)
	if err != nil && err != sql.ErrNoRows {
		return userSession, err
	}

	return
}

func (u *AuthRepo) Logout(ctx context.Context, session string) (err error) {

	var errs []error

	muErr := sync.Mutex{}
	errorHandler := func(e error) {
		muErr.Lock()
		errs = append(errs, e)
		muErr.Unlock()
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(session string, ctx context.Context) {
		defer wg.Done()
		qry := "delete from user_session where session =?"
		_, err := u.db.Exec(ctx, qry, session)
		if err != nil {
			errorHandler(err)
		}
	}(session, ctx)

	go func(session string) {
		defer wg.Done()
		redisKey := fmt.Sprintf("user_session:%s", session)
		err := u.redis.Del(redisKey)
		if err != nil {
			errorHandler(err)
		}
	}(session)

	wg.Wait()

	if len(errs) > 0 {
		err = errs[0]
		return
	}

	return nil
}
