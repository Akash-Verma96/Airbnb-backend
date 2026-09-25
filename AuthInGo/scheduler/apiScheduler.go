package scheduler

import (
	"fmt"
	"os/exec"

	"github.com/robfig/cron/v3"
)
func ApiScheduler(){
	fmt.Println("Scheduling Cron Job...")
	cronJob := cron.New()

	_, err := cronJob.AddFunc("@every 14m", func() {
		fmt.Println("CronJob Every 14 minute second!")

		authPing := exec.Command("curl", "-X", "GET", "https://airbnb-auth.onrender.com/ping")

		if err := authPing.Run(); err != nil {
			fmt.Println("Auth ping api call failed..",err)
		}

		hotelPing := exec.Command("curl", "-X", "GET", "https://airbnb-hotel-0job.onrender.com/api/v1/ping")

		if err := hotelPing.Run(); err != nil {
			fmt.Println("Hotel ping api call failed..",err)
		}

		bookingPing := exec.Command("curl", "-X", "GET", "https://airbnb-backend-glo7.onrender.com/api/v1/ping")

		if err := bookingPing.Run(); err != nil {
			fmt.Println("Booking ping api call failed..",err)
		}
	})

	if err != nil {
		fmt.Println("Error adding api to cronJob", err)
	}

	fmt.Println("Starting Cron Job...")
	cronJob.Start()
}