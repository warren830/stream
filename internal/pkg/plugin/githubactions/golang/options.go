package golang

import (
	"fmt"

	ga "github.com/merico-dev/stream/internal/pkg/plugin/common/githubactions"
)

// TODO(daniel-hutao): Options should keep as same as other plugins named Param
// Options is the struct for configurations of the githubactions plugin.
type Options struct {
	Owner    string
	Repo     string
	Branch   string
	Language *ga.Language
	Build    *Build
	Test     *Test
	Docker   *Docker
}

// validate validates the options provided by the core.
func validate(opts *Options) []error {
	retErrors := make([]error, 0)

	// owner/repo/branch
	if opts.Owner == "" {
		retErrors = append(retErrors, fmt.Errorf("owner is empty"))
	}
	if opts.Repo == "" {
		retErrors = append(retErrors, fmt.Errorf("repo is empty"))
	}
	if opts.Branch == "" {
		retErrors = append(retErrors, fmt.Errorf("branch is empty"))
	}

	// language
	if opts.Language == nil {
		retErrors = append(retErrors, fmt.Errorf("language is empty"))
	}
	if errs := opts.Language.Validate(); len(errs) != 0 {
		for _, e := range errs {
			retErrors = append(retErrors, fmt.Errorf("language is invalid: %s", e))
		}
	}

	// jobs
	if opts.Test == nil {
		retErrors = append(retErrors, fmt.Errorf("test is empty"))
	}
	if errs := opts.Test.Validate(); len(errs) != 0 {
		for _, e := range errs {
			retErrors = append(retErrors, fmt.Errorf("test is invalid: %s", e))
		}
	}

	if opts.Docker == nil {
		return retErrors
	}

	if errs := opts.Docker.Validate(); len(errs) != 0 {
		for _, e := range errs {
			retErrors = append(retErrors, fmt.Errorf("docker is invalid: %s", e))
		}
	}

	return retErrors
}
