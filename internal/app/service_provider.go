package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	noteV1 "github.com/olezhek28/auth/internal/api/note_v1"
	"github.com/olezhek28/auth/internal/client/pg"
	"github.com/olezhek28/auth/internal/closer"
	"github.com/olezhek28/auth/internal/config"
	noteRepository "github.com/olezhek28/auth/internal/repository/note"
	noteService "github.com/olezhek28/auth/internal/service/note"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgClient       pg.Client
	noteRepository noteRepository.Repository
	noteService    noteService.Service

	noteImpl *noteV1.Implementation
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

func (s *serviceProvider) GetNoteService(ctx context.Context) noteService.Service {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}

func (s *serviceProvider) GetNoteImpl(ctx context.Context) *noteV1.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = noteV1.NewImplementation(s.GetNoteService(ctx))
	}

	return s.noteImpl
}
