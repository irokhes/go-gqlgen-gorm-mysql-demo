package main

import (
	"fmt"
	"go-gqlgen-gorm-mysql-demo/graph"
	"go-gqlgen-gorm-mysql-demo/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/soberkoder/go-orders-graphql-api/models"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:password@tcp(localhost:3306)/db?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&models.Order{}, &models.Item{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	initDB()
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
