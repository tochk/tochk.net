package datastruct

import "time"

type Articles struct {
	ID             int       `db:"id"`
	Title          string    `db:"title"`
	ShortText      string    `db:"short_text"`
	Text           string    `db:"text"`
	Language       string    `db:"language"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	TeamMembersIDs []int     `db:"-"`
	TagsIDs        []int     `db:"-"`
}

type ArticlesTags struct {
	ArticleID int `db:"article_id"`
	TagID     int `db:"tag_id"`
}

type ArticlesTeamMembers struct {
	ArticleID    int `db:"article_id"`
	TeamMemberID int `db:"team_member_id"`
}
