package note

import (
	"context"
	"time"

	"github.com/olezhek28/auth/internal/model"
	"github.com/olezhek28/auth/internal/sys/validate"
)

func (s *service) Create(ctx context.Context, info *model.Info) (int64, error) {
	//err := validate.Validate(
	//	ctx,
	//	validateCreatedAt(info.CreatedAt),
	//	validateTitle(info.Title),
	//)
	//if err != nil {
	//	return 0, err
	//}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	time.Sleep(6 * time.Second)

	id, err := s.noteRepository.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func validateCreatedAt(createdAt time.Time) validate.Condition {
	return func(ctx context.Context) error {
		if !createdAt.Equal(time.Now()) {
			return validate.NewValidationErrors("createdAt is not equal to current time")
		}

		return nil
	}
}

func validateTitle(title string) validate.Condition {
	return func(ctx context.Context) error {
		if len(title) < 3 {
			return validate.NewValidationErrors("title is too short")
		}

		return nil
	}
}
