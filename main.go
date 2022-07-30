package main

import (
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"net/http"

	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"

	httpHandler "jamtangan/handler/http"
	"jamtangan/repository/brand"
	"jamtangan/repository/product"
	"jamtangan/repository/transaction"
	"jamtangan/usecase/admin"
	"jamtangan/usecase/customer"
)

// @title Jam Tangan API
// @version 0.1
// @description Jam Tangan API.
// @host localhost:8000

//go:embed seed/mysql/*.sql
var embedMySQLMigrations embed.FS

func main() {
	var command = flag.String("command", "http", "command (default: http)")
	var gooseDown = flag.Int64("goose-down", 0, "goose down command (default: 0)")
	flag.Parse()

	viper.SetConfigFile("config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	snowflakeNodeCfg := viper.GetInt64("snowflake.node")
	snowflakeNode, err := snowflake.NewNode(snowflakeNodeCfg)
	if err != nil {
		panic(err)
	}

	sqlHost := viper.GetString(`sql.host`)
	sqlPort := viper.GetString(`sql.port`)
	sqlUser := viper.GetString(`sql.user`)
	sqlPass := viper.GetString(`sql.password`)
	sqlDBName := viper.GetString(`sql.db_name`)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", sqlUser, sqlPass, sqlHost, sqlPort, sqlDBName)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	if command != nil && *command == "migrate" {
		goose.SetBaseFS(embedMySQLMigrations)

		if err = goose.SetDialect("mysql"); err != nil {
			panic(err)
		}

		if gooseDown != nil && *gooseDown != 0 {
			if err = goose.DownTo(sqlDB, "seed/mysql", *gooseDown); err != nil {
				panic(err)
			}
			return
		}

		if err = goose.Up(sqlDB, "seed/mysql"); err != nil {
			panic(err)
		}

		return
	}

	brandRepository := brand.NewRepository(sqlDB, snowflakeNode)
	productRepository := product.NewRepository(sqlDB, snowflakeNode)
	transactionRepository := transaction.NewRepository(sqlDB, snowflakeNode)

	adminUseCase := admin.NewUseCase(brandRepository, productRepository)
	customerUseCase := customer.NewUseCase(productRepository, transactionRepository)

	h := httpHandler.NewHandler(adminUseCase, customerUseCase)

	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	serverAddress := fmt.Sprintf("%s:%s", serverHost, serverPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.Handle("/", http.FileServer(http.Dir("handler/http/docs")))
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger.json", serverAddress)),
	))

	mux.HandleFunc("/brand", h.Brand)
	mux.HandleFunc("/product", h.Product)
	mux.HandleFunc("/product/brand", h.ProductBrand)
	mux.HandleFunc("/transaction", h.Transaction)

	server := new(http.Server)
	server.Addr = serverAddress
	server.Handler = mux

	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
