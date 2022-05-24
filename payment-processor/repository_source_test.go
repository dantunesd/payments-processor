package payment

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
)

func TestSourcesRepository_GetByID(t *testing.T) {
	type fields struct {
		db IQuerier
	}
	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Source
		wantErr bool
	}{
		{
			"get with success",
			fields{
				QuerierMock{
					queryRowContext: func(ctx context.Context, query string, args ...interface{}) IScanner {
						return &ScannerMock{
							scan: func(dest ...interface{}) error {
								return nil
							},
						}
					},
				},
			},
			args{},
			Source{},
			false,
		},
		{
			"no rows in result set",
			fields{
				QuerierMock{
					queryRowContext: func(ctx context.Context, query string, args ...interface{}) IScanner {
						return &ScannerMock{
							scan: func(dest ...interface{}) error {
								return sql.ErrNoRows
							},
						}
					},
				},
			},
			args{},
			Source{},
			true,
		},
		{
			"fail on query",
			fields{
				QuerierMock{
					queryRowContext: func(ctx context.Context, query string, args ...interface{}) IScanner {
						return &ScannerMock{
							scan: func(dest ...interface{}) error {
								return errors.New("invalid syntax")
							},
						}
					},
				},
			},
			args{},
			Source{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSourcesRepository(tt.fields.db)

			got, err := s.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SourcesRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SourcesRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

type QuerierMock struct {
	queryRowContext func(ctx context.Context, query string, args ...interface{}) IScanner
}

func (q QuerierMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) IScanner {
	return q.queryRowContext(ctx, query, args)
}

type ScannerMock struct {
	scan func(dest ...interface{}) error
}

func (s *ScannerMock) Scan(dest ...interface{}) error {
	return s.scan(dest)
}
