package nodejs

import (
	"fmt"

	"github.com/merico-dev/stream/internal/pkg/plugin/common/githubactions"

	"github.com/mitchellh/mapstructure"

	"github.com/merico-dev/stream/pkg/util/github"
	"github.com/merico-dev/stream/pkg/util/log"
)

// Create sets up GitHub Actions workflow(s).
func Create(options map[string]interface{}) (map[string]interface{}, error) {
	var opts Options

	err := mapstructure.Decode(options, &opts)
	if err != nil {
		return nil, err
	}

	ghOptions := &github.Option{
		Owner:    opts.Owner,
		Repo:     opts.Repo,
		NeedAuth: true,
	}

	if errs := validate(&opts); len(errs) != 0 {
		for _, e := range errs {
			log.Errorf("Options error: %s.", e)
		}
		return nil, fmt.Errorf("opts are illegal")
	}

	ghClient, err := github.NewClient(ghOptions)
	if err != nil {
		return nil, err
	}

	log.Debugf("Language is: %s.", githubactions.GetLanguage(opts.Language))

	for _, w := range workflows {
		if err := ghClient.AddWorkflow(w, opts.Branch); err != nil {
			return nil, err
		}
	}

	return githubactions.BuildState(opts.Owner, opts.Repo), nil
}
