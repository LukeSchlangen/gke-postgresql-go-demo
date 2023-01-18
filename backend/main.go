// main.go
package main

import (
	"log"
	"net/http"
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"

)

func main() {
    http.HandleFunc("/create-table", func(w http.ResponseWriter, r *http.Request) {
		// Create database connection
		dsn := fmt.Sprintf("user=%s database=%s", "gke-quickstart-service-account@gke-postgresql-0018.iam", "quickstart_db")
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Error parsing configuration with pgx.ParseConfig(dsn): %v", err)
		}
	
		d, err := cloudsqlconn.NewDialer(context.Background(), cloudsqlconn.WithIAMAuthN())
		if err != nil {
			log.Fatalf("Error connecting to database with sql.Open: %v", err)
		}
		

		config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
			connectionString := "gke-postgresql-0018:us-central1:quickstart-instanc"
			log.Printf("Intentionally using incorrect connection string %s", connectionString)
			return d.Dial(ctx, connectionString)
		}
		dbURI := stdlib.RegisterConnConfig(config)
		dbPool, err := sql.Open("pgx", dbURI)
		if err != nil {
			log.Fatalf("Error connecting to database with sql.Open: %v", err)
		}
		// Create table
		createItems := `CREATE TABLE items (
			id SERIAL NOT NULL,
			title VARCHAR(1000) NOT NULL,
			PRIMARY KEY (id)
		);`
		_, err = dbPool.Exec(createItems)
		if err != nil {
			log.Fatalf("Error creating table with dbPool.Exec: %v", err)
		}
        w.Write([]byte("Successfully created the table!"))
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello Go on GKE!"))
    })
	log.Print("Running main func")
    http.ListenAndServe(":8080", nil)
}
