package nodejs

import (
	"fmt"

	"github.com/merico-dev/stream/internal/pkg/plugin/common/githubactions"

	"github.com/mitchellh/mapstructure"

	"github.com/merico-dev/stream/pkg/util/github"
	"github.com/merico-dev/stream/pkg/util/log"
)

func Read(options map[string]interface{}) (map[string]interface{}, error) {
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

	path, err := ghClient.GetWorkflowPath()
	if err != nil {
		return nil, err
	}
	if path == "" {
		// file not found
		return nil, nil
	}

	log.Debugf("Language is: %s.", githubactions.GetLanguage(opts.Language))

	return githubactions.BuildReadState(path), nil
}
