package timer

import (
	"context"
	"database/sql"
	"time"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/store/timer/models"
	"github.com/google/uuid"
)

var _ TimerStore = &pgTimerStore{}

type pgTimerStore struct {
	conn *sql.DB
	log  logger.Logger
}

func newPgStore(conn *sql.DB, log logger.Logger) *pgTimerStore {
	return &pgTimerStore{conn, log}
}

func (s *pgTimerStore) Start(ctx context.Context, taskId, userId uuid.UUID) (uuid.UUID, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()
	startTimestamp := time.Now()
	row := tx.QueryRowContext(ctx, `INSERT INTO public.work_timer (task_id, user_id, timestamp_start) VALUES
	($1, $2, $3) RETURNING id`, taskId, userId, startTimestamp)
	if err := row.Err(); err != nil {
		return uuid.Nil, err
	}
	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, tx.Commit()
}

func (s *pgTimerStore) Stop(ctx context.Context, timerId uuid.UUID) (models.TimerResult, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return models.TimerResult{}, err
	}
	defer tx.Rollback()
	row := tx.QueryRowContext(ctx, `DELETE FROM public.work_timer WHERE id = $1 RETURNING id, task_id, timestamp_start`, timerId)
	if err := row.Err(); err != nil {
		return models.TimerResult{}, err
	}
	var workTime models.TimerResult
	var timerStart time.Time
	if err := row.Scan(&workTime.Id, &workTime.TaskId, &timerStart); err != nil {
		return models.TimerResult{}, err
	}
	return workTime, tx.Commit()
}
