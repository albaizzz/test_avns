package healthcheck

import (
	"fmt"
	"net/http"
	"sync"
	"test_avns/apitest/interfaces"
	"time"
)

type serviceItem struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Remark      string `json:"remark"`
	StatusCode  int    `json:"-"`
}

type Check func() error

const DefaultTimeoutCheck = 1500

type HealthService struct {
	readiness   map[string]serviceItem
	healthMutex sync.RWMutex
	dbMaster    interfaces.DatabasePing
	dbSlave     interfaces.DatabasePing
	redisCache  interfaces.DatabasePing
}

func NewHealthService(dbMaster, dbSlave, redisCache interfaces.DatabasePing) *HealthService {
	return &HealthService{
		dbMaster:   dbMaster,
		dbSlave:    dbSlave,
		redisCache: redisCache,
		readiness:  make(map[string]serviceItem),
	}
}

func (s *HealthService) HealthStatus() (httpStatusCode int, result interface{}) {

	wg := sync.WaitGroup{}

	s.addReadiness("db_test_avns/apitest_master", func() error {
		if s.dbMaster == nil {
			return fmt.Errorf("database is nil")
		}
		return s.dbMaster.Ping()

	}, &wg)

	s.addReadiness("db_test_avns/apitest_slave", func() error {
		if s.dbSlave == nil {
			return fmt.Errorf("database is nil")
		}
		return s.dbMaster.Ping()

	}, &wg)

	s.addReadiness("redis", func() error {
		if s.redisCache == nil {
			return fmt.Errorf("redis is nil")
		}
		return s.redisCache.Ping()

	}, &wg)

	wg.Wait()

	return s.collect(), s.readiness
}

func (s *HealthService) collect() int {
	s.healthMutex.RLock()
	defer s.healthMutex.RUnlock()

	status := http.StatusOK

	for _, x := range s.readiness {

		if x.StatusCode != http.StatusOK {
			status = x.StatusCode
		}
	}

	return status

}

func (s *HealthService) addService(svc serviceItem) {

	s.healthMutex.Lock()
	defer s.healthMutex.Unlock()

	s.readiness[svc.ServiceName] = svc
}

func (s *HealthService) addReadiness(name string, check Check, wg *sync.WaitGroup) {

	wg.Add(1)
	go func() {

		i := serviceItem{
			Status:      "OK",
			ServiceName: name,
			Remark:      "",
			StatusCode:  http.StatusOK,
		}

		if err := check(); err != nil {
			i.StatusCode = http.StatusServiceUnavailable
			i.Remark = err.Error()
			i.Status = "unhealthy"
		}

		s.addService(i)

		wg.Done()
	}()
}

func (s *HealthService) HTTPHealthyCheck(url string, timeout time.Duration) Check {
	cl := http.DefaultClient
	cl.Timeout = timeout
	cl.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return func() error {
		resp, err := cl.Get(url)
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("returned status %d", resp.StatusCode)
		}
		return nil
	}
}
