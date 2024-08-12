package handler

import (
	"api-gateway/genproto/health_analytics"
	"api-gateway/genproto/user"
	"log/slog"
	"github.com/casbin/casbin/v2"
)

type Handler struct {
	HealthService health_analytics.HealthAnalyticsServiceClient
	UserService   user.UserServiceClient
	Log           *slog.Logger
	Enforcer     *casbin.Enforcer
}
