package routes

import (
	"test_avns/apitest/services"
	"fmt"
	"time"

	"test_avns/apitest/handlers/healthcheck"
	"test_avns/apitest/handlers/user"
	"test_avns/apitest/infrastructures"
	"test_avns/apitest/infrastructures/adapter"
	"test_avns/apitest/repositories"

	config "github.com/spf13/viper"
)

// InitHealthController set the health controller dependency
func (r *Route) InitHealthController() *healthcheck.HealthCheckHandler {
	healthController := new(healthcheck.HealthCheckHandler)

	dbMaster, _ := r.Database.GetDBMaster()
	dbSlave, _ := r.Database.GetDBSlave()

	healthController.HealthService = healthcheck.NewHealthService(dbMaster, dbSlave, r.Redis)

	return healthController
}

func (r *Route) InitUserHander() *user.UserHandler {
	userHandler := new(user.UserHandler)
	redis := InitRedis()
	userRepo := repositories.NewUserRepository(r.MySQLConnector, redis)
	authRepo := repositories.NewAuthRepo(r.MySQLConnector, redis)
	userService := user.NewUserService(userRepo, authRepo)
	userHandler.UserService = userService
	return userHandler
}

func (r *Route) InitAuthenticator() *services.Authenticator {
	redis := InitRedis()
	authRepo := repositories.NewAuthRepo(r.MySQLConnector, redis)
	authenticator := services.NewAuthenticator(authRepo)
	return authenticator
}

// InitRedis set the redis values
func InitRedis() *infrastructures.Redis {
	redis := new(infrastructures.Redis)
	redis.Host = config.GetString("redis.host")

	// redis.Password = config.GetString("redis.password")
	redis.DB = config.GetInt("redis.db")
	redis.Port = config.GetInt("redis.port")

	return redis
}

// SetupMySQL config
func SetupMySQL() *adapter.MySQLConfig {

	return &adapter.MySQLConfig{
		User:            config.GetString("database_sql_master.user"),
		Password:        config.GetString("database_sql_master.password"),
		Host:            fmt.Sprintf("%s:%d", config.GetString("database_sql_master.host"), config.GetInt("database_sql_master.port")),
		Name:            config.GetString("database_sql_master.db_name"),
		Charset:         config.GetString("database_sql_master.charset"),
		MaxOpenConns:    config.GetInt("database.max_open_connection"),
		MaxIdleConns:    config.GetInt("database.max_idle_connection"),
		ConnMaxLifetime: config.GetDuration("database.max_life_time") * time.Second,
		Timeout:         config.GetDuration("database.timeout") * time.Second,
	}

}
