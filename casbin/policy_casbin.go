package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

const (
	host     = "localhost"
	port     = "5432"
	dbname   = "casbin"
	username = "postgres"
	password = "salom"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS casbin")
	if err != nil {
		logger.Error("Error dropping Casbin database", "error", err.Error())
		return nil, err
	}

	adapter, err := xormadapter.NewAdapter("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password))
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}
// patient, doctor, admin
	policies := [][]string{
		{"doctor", "/health/medical_recordsAdd", "POST"},
		{"admin", "/health/medical_recordsAdd", "POST"},
		{"doctor", "/health/medical_recordsGet/:id", "GET"},
		{"admin", "/health/medical_recordsGet/:id", "GET"},
		{"patient", "/health/medical_recordsGet/:id", "GET"},
		{"doctor", "/health/medical_recordsUp", "PUT"},
		{"admin", "/health/medical_recordsUp", "PUT"},
		{"patient", "/health/medical_recordsUp", "PUT"},
		{"admin", "/health/medical_recordsDel/:id", "DELETE"},
		{"admin", "/health/medical_records/user/:userId", "GET"},

		{"patient", "/health/lifestyleAdd", "POST"},
		{"patient", "/health/getalllifestyledata/:limit/:page", "GET"},
		{"patient", "/health/lifestyleGet/:id", "GET"},
		{"patient", "/health/lifestyleUp", "PUT"},
		{"patient", "/health/lifestyleDel/:id", "DELETE"},

		{"patient", "/health/wearable-dataAdd", "POST"},
		{"patient", "/health/wearabledata/:limit/:page", "GET"},
		{"patient", "/health/wearable-dataGet/:id", "GET"},
		{"patient", "/health/wearable-dataUp", "PUT"},
		{"patient", "/health/wearable-dataDel/:id", "DELETE"},

		{"doctor", "/health/recommendationsAdd", "POST"},
		{"patient", "/health/monitoring/:user_id/realtime", "GET"},
		{"patient", "/health/summary/:user_id/daily/:date", "GET"},
		{"patient", "/health/summary/:user_id/weekly/:start_date", "GET"},
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}
	return enforcer, nil
}