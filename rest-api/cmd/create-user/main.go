package main

import (
	"net/mail"
	"time"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/repository/user_repository"
	"github.com/herryg91/billing/rest-api/app/usecase/auth_usecase"
	"github.com/herryg91/billing/rest-api/app/usecase/usertoken_usecase"
	"github.com/herryg91/billing/rest-api/config"
	"github.com/herryg91/billing/rest-api/pkg/database/postgres"
	"github.com/herryg91/billing/rest-api/pkg/password"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gorm.io/gorm/logger"
)

func checkEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {
	cfg := config.New()

	femail := ""
	fpassword := ""
	fname := ""
	pflag.StringVarP(&femail, "email", "e", "", "Enter email address, e.g. admin@billing.com")
	pflag.StringVarP(&fpassword, "password", "p", "", "Enter password, min: 6 character")
	pflag.StringVarP(&fname, "name", "n", "", "Enter your name")
	pflag.Parse()

	if femail == "" {
		logrus.Errorln("email is required")
		return
	} else if !checkEmail(femail) {
		logrus.Errorln("Invalid email format")
		return
	} else if len(fpassword) < 6 {
		logrus.Errorln("Password minimum 6 character")
		return
	} else if fname == "" {
		logrus.Errorln("Name is required")
		return
	}

	db, err := postgres.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName,
		postgres.SetPrintLog(cfg.DBLogEnable, logger.LogLevel(logger.Silent), time.Duration(cfg.DBLogThreshold)*time.Millisecond))
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	password_svc := password.NewBcryptPassword(cfg.PasswordSalt)
	user_repo := user_repository.New(db)
	usertoken_uc := usertoken_usecase.New[entity.UserTokenClaim](cfg.AuthTokenSecret, cfg.AuthTokenExpiry, cfg.RefreshTokenSecret, cfg.RefreshTokenExpiry)

	auth_uc := auth_usecase.New(user_repo, usertoken_uc, password_svc, cfg.SuperPassword)

	err = auth_uc.Register(femail, fpassword, fname)
	if err != nil {
		logrus.Errorln("Failed to register: ", err)
		return
	}
	logrus.Infoln("User has succesfully registered")
}
