package worktime

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"time"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/store/work-time/models"
	"github.com/google/uuid"
)

type pgWorkTimeStore struct {
	conn *sql.DB
	log  logger.Logger
}

func newPgWorkTimeStore(conn *sql.DB, log logger.Logger) *pgWorkTimeStore {
	return &pgWorkTimeStore{conn, log}
}

func (s *pgWorkTimeStore) Create(ctx context.Context, taskId uuid.UUID, ms int64) (models.WorkTime, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return models.WorkTime{}, err
	}
	defer tx.Rollback()

	workTime := models.WorkTime{
		TaskId:    taskId,
		Ms:        ms,
		Timestamp: time.Now(),
	}

	row := tx.QueryRowContext(ctx, `INSERT INTO public.work_time(
	task_id, milliseconds, "timestamp")
	VALUES ($1, $2, $3) RETURNING id;`, workTime.TaskId, workTime.Ms, workTime.Timestamp)
	if err := row.Err(); err != nil {
		return models.WorkTime{}, err
	}
	if err := row.Scan(&workTime.Id); err != nil {
		return models.WorkTime{}, err
	}
	return workTime, tx.Commit()
}

func (s *pgWorkTimeStore) GetByUser(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) ([]models.TaskWorktime, error) {
	rows, err := s.conn.QueryContext(ctx, `SELECT t.id, SUM(wt.milliseconds) FROM public.work_time wt 
		LEFT JOIN public.task t
		ON wt.task_id = t.id
		WHERE t.assigned_user_id = $1 
		AND wt."timestamp" BETWEEN $2 AND $3
		GROUP BY t.id, wt.task_id
		ORDER BY SUM(wt.milliseconds) desc`,
		userId,
		start,
		end,
	)
	if err != nil {
		return []models.TaskWorktime{}, err
	}
	var workTimes []models.TaskWorktime
	for rows.Next() {
		if err := rows.Err(); err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return []models.TaskWorktime{}, err
			}
		}
		var workTime models.TaskWorktime
		if err := rows.Scan(&workTime.TaskId, &workTime.Ms); err != nil {
			return []models.TaskWorktime{}, err
		}
		workTimes = append(workTimes, workTime)
	}
	return workTimes, nil
}
