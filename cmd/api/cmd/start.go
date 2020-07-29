package cmd

import (
	"context"
	"fmt"
	"net/http"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/judwhite/go-svc/svc"
	"github.com/justoxh/go-server/common/wrapper"
	"github.com/justoxh/go-server/config"
	"github.com/justoxh/go-server/log"
	"github.com/justoxh/go-server/router"
	"github.com/spf13/cobra"
)

type Application struct {
	wrapper.Wrapper
	ginEngine  *gin.Engine
	httpServer *http.Server
}

var cfgFile *string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
	server(.exe) start -c config.toml
	start the api`,
	Run: func(cmd *cobra.Command, args []string) {
		app := &Application{}
		if err := svc.Run(app, syscall.SIGINT, syscall.SIGTERM); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	cfgFile = startCmd.Flags().StringP("config", "c", "", "api config file (required)")
	startCmd.MarkFlagRequired("config")
}

func (app *Application) Init(env svc.Environment) error {
	cfg, err := config.Load(cfgFile)
	if err != nil {
		return err
	}
	log.InitLog(&cfg.Logger)
	log.Log.Info(cfg)
	app.ginEngine = router.InitRouter(&cfg)
	return nil
}

func (app *Application) Start() error {
	app.Wrap(func() {
		cfg := config.GetConfig().Server
		app.httpServer = &http.Server{
			Handler:        app.ginEngine,
			Addr:           cfg.ListenAddr,
			ReadTimeout:    cfg.ReadTimeout * time.Second,
			WriteTimeout:   cfg.WriteTimeout * time.Second,
			IdleTimeout:    cfg.IdleTimeout * time.Second,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		}
		log.Log.Info("Listen on :", cfg.ListenAddr)
		if err := app.httpServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	})
	return nil
}

func (app *Application) Stop() error {
	if app.httpServer != nil {
		if err := app.httpServer.Shutdown(context.Background()); err != nil {
			fmt.Printf("http shutdown error:%v\n", err)
		}
		fmt.Println("http shutdown")
	}
	app.Wait()
	return nil
}
