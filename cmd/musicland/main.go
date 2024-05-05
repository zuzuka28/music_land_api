package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"github.com/zuzuka28/music_land_api/internal/config"
	"github.com/zuzuka28/music_land_api/internal/handler/rest"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/middleware/auth"
	trackhandler "github.com/zuzuka28/music_land_api/internal/handler/rest/track"
	userhandler "github.com/zuzuka28/music_land_api/internal/handler/rest/user"
	trackrepo "github.com/zuzuka28/music_land_api/internal/repository/track"
	userrepo "github.com/zuzuka28/music_land_api/internal/repository/user"
	authsrv "github.com/zuzuka28/music_land_api/internal/service/auth"
	tracksrv "github.com/zuzuka28/music_land_api/internal/service/track"
	usersrv "github.com/zuzuka28/music_land_api/internal/service/user"
	"github.com/zuzuka28/music_land_api/pkg/logging"
	"github.com/zuzuka28/music_land_api/pkg/minio"
	"xorm.io/xorm"
)

//revive:disable-next-line:cyclomatic,function-length
func runServer(c *cli.Context) error {
	cfg, err := config.NewAPI(c.String("config"))
	if err != nil {
		return fmt.Errorf("new config: %w", err)
	}

	lg := logging.NewLogger(logging.LogLevel(cfg.LogLevel))

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.User, cfg.Storage.Password, cfg.Storage.Name)

	eng, err := xorm.NewEngine("postgres", dbinfo)
	if err != nil {
		return fmt.Errorf("create xorm engine: %w", err)
	}

	fscli, err := minio.NewClient(cfg.FileStorage)
	if err != nil {
		return fmt.Errorf("create file storage client: %w", err)
	}

	trepo, err := trackrepo.NewRepository(eng)
	if err != nil {
		return fmt.Errorf("create track repository: %w", err)
	}

	urepo, err := userrepo.NewRepository(eng)
	if err != nil {
		return fmt.Errorf("create user repository: %w", err)
	}

	usrv := usersrv.NewService(urepo)
	tsrv := tracksrv.NewService(trepo, fscli)
	asrv := authsrv.NewService(usrv)

	uhandler := userhandler.NewHandler(usrv)
	thandler := trackhandler.NewHandler(tsrv)
	amw := auth.BasicMiddleware(asrv)

	api := rest.NewHandler(
		uhandler,
		thandler,
		amw,
	)

	addr := cfg.Service.Host + ":" + strconv.Itoa(cfg.Service.Port)

	errCh := make(chan error)

	go func() {
		if err := api.Run(addr); err != nil {
			errCh <- fmt.Errorf("run webserver: %w", err)
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return err

	case sig := <-osSignals:
		lg.Info("graceful shutdown initiated", sig)

		// TODO: Add graceful shutdown logic here

		return nil
	}
}

func main() {
	app := &cli.App{ //nolint:exhaustruct
		Name:  "music land API",
		Usage: "music land API",
		Flags: []cli.Flag{
			&cli.StringFlag{ //nolint:exhaustruct
				Name:  "config",
				Value: "./config.yml",
				Usage: "path to the config file",
			},
		},
		Action:   runServer,
		Commands: []*cli.Command{},
	}

	if err := app.Run(os.Args); err != nil {
		logging.NewLogger(logging.LogLevelError).Error("can't run application: " + err.Error())
	}
}
