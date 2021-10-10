package datastruct

import "time"

type Projects struct {
	ID               int       `db:"id"`
	Title            string    `db:"title"`
	ShortDescription string    `db:"short_description"`
	Description      string    `db:"description"`
	ImageURL         string    `db:"image_url"`
	Language         string    `db:"language"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
