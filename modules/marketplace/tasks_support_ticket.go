package marketplace

import (
	"github.com/jasonlvhit/gocron"
)

func StartSupportTicketCron() {
	gocron.Every(15).Minutes().Do(RefreshSerpItemsMaterializedView)
}
