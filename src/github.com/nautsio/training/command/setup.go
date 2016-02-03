package command

import (
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

func Setup(c *cli.Context) {
	label := c.String("label")
	provider := c.String("provider")

	// Validate global flags
	if label == "" || provider == "" {
		log.Fatal("[ERROR] Missing required flags, see 'training setup --help' for usage information")
	}

	options := make(map[string]string)
	options["region"] = "eu-central-1"
	options["project"] = label
	options["count"] = strconv.Itoa(c.Int("count"))

	if _, err := os.Stat("./terraform/" + provider); err != nil {
		log.Fatal("[ERROR] - Provider directory not found")
	}

	switch provider {
	case "aws":
		aws(c, &options)
	default:
		log.Fatal("[ERROR] Unsupported provider")
	}

	tfCmd := "terraform plan"
	for k, v := range options {
		tfCmd += " -var " + k + "=" + v
	}
	log.Println(tfCmd)
}

func aws(c *cli.Context, options *map[string]string) {
	accessKey := c.String("aws-access-key")
	secretKey := c.String("aws-secret-key")

	if accessKey == "" || secretKey == "" {
		log.Fatal("[ERROR] Missing AWS specific flags, see 'training setup --help' for usage information")
	}

	(*options)["access_key"] = accessKey
	(*options)["secret_key"] = secretKey
}
