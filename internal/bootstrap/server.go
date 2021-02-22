package bootstrap

import (
	"context"
	"fmt"
	"github.com/feildrixliemdra/go-restful-api/internal/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewHTTPServer() {

	//err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	//	pathTemplate, err := route.GetPathTemplate()
	//	if err == nil {
	//		fmt.Println("ROUTE:", pathTemplate)
	//	}
	//	pathRegexp, err := route.GetPathRegexp()
	//	if err == nil {
	//		fmt.Println("Path regexp:", pathRegexp)
	//	}
	//	queriesTemplates, err := route.GetQueriesTemplates()
	//	if err == nil {
	//		fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
	//	}
	//	queriesRegexps, err := route.GetQueriesRegexp()
	//	if err == nil {
	//		fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
	//	}
	//	methods, err := route.GetMethods()
	//	if err == nil {
	//		fmt.Println("Methods:", strings.Join(methods, ","))
	//	}
	//	fmt.Println()
	//	return nil
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	r := router.NewRouter()

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", viper.GetInt("app.port")),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}

	}()

	log.Infof("server up and running on port %v", viper.GetInt("app.port"))

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("server shutting down successfully")
	os.Exit(0)
}
