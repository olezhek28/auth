package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	accessV1 "github.com/olezhek28/auth/internal/api/access_v1"
	authV1 "github.com/olezhek28/auth/internal/api/auth_v1"
	noteV1 "github.com/olezhek28/auth/internal/api/note_v1"
	"github.com/olezhek28/auth/internal/client/pg"
	"github.com/olezhek28/auth/internal/closer"
	"github.com/olezhek28/auth/internal/config"
	accessRepository "github.com/olezhek28/auth/internal/repository/access"
	noteRepository "github.com/olezhek28/auth/internal/repository/note"
	userRepository "github.com/olezhek28/auth/internal/repository/user"
	accessService "github.com/olezhek28/auth/internal/service/access"
	authService "github.com/olezhek28/auth/internal/service/auth"
	noteService "github.com/olezhek28/auth/internal/service/note"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	authConfig    config.AuthConfig

	pgClient         pg.Client
	noteRepository   noteRepository.Repository
	userRepository   userRepository.Repository
	accessRepository accessRepository.Repository

	noteService   noteService.Service
	authService   authService.Service
	accessService accessService.Service

	noteImpl   *noteV1.Implementation
	authImpl   *authV1.Implementation
	accessImpl *accessV1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetHTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) GetSwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) GetAuthConfig() config.AuthConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) GetPgClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(s.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		cl, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatalf("failed to get pg client: %s", err.Error())
		}

		err = cl.PG().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.pgClient = cl
	}

	return s.pgClient
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) noteRepository.Repository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.GetPgClient(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetUserRepository(ctx context.Context) userRepository.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.GetPgClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) GetAccessRepository(ctx context.Context) accessRepository.Repository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.GetPgClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) noteService.Service {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}

func (s *serviceProvider) GetAuthService(ctx context.Context) authService.Service {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.GetAuthConfig(),
			s.GetUserRepository(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) GetAccessService(ctx context.Context) accessService.Service {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.GetAuthConfig(),
			s.GetAccessRepository(ctx),
		)
	}

	return s.accessService
}

func (s *serviceProvider) GetNoteImpl(ctx context.Context) *noteV1.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = noteV1.NewImplementation(s.GetNoteService(ctx))
	}

	return s.noteImpl
}

func (s *serviceProvider) GetAuthImpl(ctx context.Context) *authV1.Implementation {
	if s.authImpl == nil {
		s.authImpl = authV1.NewImplementation(s.GetAuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) GetAccessImpl(ctx context.Context) *accessV1.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = accessV1.NewImplementation(s.GetAccessService(ctx))
	}

	return s.accessImpl
}
