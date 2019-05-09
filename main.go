package main

import (
	"flag"
	"go-ssm/src/utils"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	flag.Parse()

	SSMKeys, err := utils.GetEnvStr(utils.GetEnvArgs{Key: "SSM_KEYS"})
	utils.ErrorExit(err)

	sess, err := session.NewSession()
	utils.ErrorExit(err)

	ssmsvc := ssm.New(sess)
	withDecryption := true

	for _, key := range strings.Split(SSMKeys, ",") {
		param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
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
