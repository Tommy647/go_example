package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Tommy647/go_example/internal/httpserver"
	"github.com/Tommy647/go_example/internal/middleware"
	_ "github.com/lib/pq"
)

const (
	// shutdownWait duration when attempting a graceful shutdown
	shutdownWait = 5 * time.Second

	// environment variable names
	envGreeter = `GREETER`     // which greeter to use
	dbHost     = `DB_HOST`     // database host
	dbPort     = `DB_PORT`     // database port
	dbUser     = `DB_USER`     // database user
	dbPassword = `DB_PASSWORD` // database password
	dbDbname   = `DB_DBNAME`   // database name
)

// set up a simple webserver
func main() {
	// monitor system calls to detect a shut-down (SYSTERM||SYSINT)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	// create a context to control the application main routine
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Println("waiting for os syscall")
		// block waiting for a signal on c (os.Syscall)
		oscall := <-c
		// log just for observability
		log.Printf("system call: %s", oscall.String())
		// cancel the context created above, which will cascade to other routines using this context
		cancel()
	}()

	// Open another DB connection
	dbConn, err := sql.Open("postgres", getPostgresConnection())
	if err != nil {
		log.Println("error opening the DB", err.Error())
	}
	defer dbConn.Close()
	// start the http server
	if err := serve(ctx, dbConn); err != nil {
		log.Println(err.Error())
	}
}

func serve(ctx context.Context, db *sql.DB) error {
	mux := http.NewServeMux()
	// attach the handler - this pattern works well for simple apps
	mux.Handle(
		"/hello",
		middleware.WithDefault(httpserver.HandleHello(), true),
	)

	mux.Handle("/coffee",
		middleware.WithDefault(httpserver.HandleCoffee(db), true),
	)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// ListenAndServe blocks until the service stops, so we run it in a go routine
	go func() {
		// ListenServe runs until closed
		log.Fatal(srv.ListenAndServe())
	}()
	log.Println("http server started")
	// block here until context get cancelled
	<-ctx.Done()
	log.Println("serve context cancelled: shutting down")
	// create a new context for the shutdown action - we want to time box this to just 5 seconds
	ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownWait)
	// close the context off when we leave this function - best practice: we are done, so clean up
	defer cancel()
	// attempt to gracefully shut down the http server
	if err := srv.Shutdown(ctxShutdown); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// getPostgresConnection string we need to open the connection
// gets connection details from the environment for now @todo: replace with viper
func getPostgresConnection() string {
	host := os.Getenv(dbHost)
	port, err := strconv.Atoi(os.Getenv(dbPort))
	if err != nil {
		panic("port must be a number " + err.Error())
	}
	user := os.Getenv(dbUser)
	password := os.Getenv(dbPassword)
	dbname := os.Getenv(dbDbname)

	return fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname,
	)
}
