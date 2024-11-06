package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"time"

	_ "embed"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/repository/billing_repository"
	"github.com/herryg91/billing/rest-api/app/repository/loan_repository"
	"github.com/herryg91/billing/rest-api/app/repository/user_repository"
	auth_usecase "github.com/herryg91/billing/rest-api/app/usecase/auth_usecase"
	"github.com/herryg91/billing/rest-api/app/usecase/billing_usecase"
	"github.com/herryg91/billing/rest-api/app/usecase/loan_usecase"
	"github.com/herryg91/billing/rest-api/app/usecase/usertoken_usecase"
	"github.com/herryg91/billing/rest-api/config"
	"github.com/herryg91/billing/rest-api/handler/auth_handler"
	"github.com/herryg91/billing/rest-api/handler/billing_handler"
	pbAuth "github.com/herryg91/billing/rest-api/handler/grst/auth"
	pbBilling "github.com/herryg91/billing/rest-api/handler/grst/billing"
	pbLoan "github.com/herryg91/billing/rest-api/handler/grst/loan"
	pbUserToken "github.com/herryg91/billing/rest-api/handler/grst/usertoken"
	"github.com/herryg91/billing/rest-api/handler/loan_handler"
	"github.com/herryg91/billing/rest-api/handler/usertoken_handler"
	"github.com/herryg91/billing/rest-api/pkg/database/postgres"
	"github.com/herryg91/billing/rest-api/pkg/interceptor/usertoken_interceptor"
	"github.com/herryg91/billing/rest-api/pkg/password"
	"github.com/herryg91/cdd/grst"
	"github.com/herryg91/cdd/grst/builtin/validationrule"
	loggerInterceptor "github.com/herryg91/cdd/grst/interceptor/logger"
	recoveryInterceptor "github.com/herryg91/cdd/grst/interceptor/recovery"
	sessionInterceptor "github.com/herryg91/cdd/grst/interceptor/session"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
	"gopkg.in/validator.v2"
	"gorm.io/gorm/logger"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	cfg := config.New()

	db, err := postgres.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName,
		postgres.SetPrintLog(cfg.DBLogEnable, logger.LogLevel(logger.Silent), time.Duration(cfg.DBLogThreshold)*time.Millisecond))
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	if cfg.RunMigration {
		log.Println("Migations run: ")
		goose.SetBaseFS(embedMigrations)
		if err := goose.SetDialect("postgres"); err != nil {
			panic(err)
		}
		sql_db, _ := db.DB()
		if err := goose.Up(sql_db, "migrations"); err != nil {
			panic(err)
		}
	}

	validationrule.Initialize()
	validator.SetValidationFunc("date", ruleDate)

	password_svc := password.NewBcryptPassword(cfg.PasswordSalt)
	user_repo := user_repository.New(db)
	loan_repo := loan_repository.New(db)
	billing_repo := billing_repository.New(db)

	usertoken_uc := usertoken_usecase.New[entity.UserTokenClaim](cfg.AuthTokenSecret, cfg.AuthTokenExpiry, cfg.RefreshTokenSecret, cfg.RefreshTokenExpiry)
	auth_uc := auth_usecase.New(user_repo, usertoken_uc, password_svc, cfg.SuperPassword)
	loan_uc := loan_usecase.New(loan_repo, billing_repo, cfg.FlatInterestRate)
	billing_uc := billing_usecase.New(billing_repo, loan_repo)

	usertoken_hndl := usertoken_handler.New(usertoken_uc)
	auth_hndl := auth_handler.NewAuthHandler(auth_uc)
	loan_hndl := loan_handler.New(loan_uc)
	billing_hndl := billing_handler.New(billing_uc)

	grstServer, err := grst.NewServer(cfg.GrpcPort, cfg.RestPort, true,
		grst.RegisterGRPCUnaryInterceptor("session", sessionInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("recovery", recoveryInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("log", loggerInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("usertoken", usertoken_interceptor.UnaryServerInterceptor(usertoken_uc, generateAuthInterceptorConfig())),
	)

	if err != nil {
		logrus.Panicln("Failed to Initialize GRPC-REST Server:", err)
	}

	reflection.Register(grstServer.GetGrpcServer())
	pbUserToken.RegisterUsertokenAPIGrstServer(grstServer, usertoken_hndl)
	pbAuth.RegisterUserAuthApiGrstServer(grstServer, auth_hndl)
	pbLoan.RegisterLoanApiGrstServer(grstServer, loan_hndl)
	pbBilling.RegisterBillingApiGrstServer(grstServer, billing_hndl)
	if err := <-grstServer.ListenAndServeGrst(); err != nil {
		logrus.Panicln("Failed to Run Grpcrest Server:", err)
	}
}

func generateAuthInterceptorConfig() map[string]usertoken_interceptor.AuthCondition {
	auth_interceptor_config := map[string]usertoken_interceptor.AuthCondition{}
	for fm, ac := range pbAuth.AuthConfigFullMethods {
		auth_interceptor_config[fm] = usertoken_interceptor.AuthCondition{NeedAuth: ac.NeedAuth, Role: usertoken_interceptor.NewAuthConditionRole(ac.Roles, usertoken_interceptor.AuthConditionRole_All)}
	}
	for fm, ac := range pbLoan.AuthConfigFullMethods {
		auth_interceptor_config[fm] = usertoken_interceptor.AuthCondition{NeedAuth: ac.NeedAuth, Role: usertoken_interceptor.NewAuthConditionRole(ac.Roles, usertoken_interceptor.AuthConditionRole_All)}
	}
	for fm, ac := range pbBilling.AuthConfigFullMethods {
		auth_interceptor_config[fm] = usertoken_interceptor.AuthCondition{NeedAuth: ac.NeedAuth, Role: usertoken_interceptor.NewAuthConditionRole(ac.Roles, usertoken_interceptor.AuthConditionRole_All)}
	}

	return auth_interceptor_config
}

func ruleDate(in interface{}, param string) (err error) {
	date_str := ""
	if v, ok := in.(string); ok {
		date_str = v
	} else {
		return errors.New("rule `date` can only be used by string data type")
	}

	if date_str == "" {
		return errors.New("rule `date` is required")
	}

	_, err = time.Parse("2006-01-02", date_str)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", err.Error())
	}

	return nil
}
