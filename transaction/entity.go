package transaction

import (
	"time"

	"github.com/haritsrizkall/api-crowdfunding/campaign"
	"github.com/haritsrizkall/api-crowdfunding/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	Campaign   campaign.Campaign
	UserID     int
	User       user.User
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
