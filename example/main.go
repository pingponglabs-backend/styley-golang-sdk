package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pingponglabs-backend/styley-golang-sdk/sdk"
	"github.com/pingponglabs-backend/styley-golang-sdk/sdk/deployments"
)

func main() {
	client := sdk.NewClient(sdk.WithKey(os.Getenv("X_STYLEY_KEY")))

	deploymentReq := deployments.CreateDeployment{
		Name:    "Background Remover Pro",
		ModelId: "3ca20b58-a07b-492a-8294-e366b5122947",
		Args: map[string]interface{}{
			"image_file": "https://cdn2.styley.ai/d99326d0-45c6-11f0-a674-246e96be0954.jpg",
		},
		OutputFormat: "jpeg",
		OutputWidth:  300,
		OutputHeight: 300,
		Synchronous:  true,
	}
	deplyementResponse, err := client.Deployments().Create(deploymentReq)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Printf("deplyementResponse:: %+v", deplyementResponse)

}

func RunGetJobExample(client *sdk.Client) {
	jobId := "fd2d1136-3da3-421d-83bf-cf5e9c528302"
	job, err := client.Deployments().GetJob(jobId)
	if err != nil {
		panic(err)
	}

	log.Println("job: ", job)
}
