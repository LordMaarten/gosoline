package dispatcher

import (
	"context"
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/db-repo"
	"github.com/applike/gosoline/pkg/mon"
)

type Repository struct {
	db_repo.Repository
	dispatcher Dispatcher
	logger     mon.Logger
}

func NewRepository(config cfg.Config, logger mon.Logger, repo db_repo.Repository) db_repo.Repository {
	disp := Get()

	return &Repository{
		Repository: repo,
		dispatcher: disp,
		logger:     logger,
	}
}

func (r Repository) Create(ctx context.Context, value db_repo.ModelBased) error {
	err := r.Repository.Create(ctx, value)

	if err != nil {
		return err
	}

	eventName := fmt.Sprintf("%s.%s", r.Repository.GetModelName(), db_repo.Create)
	errs := r.dispatcher.Fire(ctx, eventName, value)

	for _, err := range errs {
		if err != nil {
			r.logger.Error(err, "error on "+db_repo.Create+" for event "+eventName)
		}
	}

	return nil
}

func (r Repository) Update(ctx context.Context, value db_repo.ModelBased) error {
	err := r.Repository.Update(ctx, value)

	if err != nil {
		return err
	}

	eventName := fmt.Sprintf("%s.%s", r.Repository.GetModelName(), db_repo.Update)
	errs := r.dispatcher.Fire(ctx, eventName, value)

	for _, err := range errs {
		if err != nil {
			r.logger.Error(err, "error on "+db_repo.Update+" for event "+eventName)
		}
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, value db_repo.ModelBased) error {
	err := r.Repository.Delete(ctx, value)

	if err != nil {
		return err
	}

	eventName := fmt.Sprintf("%s.%s", r.Repository.GetModelName(), db_repo.Delete)
	errs := r.dispatcher.Fire(ctx, eventName, value)

	for _, err := range errs {
		if err != nil {
			r.logger.Error(err, "error on "+db_repo.Delete+" for event "+eventName)
		}
	}

	return nil
}
