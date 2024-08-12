package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/casbin"
	"api-gateway/config"
	"api-gateway/genproto/health_analytics"
	"api-gateway/genproto/user"
	"api-gateway/pkg/logger"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf := config.Load()
	hand := NewHandler()
	
	router := api.NewRouter(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(conf.API_GATEWAY))
}

func NewHandler() *handler.Handler {
	conf := config.Load()
	health, err := grpc.Dial(conf.HEALTH_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	userr, err := grpc.NewClient(conf.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	users:=user.NewUserServiceClient(userr)
	healthh := health_analytics.NewHealthAnalyticsServiceClient(health)

	logs := logger.NewLogger()
	en, err := casbin.CasbinEnforcer(logs)
	if err != nil {
		log.Fatal("error in creating casbin enforcer", err)
	}
	
	return &handler.Handler{
		HealthService: healthh,
		UserService: users,
		Log: logger.NewLogger(),
		Enforcer: en,
	} 
}
