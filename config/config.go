package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE   string
	HEALTH_SERVICE string
	SIGNING_KEY    string
	API_GATEWAY    string
}

func Load() Config {
	if err := SignKeyniUzgartirish(); err != nil {
		log.Fatalf("SIGNING_KEY ni yangilashda xatolik: %v", err)
	}
	
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":50051"))
	config.HEALTH_SERVICE = cast.ToString(Coalesce("HEALTH_SERVICE", ":50052"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "lingualeap"))
	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", ":50053"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func SignKeyniUzgartirish() error {
	userServiceEnvPath := "C:/imtixon4oy/UserService/.env"
	if err := godotenv.Load(userServiceEnvPath); err != nil {
		log.Fatal(".env faylini yuklashda xatolik yuz berdi")
		return err
	}

	signingKey := os.Getenv("SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("SIGNING_KEY qiymati topilmadi")
		return fmt.Errorf("SIGNING_KEY qiymati topilmadi")
	}

	apiGatewayEnvPath := ".env"
	if err := godotenv.Load(apiGatewayEnvPath); err != nil {
		log.Fatal(".env faylini yuklashda xatolik yuz berdi")
		return err
	}

	err := godotenv.Write(map[string]string{
		"USER_SERVICE":   os.Getenv("USER_SERVICE"),
		"HEALTH_SERVICE": os.Getenv("HEALTH_SERVICE"),
		"API_GATEWAY":    os.Getenv("API_GATEWAY"),
		"SIGNING_KEY":    signingKey,
	}, apiGatewayEnvPath)

	if err != nil {
		log.Fatalf(".env fayliga SIGNING_KEY yozishda xatolik yuz berdi: %v", err)
	}
	return nil
}
