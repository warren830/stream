package nodejs

import (
	ga "github.com/merico-dev/stream/internal/pkg/plugin/common/githubactions"
	github "github.com/merico-dev/stream/pkg/util/github"
)

var workflows = []*github.Workflow{
	{CommitMessage: ga.CommitMessage, WorkflowFileName: ga.MainBuilderFileName, WorkflowContent: mainPipeline},
}
