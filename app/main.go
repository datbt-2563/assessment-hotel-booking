package main

import (
	"fmt"
	"hotel-booking/config"

	reservationHttp "hotel-booking/reservation/delivery/http"
	reservation "hotel-booking/reservation/provider"

	customerHttp "hotel-booking/customer/delivery/http"
	customer "hotel-booking/customer/provider"

	roomHttp "hotel-booking/room/delivery/http"
	room "hotel-booking/room/provider"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// configuration
	viper.SetDefault(config.MySqlUsernameConfig, "user")
	viper.SetDefault(config.MySqlPasswordConfig, "password")
	viper.SetDefault(config.MySqlAddrConfig, "localhost:3306")
	viper.SetDefault(config.MySqlDatabaseName, "database")

	viper.SetDefault(config.HttpHandlerPort, "8020")
	viper.SetDefault(config.Workdir, "/app/")

	// viper.SetEnvPrefix("hotel-booking")
	viper.AutomaticEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.Get(config.MySqlUsernameConfig), viper.Get(config.MySqlPasswordConfig), viper.Get(config.MySqlAddrConfig), viper.Get(config.MySqlDatabaseName))
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// DI

	customerRepo := customer.NewCustomerRepository(dbClient)
	customerUsecase := customer.NewCustomerUsecase(customerRepo)

	roomRepo := room.NewRoomRepository(dbClient)
	roomUsecase := room.NewRoomUsecase(roomRepo)

	reservationRepo := reservation.NewReservationRepository(dbClient)
	reservationUsecase := reservation.NewReservationUsecase(reservationRepo)

	// run http history handler
	log.Println("Create HTTP handler ...")
	httpPort := viper.Get(config.HttpHandlerPort)
	r := gin.Default()

	customerHttp.AddCustomerHandler(r, customerUsecase)
	roomHttp.AddRoomHandler(r, roomUsecase)
	reservationHttp.AddReservationHandler(r, reservationUsecase)

	err = r.Run(fmt.Sprintf(":%s", httpPort))
	if err != nil {
		log.Fatalln(err)
	}
}
