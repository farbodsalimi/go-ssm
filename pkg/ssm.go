package ssm

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"go-ssm/internal/utils"
)

// Run fetches ssm keys
func Run() {
	flag.Parse()

	SSMKeys, err := utils.GetEnvStr(utils.GetEnvArgs{Key: "SSM_KEYS"})
	if err != nil {
		log.Fatalf("failed to read SSM_KEYS: %v", err)
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	ssmsvc := ssm.NewFromConfig(cfg)
	withDecryption := true

	for key := range strings.SplitSeq(SSMKeys, ",") {
		param, err := ssmsvc.GetParameter(ctx, &ssm.GetParameterInput{
			Name:           &key,
			WithDecryption: &withDecryption,
		})
		if err != nil {
			log.Fatalf("failed to get SSM parameter %q: %v", key, err)
		}

		keyPath := strings.Split(*param.Parameter.Name, "/")
		keyName := keyPath[len(keyPath)-1]

		os.Setenv(keyName, *param.Parameter.Value)
	}

	path, err := exec.LookPath(flag.Args()[0])
	if err != nil {
		log.Fatalf("failed to find executable %q: %v", flag.Args()[0], err)
	}

	execErr := syscall.Exec(path, flag.Args()[0:], os.Environ())
	if execErr != nil {
		log.Fatalf("failed to exec %q: %v", path, execErr)
	}
}
