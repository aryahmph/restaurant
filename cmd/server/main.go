package main

import (
	"fmt"
	"github.com/aryahmph/restaurant/common/env"
	httpCommon "github.com/aryahmph/restaurant/common/http"
	jwtCommon "github.com/aryahmph/restaurant/common/jwt"
	passCommon "github.com/aryahmph/restaurant/common/password"
	dbCommon "github.com/aryahmph/restaurant/common/postgresql"
	authDelivery "github.com/aryahmph/restaurant/internal/delivery/auth/http"
	categoryDelivery "github.com/aryahmph/restaurant/internal/delivery/category/http"
	userDelivery "github.com/aryahmph/restaurant/internal/delivery/user/http"
	categoryRepo "github.com/aryahmph/restaurant/internal/repository/category/postgresql"
	userRepo "github.com/aryahmph/restaurant/internal/repository/user/postgresql"
	authUc "github.com/aryahmph/restaurant/internal/usecase/auth"
	categoryUc "github.com/aryahmph/restaurant/internal/usecase/category"
	userUc "github.com/aryahmph/restaurant/internal/usecase/user"
	"github.com/gin-contrib/cors"
	"log"
)

func main() {
	cfg := env.LoadConfig()
	db := dbCommon.NewPostgreSql(cfg.PostgresURL)
	httpServer := httpCommon.NewHTTPServer()
	passwordManager := passCommon.NewPasswordHashManager()
	jwtManager := jwtCommon.NewJWTManager(cfg.AccessTokenKey)

	httpServer.Router.Use(httpCommon.MiddlewareErrorHandler())
	httpServer.Router.Use(cors.Default())
	httpServer.Router.RedirectTrailingSlash = true
	root := httpServer.Router.Group("/api")

	userRepository := userRepo.NewPostgreSqlUserRepositoryImpl(db)
	userUsecase := userUc.NewUserUsecaseImpl(userRepository, passwordManager, jwtManager)
	userDelivery.NewHTTPUserDelivery(root.Group("/users"), userUsecase, jwtManager)

	authUsecase := authUc.NewAuthUsecase(userRepository, passwordManager, jwtManager)
	authDelivery.NewHTTPAuthDelivery(root.Group("/auth"), authUsecase)

	categoryRepository := categoryRepo.NewPostgreSqlCategoryRepositoryImpl(db)
	categoryUsecase := categoryUc.NewCategoryUsecaseImpl(userRepository, categoryRepository)
	categoryDelivery.NewHTTPCategoryDelivery(root.Group("/categories"), categoryUsecase, jwtManager)

	log.Fatalln(httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
