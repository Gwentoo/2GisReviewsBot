package period

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func RunPeriodTask(b *tele.Bot, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := CheckNewReviews(b)
			if err != nil {
				log.Printf("Ошибка в периодической задаче: %v", err)
			}
		}
	}

}
