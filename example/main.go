package main

import (
	"log"
	"os"

	"github.com/pingponglabs-backend/sdks/pingpong-golang-sdk/sdk"
	"github.com/pingponglabs-backend/sdks/pingpong-golang-sdk/sdk/deployments"
)

func main() {
	client := sdk.NewClient(sdk.WithKey(os.Getenv("X_STYLEY_KEY")))

	deployment, err := client.Deployments().Create(deployments.CreateDeployment{
		Name: "Background Removal",
		Args: map[string]interface{}{
			"input": "https://cdn.mediamagic.dev/media/c7dbd266-3aa3-11ed-8e27-e679ed67c206.jpeg",
		},
		Model: "844218fa-c5d0-4cee-90ce-0b42d226ac8d",
		Sync:  true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("deployment: %v", deployment)
}

func RunGetJobExample(client *sdk.Client) {
	jobId := "fd2d1136-3da3-421d-83bf-cf5e9c528302"
	job, err := client.Deployments().GetJob(jobId)
	if err != nil {
		panic(err)
	}

	log.Println("job: ", job)
}
