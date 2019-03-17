package routes

import (
	"time"

	"test_avns/apitest/infrastructures"
	"test_avns/apitest/infrastructures/adapter"
	"test_avns/apitest/interfaces"

	muxtrace "github.com/DataDog/dd-trace-go/contrib/gorilla/mux"
	"github.com/DataDog/dd-trace-go/tracer"
	"github.com/gorilla/mux"
	config "github.com/spf13/viper"
	"github.com/urfave/negroni"
)

type Route struct {
	Database       interfaces.ISQL
	Redis          interfaces.IRedis
	MySQLConnector adapter.MySQLAdapter
}

func (r *Route) Init() *mux.Router {

	if config.GetBool("app.debug") == true {
		tracer.DefaultTracer.SetDebugLogging(true)
	}

	//Initialize controller
	healthController := r.InitHealthController()
	userController := r.InitUserHander()
	authenticator := r.InitAuthenticator()

	router := mux.NewRouter()
	muxTracer := muxtrace.NewMuxTracer(config.GetString("app.name"), tracer.DefaultTracer)
	internal := router.PathPrefix("/v1").Subrouter()

	var exBase = mux.NewRouter()
	router.PathPrefix("/v1/ex").Handler(negroni.New(
		negroni.HandlerFunc(authenticator.Authenticate),
		negroni.Wrap(exBase),
	))

	var va = exBase.PathPrefix("/v1/ex").Subrouter().StrictSlash(true)

	muxTracer.HandleFunc(va, "/user/{username}", userController.GetUserByUsername).Methods("GET")
	muxTracer.HandleFunc(va, "/user/{id}", userController.EditUser).Methods("PUT")
	muxTracer.HandleFunc(va, "/user/{id}", userController.DeleteUser).Methods("DELETE")
	muxTracer.HandleFunc(va, "/user/logout", userController.Logout).Methods("POST")

	muxTracer.HandleFunc(internal, "/user/register", userController.Register).Methods("GET")
	muxTracer.HandleFunc(internal, "/user/auth", userController.AuthenticatUser).Methods("GET")

	// Health
	muxTracer.HandleFunc(internal, "/health/status", healthController.GetStatus).Methods("GET")
	muxTracer.HandleFunc(internal, "/health/liveness", healthController.GetStatus).Methods("GET")
	muxTracer.HandleFunc(internal, "/health/readiness", healthController.Readiness).Methods("GET")

	return router
}

// InitSQL set the sql values
func InitSQL() *infrastructures.SQLInfrastructure {
	sqlInfra := new(infrastructures.SQLInfrastructure)

	sqlInfra.SQLSlave.User = config.GetString("database_sql_slave.user")
	sqlInfra.SQLSlave.Password = config.GetString("database_sql_slave.password")
	sqlInfra.SQLSlave.Host = config.GetString("database_sql_slave.host")
	sqlInfra.SQLSlave.Port = config.GetInt("database_sql_slave.port")
	sqlInfra.SQLSlave.DBName = config.GetString("database_sql_slave.db_name")
	sqlInfra.SQLSlave.Charset = config.GetString("database_sql_slave.charset")

	sqlInfra.SQLMaster.User = config.GetString("database_sql_master.user")
	sqlInfra.SQLMaster.Password = config.GetString("database_sql_master.password")
	sqlInfra.SQLMaster.Host = config.GetString("database_sql_master.host")
	sqlInfra.SQLMaster.Port = config.GetInt("database_sql_master.port")
	sqlInfra.SQLMaster.DBName = config.GetString("database_sql_master.db_name")
	sqlInfra.SQLMaster.Charset = config.GetString("database_sql_master.charset")

	sqlInfra.OpenConnection()
	sqlInfra.SetConnMaxLifetime(config.GetDuration("database.max_life_time") * time.Second)
	sqlInfra.SetMaxIdleConns(config.GetInt("database.max_idle_connection"))
	sqlInfra.SetMaxOpenConns(config.GetInt("database.max_open_connection"))

	return sqlInfra
}
