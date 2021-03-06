package server

import (
	"net"

	"github.com/Q-n-A/Q-n-A/util/profiler"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server サーバー
type Server struct {
	e      *echo.Echo
	s      *grpc.Server
	logger *zap.Logger
	c      *Config
}

// Config サーバー用設定
type Config struct {
	DevMode  bool
	GRPCAddr string
	RESTAddr string
}

// NewServer 新しいサーバーを生成
func NewServer(e *echo.Echo, s *grpc.Server, logger *zap.Logger, Config *Config) *Server {
	return &Server{
		e:      e,
		s:      s,
		logger: logger,
		c:      Config,
	}
}

// Run サーバーを起動
func (s *Server) Run() {
	// DevModeがtrueならfgprofサーバーを起動
	if s.c.DevMode {
		go func() {
			err := profiler.StartFgprof(s.logger)
			if err != nil {
				s.logger.Panic("failed to start fgprof server", zap.Error(err))
			}
		}()
	}

	// gRPC用リスナーの作成
	lis, err := net.Listen("tcp", s.c.GRPCAddr)
	if err != nil {
		s.logger.Panic("failed to setup Listener", zap.Error(err))
	}

	// goroutineでgRPCサーバーを起動
	go func() {
		s.logger.Info("Starting gRPC server on " + s.c.GRPCAddr)

		if err := s.s.Serve(lis); err != nil {
			s.logger.Panic("failed to run gRPC server", zap.Error(err))
		}
	}()

	// REST APIサーバーを起動
	s.logger.Info("Starting REST API server on " + s.c.RESTAddr)
	err = s.e.Start(s.c.RESTAddr)
	s.logger.Panic("failed to run REST API server", zap.Error(err))
}
