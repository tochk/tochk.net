package datastruct

import "time"

type Projects struct {
	ID               int       `db:"id"`
	Title            string    `db:"title"`
	ShortDescription string    `db:"short_description"`
	Description      string    `db:"description"`
	ImageURL         string    `db:"image_url"`
	Language         string    `db:"language"`
	RedirectURL      string    `db:"redirect_url"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	TeamMembersIDs   []int     `db:"-"`
	TagsIDs          []int     `db:"-"`
}

type ProjectsTags struct {
	ProjectID int `db:"project_id"`
	TagID     int `db:"tag_id"`
}

type ProjectsTeamMembers struct {
	ProjectID    int `db:"project_id"`
	TeamMemberID int `db:"team_member_id"`
}
