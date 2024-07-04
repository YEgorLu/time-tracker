package profile

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/YEgorLu/time-tracker/internal/store/profile/models"
)

var _ ProfileStore = &pgProfileStore{}

type pgProfileStore struct {
	conn *sql.DB
}

func newPgProfileStore(conn *sql.DB) *pgProfileStore {
	return &pgProfileStore{
		conn: conn,
	}
}

// Create implements ProfileStore.
func (p *pgProfileStore) Create(ctx context.Context, profile models.Profile) (models.Profile, error) {
	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return models.Profile{}, err
	}
	defer tx.Rollback()
	row := tx.QueryRowContext(ctx, `INSERT INTO public.profile(
	pass_serie, pass_number, name, surname, patronymic, address)
	VALUES ($1, $2, $3, $4, $5, $6)
	returning pass_serie, pass_number, name, surname, patronymic, address;`,
		profile.PassportSerie, profile.PassportNumber, profile.Name, profile.Surname, profile.Patronymic, profile.Address)
	if err = row.Err(); err != nil {
		return models.Profile{}, err
	}
	var createdProfile models.Profile
	if err := row.Scan(
		&createdProfile.PassportSerie,
		&createdProfile.PassportNumber,
		&createdProfile.Name,
		&createdProfile.Surname,
		&createdProfile.Patronymic,
		&createdProfile.Address,
	); err != nil {
		return models.Profile{}, err
	}
	return createdProfile, tx.Commit()
}

// Delete implements ProfileStore.
func (p *pgProfileStore) Delete(ctx context.Context, passportSerie, passportNumber string) error {
	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `DELETE FROM public.profile
	WHERE pass_serie = $1 AND pass_number = $2;`, passportSerie, passportNumber)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// GetMany implements ProfileStore.
func (p *pgProfileStore) GetMany(ctx context.Context, page, size int, filter models.ProfileFilter) ([]models.Profile, error) {
	query := `SELECT pass_serie, pass_number, name, surname, patronymic, address
	FROM public.profile
	WHERE `
	var searchParams []string
	args := make([]any, 0, len(filter.Name)+len(filter.Surname)+len(filter.Patronymic)+len(filter.Address))
	parameterTypes := [...]struct {
		argName string
		values  []string
	}{
		{"name", filter.Name},
		{"surname", filter.Surname},
		{"patronymic", filter.Patronymic},
		{"address", filter.Address},
	}

	paramIndex := 0
	for _, parameterType := range parameterTypes {
		if len(parameterType.values) > 0 {
			maskSl := make([]string, 0, len(parameterType.values))
			for _, v := range parameterType.values {
				paramIndex++
				maskSl = append(maskSl, "$"+strconv.Itoa(paramIndex))
				args = append(args, v)
			}
			searchParams = append(searchParams, parameterType.argName+` IN (`+strings.Join(maskSl, ", ")+`) `)
		}
	}
	query += strings.Join(searchParams, " AND ")
	rows, err := p.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return []models.Profile{}, err
	}
	defer rows.Close()
	var profiles []models.Profile
	for rows.Next() {
		err := rows.Err()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return []models.Profile{}, err
			}
		}

		var profile models.Profile
		if err = rows.Scan(
			&profile.PassportSerie,
			&profile.PassportNumber,
			&profile.Name,
			&profile.Surname,
			&profile.Patronymic,
			&profile.Address,
		); err != nil {
			return []models.Profile{}, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// GetOne implements ProfileStore.
func (p *pgProfileStore) GetOne(ctx context.Context, passportSerie, passportNumber string) (models.Profile, error) {
	row := p.conn.QueryRowContext(ctx, `SELECT pass_serie, pass_number, name, surname, patronymic, address
	FROM public.profile;
	WHERE pass_serie = $1 AND pass_number = $2`, passportSerie, passportNumber)
	if err := row.Err(); err != nil {
		return models.Profile{}, err
	}
	var profile models.Profile
	err := row.Scan(
		&profile.PassportSerie,
		&profile.PassportNumber,
		&profile.Name,
		&profile.Surname,
		&profile.Patronymic,
		&profile.Address,
	)
	return profile, err
}

// Update implements ProfileStore.
func (p *pgProfileStore) Update(ctx context.Context, profile models.Profile) error {
	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tx.ExecContext(ctx, `UPDATE public.profile
	SET name=$3, surname=$4, patronymic=$5, address=$6
	WHERE pass_serie=$1 AND pass_number=$2;`,
		profile.PassportSerie,
		profile.PassportNumber,
		profile.Name,
		profile.Surname,
		profile.Patronymic,
		profile.Address,
	)
	return tx.Commit()
}
