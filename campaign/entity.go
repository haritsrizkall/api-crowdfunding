package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BakerCount       int
	Slug             string
	CreatedAt        time.Time
	updatedAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	updatedAt  time.Time
}
