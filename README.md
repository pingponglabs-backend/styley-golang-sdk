# **Styley - Golang SDK**

The Styley SDK for Go is a developer-friendly Software Development Kit (SDK) that simplifies interaction with Styley API services directly from your Go applications. It abstracts complexities like HTTP requests, authentication, and response parsing, letting you focus on business logic.

---

### **Key Features**
- **Easy to Use**: Simple method calls, no custom HTTP requests needed.
- **Built-in Authentication**: Set your API key once and forget about it.
- **Efficient Development**: Reduce boilerplate code for API interactions.
- **Integration Ready**: Works seamlessly with Go frameworks like `gin` and `echo`.

---

## **Contents**
- [Installation Guide](#installation-guide)
- [Examples](#examples)
- [Summary of Available Methods](#summary-of-available-methods)

---

## **Installation Guide**

### **📥 1. Install Golang**
Follow these steps to install Go:

1. Visit the [official Golang download page](https://go.dev/doc/install).
2. Download the appropriate installer for your operating system:
   - **Windows**: `.msi` file
   - **MacOS**: `.pkg` file
   - **Linux**: `.tar.gz` file
3. Run the installer and follow the instructions.

---

### **⚙️ 2. Verify Installation**
Check if Go is installed correctly:

```bash
go version
```

Expected output:

```bash
go version go1.21.1 linux/amd64
```

If you see "command not found," ensure the Go binary is in your `PATH`.

---

### **📁 3. Setup Go Workspace**
Organize your projects using a Go workspace:

1. Create a workspace directory:

    ```bash
    mkdir -p $HOME/go/src/github.com/styley/styleygolang-sdk
    cd $HOME/go/src/github.com/styley/styleygolang-sdk
    ```

2. Initialize the Go module:

    ```bash
    go mod init github.com/styley/styleygolang-sdk
    ```

---

### **📘 SDK Installation**
Install the Styley SDK using:

```bash
go get git.mediamagic.ai/styley-sdks/styley-golang-sdk@latest

go get git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk/deployments

or 

# After import use 
go mod tidy

```
This adds the SDK to your `go.mod` file.

---

### **🔑 Environment Variables**
Set the following environment variable to authenticate API requests:

```bash
export X_STYLEY_KEY=***************************
```

Alternatively, create a `.env` file with:

```bash
X_STYLEY_KEY=***************************
```

---

## **Examples**

### 📤 **Create Deployments**

 The **Create Deployment** method allows you to create a new deployment using a model name, arguments, and deployment name.The method returns an output with an job_id which can be used to fetch the final results.


```go
package main

import (
	"fmt"
	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk"
	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk/deployments"
)

func main() {
	client := sdk.NewClient()

	deployment, err := client.Deployments().Create(deployments.CreateDeployment{
		Name: "Gif Maker",
		Args: map[string]interface{}{
			"height":          "384",
			"mp4":             "true",
			"negative_prompt": "blurry",
			"prompt":          "A dog with a stick",
			"scheduler":       "EulerAncestralDiscreteScheduler",
			"seed":            "1",
			"steps":           "30",
			"width":           "672",
		},
		Model: "66a8ccd5-ed5d-4133-b0c8-c3862a575a9e",
	})

	if err != nil {
		fmt.Println("Error creating deployment:", err)
		return
	}

	fmt.Printf("Deployment created successfully: %+v \n", deployment)
}

```

---

### 📄 **Get Job**
Fetch details of a job by its `JobID`:

```go
package main

import (
	"fmt"

	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk"
)

func main() {
	client := sdk.NewClient()

// <job_id> - replace with actual jobID from Deployement response
	job, err := client.Deployments().GetJob("<job_id>")

	if err != nil {
		fmt.Println("Error fetching job:", err)
		return
	}

	fmt.Printf("Job details: %+v\n", job)
}

```

---

### 📜 **List Deployments**
List all deployments in your account:

```go
package main

import (
	"fmt"

	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk"
)

func main() {
	client := sdk.NewClient()

	deployments, err := client.Deployments().List()
	if err != nil {
		fmt.Println("Error listing deployments:", err)
		return
	}

	fmt.Printf("Deployments: %+v\n", deployments)
}

```

---

### 🔍 **Get Model by ID**
Retrieve model details using its `ModelID`:

```go
package main

import (
	"fmt"

	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk"
)

func main() {
	client := sdk.NewClient()
	model, err := client.Models().GetByID("66a8ccd5-ed5d-4133-b0c8-c3862a575a9e")

	if err != nil {
		fmt.Println("Error fetching model by ID:", err)
		return
	}

	fmt.Printf("Model details: %+v\n", model)
}

```

---

### 🔍 **Get Model by Name**
Retrieve model details using its `ModelName`:

```go
package main

import (
	"fmt"

	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk"
)

func main() {
	client := sdk.NewClient()

	model, err := client.Models().GetByName("Gif Maker")

	if err != nil {
		fmt.Println("Error fetching model by name:", err)
		return
	}

	fmt.Printf("Model details: %+v\n", model)
}

```

---

### 📜 **List Models**
List all available models:

```go
package main

import (
	"fmt"

	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/sdk"
)

func main() {
	client := sdk.NewClient()

	models, err := client.Models().List()

	if err != nil {
		fmt.Println("Error listing models:", err)
		return
	}

	fmt.Printf("Models: %+v\n", models)
}

```

---

## **Summary of Available Methods**

| **Class**       | **Method**          | **Description**                           |
|------------------|---------------------|-------------------------------------------|
| **Deployments**  | `Create(payload)`   | Create a new deployment.                  |
| **Deployments**  | `GetByID(id)`       | Get deployment details by deployment ID.  |
| **Deployments**  | `List()`            | List all deployments.                     |
| **Deployments**  | `GetJob(jobID)`     | Get the status of a deployment job.       |
| **Models**       | `List()`            | List all available models.                |
| **Models**       | `GetByID(id)`       | Get model details by model ID.            |
| **Models**       | `GetByName(name)`   | Get model details by model name.          |

---