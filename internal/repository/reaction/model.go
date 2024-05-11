package reaction

import "time"

type reaction struct {
	UserID             string
	ReactionType       string
	ReactionTargetID   string
	ReactionTargetType string
	Created            time.Time `xorm:"created"`
	Updated            time.Time `xorm:"updated"`
}
