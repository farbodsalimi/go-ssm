package ssm

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"go-ssm/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// Run fetches ssm keys
func Run() {
	flag.Parse()

	SSMKeys, err := utils.GetEnvStr(utils.GetEnvArgs{Key: "SSM_KEYS"})
	utils.ErrorExit(err)

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	utils.ErrorExit(err)

	ssmsvc := ssm.NewFromConfig(cfg)
	withDecryption := true

	for key := range strings.SplitSeq(SSMKeys, ",") {
		param, err := ssmsvc.GetParameter(ctx, &ssm.GetParameterInput{
			Name:           &key,
			WithDecryption: &withDecryption,
		})
		utils.ErrorExit(err)

		keyPath := strings.Split(*param.Parameter.Name, "/")
		keyName := keyPath[len(keyPath)-1]

		os.Setenv(keyName, *param.Parameter.Value)
	}

	path, err := exec.LookPath(flag.Args()[0])
	utils.ErrorExit(err)

	execErr := syscall.Exec(path, flag.Args()[0:], os.Environ())
	utils.ErrorExit(execErr)
}
