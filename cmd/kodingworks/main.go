package main

import (
	guests "kodingworks/guests"
	hotels "kodingworks/hotels"
	orderGuests "kodingworks/order_guests"
	orderItems "kodingworks/order_items"
	orders "kodingworks/orders"
	roomAvailabality "kodingworks/room_availabilities"
	roomRates "kodingworks/room_rates"
	rooms "kodingworks/rooms"
	utils "kodingworks/utils"
	"os"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
)

type conf struct {
	AppPort int
	DbPort  int
	DbHost  string
	DbUser  string
	DbPass  string
	DbName  string
}

var config *conf

func main() {

	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	mode := ""
	if len(os.Getenv("LOCAL")) > 0 {
		mode = "local"
	} else {
		mode = "heroku"
	}

	c := viper.Sub(mode)
	// Configurations
	config = &conf{
		AppPort: 8000,
		DbPort:  c.GetInt("DB_PORT"),
		DbHost:  c.GetString("DB_HOST"),
		DbUser:  c.GetString("DB_USER"),
		DbPass:  c.GetString("DB_PASS"),
		DbName:  c.GetString("DB_NAME"),
	}

	log.Println("Server is running on port " + strconv.Itoa(config.AppPort))
	log.Println("Configurations:")
	log.Println("App Port:\t ", config.AppPort)
	log.Println("DB Host :\t ", config.DbHost)
	log.Println("DB Port :\t ", config.DbPort)
	log.Println("DB Name :\t ", config.DbName)
	log.Println("DB User :\t ", config.DbUser)
	log.Println("----------------------------")
	log.Println("Registered Apps:")

	// Init DB
	dbConnPool := initDb(config)

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET") // root route
	api := router.PathPrefix("/api/v1").Subrouter()     // Set api route prefix

	// Registering Routes
	guestAPI := guests.GuestsApi{
		Router: api,
		Db:     dbConnPool,
	}
	guestAPI.Register()

	hotelAPI := hotels.HotelsApi{Router: api, Db: dbConnPool}
	hotelAPI.Register()

	ordersAPI := orders.OrdersApi{Router: api, Db: dbConnPool}
	ordersAPI.Register()

	orderItemsAPI := orderItems.OrderItemsApi{Router: api, Db: dbConnPool}
	orderItemsAPI.Register()

	orderGuestsAPI := orderGuests.OrderGuestSApi{Router: api, Db: dbConnPool}
	orderGuestsAPI.Register()

	roomsAPI := rooms.RoomsApi{Router: api, Db: dbConnPool}
	roomsAPI.Register()

	roomRatesAPI := roomRates.RoomRatesApi{Router: api, Db: dbConnPool}
	roomRatesAPI.Register()

	roomAvailabilitiesAPI := roomAvailabality.RoomAvailabilitiesApi{Router: api, Db: dbConnPool}
	roomAvailabilitiesAPI.Register()

	// Serve

	log.Fatal(http.ListenAndServe(
		":"+strconv.Itoa(config.AppPort),
		handlers.CORS(
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization",
				"Cache-Control",
			}),
			handlers.ExposedHeaders([]string{"Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(router)))
}

func initDb(c *conf) (dbConnPool *pgx.ConnPool) {
	var err error
	pgxConf := &pgx.ConnConfig{
		Port:     uint16(c.DbPort),
		Host:     c.DbHost,
		User:     c.DbUser,
		Password: c.DbPass,
		Database: c.DbName,
	}

	dbConnPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     *pgxConf,
		MaxConnections: 5,
	})
	if err != nil {
		panic(err)
	}
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondwithJSON(
		w,
		http.StatusOK,
		map[string]interface{}{
			"message": "Kodingworks test...",
		},
	)
}
