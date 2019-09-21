package tftest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// FindTerraform attempts to find a Terraform CLI executable for plugin testing.
//
// As a first preference it will look for the environment variable
// TFTEST_TERRAFORM and return its value. If that variable is not set, it will
// look in PATH for a program named "terraform" and, if one is found, return
// its absolute path.
//
// If no Terraform executable can be found, the result is the empty string. In
// that case, the test program will usually fail outright.
func FindTerraform() string {
	if p := os.Getenv("TFTEST_TERRAFORM"); p != "" {
		return p
	}
	p, err := exec.LookPath("terraform")
	if err != nil {
		return ""
	}
	return p
}

// RunTerraform runs the configured Terraform CLI executable with the given
// arguments, returning an error if it produces a non-successful exit status.
func (wd *WorkingDir) runTerraform(args ...string) error {
	allArgs := []string{"terraform"}
	allArgs = append(allArgs, args...)

	var env []string
	for _, e := range os.Environ() {
		env = append(env, e)
	}
	env = append(env, "TF_INPUT=0")
	env = append(env, "TF_LOG=") // so logging can't pollute our stderr output

	var errBuf strings.Builder

	// FIXME: Ideally in testing.Verbose mode we'd turn on Terraform DEBUG
	// logging, perhaps redirected to a separate fd other than stderr to avoid
	// polluting it, and then propagate the log lines out into t.Log so that
	// they are visible to the person running the test. Currently though,
	// Terraform CLI is able to send logs only to either an on-disk file or
	// to stderr.

	cmd := &exec.Cmd{
		Path:   wd.h.TerraformExecPath(),
		Args:   allArgs,
		Dir:    wd.baseDir,
		Stderr: &errBuf,
	}
	err := cmd.Run()
	if tErr, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("terraform failed: %s\n\nstderr:\n%s", tErr.ProcessState.String(), errBuf.String())
	}
	return err
}

// runTerraformJSON runs the configured Terraform CLI executable with the given
// arguments and tries to decode its stdout into the given target value (which
// must be a non-nil pointer) as JSON.
func (wd *WorkingDir) runTerraformJSON(target interface{}, args ...string) error {
	allArgs := []string{"terraform"}
	allArgs = append(allArgs, args...)

	var env []string
	for _, e := range os.Environ() {
		env = append(env, e)
	}
	env = append(env, "TF_INPUT=0")
	env = append(env, "TF_LOG=") // so logging can't pollute our stderr output

	var outBuf bytes.Buffer
	var errBuf strings.Builder

	cmd := &exec.Cmd{
		Path:   wd.h.TerraformExecPath(),
		Args:   allArgs,
		Dir:    wd.baseDir,
		Stderr: &errBuf,
		Stdout: &outBuf,
	}
	err := cmd.Run()
	if err != nil {
		if tErr, ok := err.(*exec.ExitError); ok {
			err = fmt.Errorf("terraform failed: %s\n\nstderr:\n%s", tErr.ProcessState.String(), errBuf.String())
		}
		return err
	}

	return json.Unmarshal(outBuf.Bytes(), target)
}
