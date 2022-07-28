# repo-scaffolding 插件

这个插件会基于一个脚手架仓库来初始化一个 Gihub 或者 GitLab 仓库。

## 运行需求

这个插件运行需要设置以下环境变量：

- GITHUB_TOKEN

在使用插件之前请先设置这个环境变量，如果你不知道如何获取这个 token，可以查看以下文档：

- [Creating a personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

*注意：*

- 如果你执行 `dtm delete` 命令，这个 GitHub 上的仓库将会被删除。

- 如果你执行 `dtm update` 命令,  这个 GitHub 上的仓库将会被删除然后重新创建。 

- 对于 `repo-scaffolding` 插件，目前只需要 token 有 `repo`, `delete_repo` 权限即可。 

## 使用方法

**请注意这里的设置参数都是大小写敏感的**

```yaml
--8<-- "repo-scaffolding.yaml"
```

在配置文件中替换以下配置：

### destination_repo

这个是目标仓库的配置，包括以下几个配置项：

- `YOUR_GITHUB_USERNAME`
- `YOUR_GITHUB_ORGANIZATION_NAME`
- `YOUR_GITHUB_REPO_NAME`
- `YOUR_GITHUB_REPO_MAIN_BRANCH`

`owner`，`org` 和 `repo` 目前是必填的，`branch` 的默认值是  "main"。

### source_repo

这个是源脚手架仓库的配置（目前只支持 Github），包括以下几个配置：

- `YOUR_TEMPLATE_REPO_ORG`
- `YOUR_TEMPLATE_REPO_NAME`

目前这两个配置项都是必填的。

### repo_type

这个配置用于设置推送的目标仓库类型，目前支持 `GitLab` 和 `github`。

### vars

这个配置用于设置渲染源脚手架仓库时的变量，以下变量会默认设置：

```json
{
    "AppName": destination_repo.repo,
    "Repo": {
        "Name": destination_repo.repo,
        "Owner": destination_repo.owner
    }
}
```

## 示例 

### 官方支持脚手架项目

以下仓库是用于在 `source_repo` 设置的官方脚手架仓库，你可以使用这些仓库或者创建自己的脚手架仓库。

| language | org | repo |
|  ----  | ----  |----  |
| Golang | devstream-io | dtm-scaffolding-golang |
| Java Spring | spring-guides | gs-spring-boot |


### Golang

```yaml
tools:
  - name: repo-scaffolding
    instanceID: golang-scaffolding
    options:
      destination_repo:
        owner: test_owner
        org: ""
        repo: dtm-test-golang
        branch: main
      repo_type: github
      source_repo:
        org: devstream-io
        repo: dtm-scaffolding-golang
      vars:
        ImageRepo: dtm-test/golang-repo
```

这个配置在 GitHub 为用于 test_owner 创建 `dtm-test-golang` 仓库，它的生成是基于 `devstream-io/dtm-scaffolding-golang` 官方 Golang 脚手架仓库和传入的变量 `ImageRepo`。

### Java Spring

```yaml
tools:
  - name: repo-scaffolding
    instanceID: java-scaffolding
    options:
      destination_repo:
        owner: test_owner
        org: ""
        repo: dtm-test-java
        branch: main
      repo_type: github
      source_repo:
        org: spring-guides
        repo: gs-spring-boot
```

这个配置会在 GitHub 为用户 test_owner 创建 `dtm-test-java` 仓库，使用的是 Spring 官方的 `spring-guides/gs-spring-boot` 仓库。

## Outputs

这个插件有以下三个输出：

- `owner`
- `repo`
- `repoURL`