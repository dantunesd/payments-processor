package payment

import (
	"context"
	"database/sql"
)

// ISourceRepository is a interface for Source repository
type ISourceRepository interface {
	GetByID(ctx context.Context, ID string) (Source, error)
}

// SourcesRepository is a repository for Sources db
type SourcesRepository struct {
	db IQuerier
}

// NewSourcesRepository constructor for SourcesRepository
func NewSourcesRepository(db IQuerier) *SourcesRepository {
	return &SourcesRepository{
		db: db,
	}
}

// GetByID returns a Source
func (s *SourcesRepository) GetByID(ctx context.Context, ID string) (Source, error) {
	var src Source
	err := s.db.QueryRowContext(ctx, "SELECT source_id, card_number, cvv FROM sources WHERE source_id = ?", ID).Scan(
		&src.SourceID, &src.CardNumber, &src.CVV,
	)
	if err == sql.ErrNoRows {
		return src, NewInvalidContentError("Invalid source_id")
	}
	return src, NewInternalServerError(err.Error())
}
