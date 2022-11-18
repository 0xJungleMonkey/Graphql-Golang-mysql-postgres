package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/OxJungleMonkey/go-graphql/graph"
	"github.com/OxJungleMonkey/go-graphql/graph/generated"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/zenyui/gqlgen-dataloader/graph/dataloader"

	"github.com/zenyui/gqlgen-dataloader/graph/storage"
)

var mydb *gorm.DB
var psdb *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	mydb, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	mydb.LogMode(true)
	mydb.Exec("USE rocket_development")
	// var FactIntervention gorm.Model
	// dsn := os.Getenv("dsn")
	dsn := "host=localhost user=postgres password=password port=5432 sslmode=disable dbname=myapp_development"
	psdb, err = gorm.Open("postgres", dsn)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	psdb.LogMode(true)

}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB()
	db := storage.NewMemoryStorage()
	loader := dataloader.NewDataLoader(db)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  &graph.Resolver{DB1: mydb, DB2: psdb},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))
	dataloaderSrv := dataloader.Middleware(loader, srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	http.Handle("/query", dataloaderSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
