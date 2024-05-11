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
	albumhandler "github.com/zuzuka28/music_land_api/internal/handler/rest/album"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/middleware/auth"
	reactionhandler "github.com/zuzuka28/music_land_api/internal/handler/rest/reaction"
	trackhandler "github.com/zuzuka28/music_land_api/internal/handler/rest/track"
	userhandler "github.com/zuzuka28/music_land_api/internal/handler/rest/user"
	albumrepo "github.com/zuzuka28/music_land_api/internal/repository/album"
	reactionrepo "github.com/zuzuka28/music_land_api/internal/repository/reaction"
	trackrepo "github.com/zuzuka28/music_land_api/internal/repository/track"
	userrepo "github.com/zuzuka28/music_land_api/internal/repository/user"
	albumsrv "github.com/zuzuka28/music_land_api/internal/service/album"
	authsrv "github.com/zuzuka28/music_land_api/internal/service/auth"
	reactionsrv "github.com/zuzuka28/music_land_api/internal/service/reaction"
	tracksrv "github.com/zuzuka28/music_land_api/internal/service/track"
	usersrv "github.com/zuzuka28/music_land_api/internal/service/user"
	"github.com/zuzuka28/music_land_api/pkg/logging"
	"github.com/zuzuka28/music_land_api/pkg/minio"
	"github.com/zuzuka28/music_land_api/pkg/tracing"
	"xorm.io/xorm"
)

//revive:disable-next-line:cyclomatic,function-length
func runServer(c *cli.Context) error {
	cfg, err := config.NewAPI(c.String("config"))
	if err != nil {
		return fmt.Errorf("new config: %w", err)
	}

	l := logging.NewLogger(logging.LogLevel(cfg.LogLevel))

	t := tracing.NewNoop()

	if cfg.Tracing != nil {
		exp, er := tracing.NewOTLPExporter(c.Context, cfg.Tracing.Exporter)
		if er != nil {
			return fmt.Errorf("new otlp exporter: %w", er)
		}

		t = tracing.New(exp, tracing.NewResource(cfg.Tracing.Resource))
	}

	handlerT := t.Child("handler")
	serviceT := t.Child("service")
	repoT := t.Child("repository")
	fsT := t.Child("filestorage")

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.User, cfg.Storage.Password, cfg.Storage.Name)

	eng, err := xorm.NewEngine("postgres", dbinfo)
	if err != nil {
		return fmt.Errorf("create xorm engine: %w", err)
	}

	fscli, err := minio.NewClient(cfg.FileStorage, fsT.Child("minio"))
	if err != nil {
		return fmt.Errorf("create file storage client: %w", err)
	}

	trepo, err := trackrepo.NewRepository(eng, repoT.Child("track"))
	if err != nil {
		return fmt.Errorf("create track repository: %w", err)
	}

	urepo, err := userrepo.NewRepository(eng, repoT.Child("user"))
	if err != nil {
		return fmt.Errorf("create user repository: %w", err)
	}

	arepo, err := albumrepo.NewRepository(eng, repoT.Child("album"))
	if err != nil {
		return fmt.Errorf("create album repository: %w", err)
	}

	rrepo, err := reactionrepo.NewRepository(eng, repoT.Child("reaction"))
	if err != nil {
		return fmt.Errorf("create reaction repository: %w", err)
	}

	usrv := usersrv.NewService(urepo, serviceT.Child("user"))
	tsrv := tracksrv.NewService(trepo, fscli, serviceT.Child("track"))
	albsrv := albumsrv.NewService(arepo, serviceT.Child("album"))
	rsrv := reactionsrv.NewService(rrepo, serviceT.Child("reaction"))
	asrv := authsrv.NewService(usrv, serviceT.Child("auth"))

	uhandler := userhandler.NewHandler(usrv, handlerT.Child("user"))
	thandler := trackhandler.NewHandler(tsrv, handlerT.Child("track"))
	ahandler := albumhandler.NewHandler(albsrv, handlerT.Child("album"))
	rhandler := reactionhandler.NewHandler(rsrv, handlerT.Child("reaction"))
	amw := auth.BasicMiddleware(asrv)

	api := rest.NewHandler(
		uhandler,
		thandler,
		ahandler,
		rhandler,
		amw,
		l,
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
		l.Info("graceful shutdown initiated", sig)

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
