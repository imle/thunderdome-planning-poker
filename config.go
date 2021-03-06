package main

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/thunderdome/")
	viper.AddConfigPath("$HOME/.config/thunderdome/")
	viper.AddConfigPath(".")

	viper.SetDefault("http.cookie_hashkey", "strongest-avenger")
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.secure_cookie", true)
	viper.SetDefault("http.domain", "thunderdome.dev")
	viper.SetDefault("analytics.enabled", true)
	viper.SetDefault("analytics.id", "UA-140245309-1")
	viper.SetDefault("db.host", "db")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.user", "thor")
	viper.SetDefault("db.pass", "odinson")
	viper.SetDefault("db.name", "thunderdome")
	viper.SetDefault("db.ssl", false)
	viper.SetDefault("smtp.host", "localhost")
	viper.SetDefault("smtp.port", "25")
	viper.SetDefault("smtp.secure", true)
	viper.SetDefault("smtp.sender", "no-reply@thunderdome.dev")
	viper.SetDefault("config.allowedPointValues",
		[]string{"0", "1/2", "1", "2", "3", "5", "8", "13", "20", "40", "100", "?"})
	viper.SetDefault("config.defaultPointValues",
		[]string{"1", "2", "3", "5", "8", "13", "?"})
	viper.SetDefault("config.show_warrior_rank", false)
	viper.SetDefault("config.avatar_service", "default")
	viper.SetDefault("config.toast_timeout", 1000)
	viper.SetDefault("config.allow_guests", true)
	viper.SetDefault("config.allow_registration", true)

	viper.BindEnv("http.cookie_hashkey", "COOKIE_HASHKEY")
	viper.BindEnv("http.port", "PORT")
	viper.BindEnv("http.secure_cookie", "COOKIE_SECURE")
	viper.BindEnv("http.domain", "APP_DOMAIN")
	viper.BindEnv("analytics.enabled", "ANALYTICS_ENABLED")
	viper.BindEnv("analytics.id", "ANALYTICS_ID")
	viper.BindEnv("admin.email", "ADMIN_EMAIL")
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.user", "DB_USER")
	viper.BindEnv("db.pass", "DB_PASS")
	viper.BindEnv("db.name", "DB_NAME")
	viper.BindEnv("db.ssl", "DB_SSL")
	viper.BindEnv("smtp.host", "SMTP_HOST")
	viper.BindEnv("smtp.port", "SMTP_PORT")
	viper.BindEnv("smtp.secure", "SMTP_SECURE")
	viper.BindEnv("smtp.identity", "SMTP_IDENTITY")
	viper.BindEnv("smtp.user", "SMTP_USER")
	viper.BindEnv("smtp.pass", "SMTP_PASS")
	viper.BindEnv("smtp.sender", "SMTP_SENDER")
	viper.BindEnv("config.allowedPointValues", "CONFIG_POINTS_ALLOWED")
	viper.BindEnv("config.defaultPointValues", "CONFIG_POINTS_DEFAULT")
	viper.BindEnv("config.show_warrior_rank", "CONFIG_SHOW_RANK")
	viper.BindEnv("config.avatar_service", "CONFIG_AVATAR_SERVICE")
	viper.BindEnv("config.toast_timeout", "CONFIG_TOAST_TIMEOUT")
	viper.BindEnv("config.allow_guests", "CONFIG_ALLOW_GUESTS")
	viper.BindEnv("config.allow_registration", "CONFIG_ALLOW_REGISTRATION")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}
}
