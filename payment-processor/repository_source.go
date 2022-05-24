package payment

import (
	"context"
	"database/sql"
	"fmt"
)

// ISourceRepository is a interface for Source repository.
type ISourceRepository interface {
	GetByID(ctx context.Context, sourceID string) (Source, error)
}

// SourcesRepository is a repository to communicate with sources db.
type SourcesRepository struct {
	db IQuerier
}

// NewSourcesRepository SourcesRepository's constructor.
func NewSourcesRepository(db IQuerier) *SourcesRepository {
	return &SourcesRepository{
		db: db,
	}
}

// GetByID returns a Source.
func (s *SourcesRepository) GetByID(ctx context.Context, sourceID string) (Source, error) {
	var src Source
	err := s.db.QueryRowContext(ctx, "SELECT source_id, card_number, cvv FROM sources WHERE source_id = ?", sourceID).Scan(
		&src.SourceID, &src.CardNumber, &src.CVV,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return src, NewInvalidContentError("source_id not found")
		}
		return src, NewInternalServerError(fmt.Sprintf("failed execute query: %s", err.Error()))
	}

	return src, nil
}
