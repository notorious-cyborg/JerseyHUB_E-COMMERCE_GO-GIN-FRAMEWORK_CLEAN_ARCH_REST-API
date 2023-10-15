// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"jerseyhub/pkg/api"
	"jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/config"
	"jerseyhub/pkg/db"
	"jerseyhub/pkg/helper"
	"jerseyhub/pkg/repository"
	"jerseyhub/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	helper:=helper.NewHelper(cfg)

	offerRepository := repository.NewOfferRepository(gormDB)
	offerUseCase := usecase.NewOfferUseCase(offerRepository)
	offerHandler := handler.NewOfferHandler(offerUseCase)

	wishlistRepository := repository.NewWishlistRepository(gormDB)
	wishlistUseCase := usecase.NewWishlistUseCase(wishlistRepository,offerRepository)
	wishlistHandler := handler.NewWishlistHandler(wishlistUseCase)


	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository,helper)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	inventoryRepository := repository.NewInventoryRepository(gormDB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository,offerRepository,helper)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository,helper)
	otpHandler := handler.NewOtpHandler(otpUseCase)


	orderRepository := repository.NewOrderRepository(gormDB)

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository,cfg,otpRepository,inventoryRepository,orderRepository,helper)
	userHandler := handler.NewUserHandler(userUseCase)

	couponRepository := repository.NewCouponRepository(gormDB)
	couponUseCase := usecase.NewCouponUseCase(couponRepository)
	couponHandler := handler.NewCouponHandler(couponUseCase)

	orderUseCase := usecase.NewOrderUseCase(orderRepository,couponRepository,userUseCase)
	orderHandler := handler.NewOrderHandler(orderUseCase)


	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository,inventoryRepository,userUseCase)
	cartHandler := handler.NewCartHandler(cartUseCase)


	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	
	serverHTTP := http.NewServerHTTP(userHandler,adminHandler,categoryHandler,inventoryHandler,otpHandler,orderHandler,cartHandler,couponHandler,paymentHandler,offerHandler,wishlistHandler)



	return serverHTTP, nil
}
