package python

import (
	"github.com/merico-dev/stream/internal/pkg/log"
	ga "github.com/merico-dev/stream/internal/pkg/plugin/githubactions"
	"github.com/merico-dev/stream/pkg/util/github"
)

func IsHealthy(options *map[string]interface{}) (bool, error) {
	opt, err := parseAndValidateOptions(options)
	if err != nil {
		return false, err
	}

	ghOptions := &github.Option{
		Owner:    opt.Owner,
		Repo:     opt.Repo,
		NeedAuth: true,
	}
	gitHubClient, err := github.NewClient(ghOptions)
	if err != nil {
		return false, err
	}

	log.Infof("Language is: %s.", ga.GetLanguage(opt.Language))

	retMap, err := gitHubClient.VerifyWorkflows(workflows)
	if err != nil {
		return false, err
	}

	healthy := true
	for name, err := range retMap {
		if err != nil {
			healthy = false
			log.Errorf("The workflow/file %s is not ok: %s", name, err)
		} else {
			log.Successf("The workflow/file %s is ok", name)
		}
	}

	return healthy, nil
}