package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ponytail: hand-written mirror of `sqlc generate` (sql_package: pgx/v5).
// Replace by running `sqlc generate` from db/.

type DBTX interface {
	QueryRow(context.Context, string, ...any) pgx.Row
	Query(context.Context, string, ...any) (pgx.Rows, error)
}

var _ DBTX = (*pgxpool.Pool)(nil)

type Queries struct{ db DBTX }

func New(db DBTX) *Queries { return &Queries{db: db} }

// --- auth ---

const getUserByID = `SELECT "id","name","email","emailVerified","image","createdAt","updatedAt" FROM "user" WHERE "id" = $1`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	var u User
	err := q.db.QueryRow(ctx, getUserByID, id).Scan(
		&u.ID, &u.Name, &u.Email, &u.EmailVerified, &u.Image, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

const getSessionByToken = `SELECT "id","userId","token","expiresAt","ipAddress","userAgent","createdAt","updatedAt" FROM "session" WHERE "token" = $1`

func (q *Queries) GetSessionByToken(ctx context.Context, token string) (Session, error) {
	var s Session
	err := q.db.QueryRow(ctx, getSessionByToken, token).Scan(
		&s.ID, &s.UserID, &s.Token, &s.ExpiresAt, &s.IpAddress, &s.UserAgent, &s.CreatedAt, &s.UpdatedAt)
	return s, err
}

// --- health ---

type CreateHealthRecordParams struct {
	UserID      string      `json:"user_id"`
	Kind        string      `json:"kind"`
	TriageLevel pgtype.Text `json:"triage_level"`
	Notes       string      `json:"notes"`
}

const createHealthRecord = `INSERT INTO health_record (user_id, kind, triage_level, notes) VALUES ($1,$2,$3,$4) RETURNING id, user_id, kind, triage_level, notes, created_at`

func (q *Queries) CreateHealthRecord(ctx context.Context, arg CreateHealthRecordParams) (HealthRecord, error) {
	var r HealthRecord
	err := q.db.QueryRow(ctx, createHealthRecord, arg.UserID, arg.Kind, arg.TriageLevel, arg.Notes).Scan(
		&r.ID, &r.UserID, &r.Kind, &r.TriageLevel, &r.Notes, &r.CreatedAt)
	return r, err
}

const listHealthRecordsByUser = `SELECT id, user_id, kind, triage_level, notes, created_at FROM health_record WHERE user_id = $1 ORDER BY created_at DESC`

func (q *Queries) ListHealthRecordsByUser(ctx context.Context, userID string) ([]HealthRecord, error) {
	rows, err := q.db.Query(ctx, listHealthRecordsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []HealthRecord
	for rows.Next() {
		var r HealthRecord
		if err := rows.Scan(&r.ID, &r.UserID, &r.Kind, &r.TriageLevel, &r.Notes, &r.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// --- community ---

type PostMessageParams struct {
	UserID string `json:"user_id"`
	Room   string `json:"room"`
	Body   string `json:"body"`
}

const postMessage = `INSERT INTO community_message (user_id, room, body) VALUES ($1,$2,$3) RETURNING id, user_id, room, body, created_at`

func (q *Queries) PostMessage(ctx context.Context, arg PostMessageParams) (CommunityMessage, error) {
	var m CommunityMessage
	err := q.db.QueryRow(ctx, postMessage, arg.UserID, arg.Room, arg.Body).Scan(
		&m.ID, &m.UserID, &m.Room, &m.Body, &m.CreatedAt)
	return m, err
}

type ListMessagesParams struct {
	Room  string `json:"room"`
	Limit int32  `json:"limit"`
}

const listMessages = `SELECT id, user_id, room, body, created_at FROM community_message WHERE room = $1 ORDER BY created_at DESC LIMIT $2`

func (q *Queries) ListMessages(ctx context.Context, arg ListMessagesParams) ([]CommunityMessage, error) {
	rows, err := q.db.Query(ctx, listMessages, arg.Room, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CommunityMessage
	for rows.Next() {
		var m CommunityMessage
		if err := rows.Scan(&m.ID, &m.UserID, &m.Room, &m.Body, &m.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// --- media ---

type CreateMediaObjectParams struct {
	UserID      string `json:"user_id"`
	Bucket      string `json:"bucket"`
	ObjectKey   string `json:"object_key"`
	ContentType string `json:"content_type"`
}

const createMediaObject = `INSERT INTO media_object (user_id, bucket, object_key, content_type) VALUES ($1,$2,$3,$4) RETURNING id, user_id, bucket, object_key, content_type, created_at`

func (q *Queries) CreateMediaObject(ctx context.Context, arg CreateMediaObjectParams) (MediaObject, error) {
	var m MediaObject
	err := q.db.QueryRow(ctx, createMediaObject, arg.UserID, arg.Bucket, arg.ObjectKey, arg.ContentType).Scan(
		&m.ID, &m.UserID, &m.Bucket, &m.ObjectKey, &m.ContentType, &m.CreatedAt)
	return m, err
}
