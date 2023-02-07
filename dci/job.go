package dci

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"github.com/rh-ecosystem-edge/dci-manager/internal"
)

func CreateJob(test common.Test) error {
	err := createJobFolders(test)
	if err != nil {
		return err
	}
	settings := newSettings(test)
	err = settings.WriteToFile()
	if err != nil {
		return err
	}
	err = writeRunScript(test)
	if err != nil {
		return err
	}
	// Generate hooks
	hooksFolder := jobHooksFolder(test)
	installHook := &hook{hookType: HookTypeInstall}
	err = installHook.WriteHook(hooksFolder, test.Setup())
	if err != nil {
		return err
	}
	testsHook := &hook{hookType: HookTypeTests}
	err = testsHook.WriteHook(hooksFolder, test.Steps())
	if err != nil {
		return err
	}
	postRunHook := &hook{hookType: HookTypePostRun}
	err = postRunHook.WriteHook(hooksFolder, nil)
	if err != nil {
		return err
	}
	preRunHook := &hook{hookType: HookTypePreRun}
	err = preRunHook.WriteHook(hooksFolder, nil)
	if err != nil {
		return err
	}
	teardownHook := &hook{hookType: HookTypeTeardown}
	err = teardownHook.WriteHook(hooksFolder, nil)
	if err != nil {
		return err
	}
	return nil
}

func createJobFolders(test common.Test) error {
	err := internal.CreateFolderIfNotExists(agentTestFolder(test))
	if err != nil {
		return err
	}
	err = internal.CreateFolderIfNotExists(jobFolder(test))
	if err != nil {
		return err
	}
	return internal.CreateFolderIfNotExists(jobHooksFolder(test))
}

func writeRunScript(test common.Test) error {
	r, err := newRunScript(test)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/run_job.sh", jobFolder(test))

	err = r.WriteToFile(filePath)
	if err != nil {
		return err
	}
	return os.Chmod(filePath, 0754)
}

func templateParseToFile(filePath string, templateStr string, obj interface{}) error {
	tmpl, err := template.New("t").Parse(templateStr)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, obj)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, b.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func RemoveJob(test common.Test) error {
	return os.RemoveAll(jobFolder(test))
}

func agentTestFolder(test common.Test) string {
	testNameLower := strings.ToLower(test.Name())
	testNameLower = strings.ReplaceAll(testNameLower, " ", "_")
	componentNameLower := strings.ToLower(test.ComponentName())
	componentNameLower = strings.ReplaceAll(componentNameLower, " ", "_")
	return fmt.Sprintf("%s/%s/%s/%s/", internal.JobsFolderLocation, test.AgentName(), componentNameLower, testNameLower)
}

func jobFolder(test common.Test) string {
	folder := fmt.Sprintf("%s/%s-%s/", agentTestFolder(test), test.TargetVersion(), test.AppVersion())
	absPath, err := filepath.Abs(folder)
	if err != nil {
		log.Fatal(err)
	}

	return absPath
}
func jobHooksFolder(test common.Test) string {
	return fmt.Sprintf("%s/hooks/", jobFolder(test))
}
