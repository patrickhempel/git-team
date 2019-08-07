package disable

import (
	"errors"
	"os"
	"testing"

	"github.com/hekmekk/git-team/src/config"
)

var (
	unsetCommitTemplate = func() error { return nil }
	cfg                 = config.Config{TemplateFileName: "TEMPLATE_FILE", BaseDir: "BASE_DIR", StatusFileName: "STATUS_FILE"}
	loadConfig          = func() (config.Config, error) { return cfg, nil }
	fileInfo            os.FileInfo
	statFile            = func(string) (os.FileInfo, error) { return fileInfo, nil }
	removeFile          = func(string) error { return nil }
	persistDisabled     = func() error { return nil }
)

func TestDisableSucceeds(t *testing.T) {
	deps := dependencies{
		GitUnsetCommitTemplate: unsetCommitTemplate,
		LoadConfig:             loadConfig,
		StatFile:               statFile,
		RemoveFile:             removeFile,
		PersistDisabled:        persistDisabled,
	}
	err := executorFactory(deps)()

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestDisableShouldSucceedWhenUnsetCommitTemplateFailsBecauseItWasAlreadyUnset(t *testing.T) {
	deps := dependencies{
		GitUnsetCommitTemplate: func() error { return errors.New("exit status 5") },
		LoadConfig:             loadConfig,
		StatFile:               statFile,
		RemoveFile:             removeFile,
		PersistDisabled:        persistDisabled,
	}

	err := executorFactory(deps)()

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestDisableShouldFailWhenUnsetCommitTemplateFails(t *testing.T) {
	expectedErr := errors.New("failed to unset commit template")
	deps := dependencies{
		GitUnsetCommitTemplate: func() error { return expectedErr },
	}

	err := executorFactory(deps)()

	if err == nil || expectedErr != err {
		t.Errorf("expected: %s, received: %s", expectedErr, err)
		t.Fail()
	}
}

func TestDisableShouldFailWhenLoadConfigFails(t *testing.T) {
	expectedErr := errors.New("failed to load config")
	deps := dependencies{
		GitUnsetCommitTemplate: unsetCommitTemplate,
		LoadConfig:             func() (config.Config, error) { return config.Config{}, expectedErr },
	}

	err := executorFactory(deps)()

	if err == nil || expectedErr != err {
		t.Errorf("expected: %s, received: %s", expectedErr, err)
		t.Fail()
	}
}

func TestDisableShouldSucceedButNotTryToRemoveTheCommitTemplateFileWhenStatFileFails(t *testing.T) {
	deps := dependencies{
		GitUnsetCommitTemplate: unsetCommitTemplate,
		LoadConfig:             loadConfig,
		StatFile:               func(string) (os.FileInfo, error) { return fileInfo, errors.New("failed to stat file") },
		PersistDisabled:        persistDisabled,
	}

	err := executorFactory(deps)()

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestDisableShouldFailWhenRemoveFileFails(t *testing.T) {
	expectedErr := errors.New("failed to remove file")
	deps := dependencies{
		GitUnsetCommitTemplate: unsetCommitTemplate,
		LoadConfig:             loadConfig,
		StatFile:               statFile,
		RemoveFile:             func(string) error { return expectedErr },
	}

	err := executorFactory(deps)()

	if err == nil || expectedErr != err {
		t.Errorf("expected: %s, received: %s", expectedErr, err)
		t.Fail()
	}
}

func TestDisableShouldFailWhenpersistDisabledFails(t *testing.T) {
	expectedErr := errors.New("failed to save status")
	deps := dependencies{
		GitUnsetCommitTemplate: unsetCommitTemplate,
		LoadConfig:             loadConfig,
		StatFile:               statFile,
		RemoveFile:             removeFile,
		PersistDisabled:        func() error { return expectedErr },
	}

	err := executorFactory(deps)()

	if err == nil || expectedErr != err {
		t.Errorf("expected: %s, received: %s", expectedErr, err)
		t.Fail()
	}
}
