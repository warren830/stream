package golang

import (
	"github.com/spf13/viper"

	ga "github.com/merico-dev/stream/internal/pkg/plugin/common/githubactions"

	"github.com/merico-dev/stream/pkg/util/github"
	"github.com/merico-dev/stream/pkg/util/log"
)

// Update remove and set up GitHub Actions workflows.
func Update(options map[string]interface{}) (map[string]interface{}, error) {
	opts, err := parseAndValidateOptions(options)
	if err != nil {
		return nil, err
	}

	ghOptions := &github.Option{
		Owner:    opts.Owner,
		Repo:     opts.Repo,
		NeedAuth: true,
	}
	ghClient, err := github.NewClient(ghOptions)
	if err != nil {
		return nil, err
	}

	log.Debugf("Language is %s.", ga.GetLanguage(opts.Language))

	if opts.Docker != nil && opts.Docker.Enable {
		for _, secret := range []string{"DOCKERHUB_USERNAME", "DOCKERHUB_TOKEN"} {
			if err := ghClient.DeleteRepoSecret(secret); err != nil {
				return nil, err
			}
		}

		if err := ghClient.AddRepoSecret("DOCKERHUB_USERNAME", viper.GetString("dockerhub_username")); err != nil {
			return nil, err
		}
		if err := ghClient.AddRepoSecret("DOCKERHUB_TOKEN", viper.GetString("dockerhub_token")); err != nil {
			return nil, err
		}
	}

	for _, pipeline := range workflows {
		err := ghClient.DeleteWorkflow(pipeline, opts.Branch)
		if err != nil {
			return nil, err
		}

		content, err := renderTemplate(pipeline, opts)
		if err != nil {
			return nil, err
		}
		pipeline.WorkflowContent = content

		err = ghClient.AddWorkflow(pipeline, opts.Branch)
		if err != nil {
			return nil, err
		}
	}

	return ga.BuildState(opts.Owner, opts.Repo), nil
}
