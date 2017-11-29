package deploy

import (
	"os"
	"os/exec"
	filePath "path/filepath"

	"github.com/1backend/1backend/backend/api-pack"
	apiTypes "github.com/1backend/1backend/backend/api-pack/types"
	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
)

// @todo revisit this output return: likely output should go into error
// because the error is only something like "exit status 1"
func (d Deployer) GenerateAPIs(project *domain.Project, buildId string) (string, error) {
	noDefs := true
	if project.Types != "" {
		noDefs = false
	}
	for _, v := range project.Endpoints {
		if v.Input != "" || v.Output != "" {
			noDefs = false
		}
	}
	if noDefs {
		return "", nil
	}
	projects := []domain.Project{}
	err := d.db.Where("author = ? AND name <> ?", project.Author, project.Name).Find(&projects).Error
	if err != nil {
		return "", err
	}
	projectNames := []string{}
	for _, v := range projects {
		projectNames = append(projectNames, v.Name)
	}
	context, err := apiTypes.GetContext(project, projectNames)
	if err != nil {
		return "", err
	}
	reposPath := config.C.Path + "/repos"
	err = os.MkdirAll(reposPath, 0700)
	if err != nil {
		return "", err
	}
	output, err := exec.Command("/bin/bash", config.C.Path+"/bash/create-git-repo.sh",
		reposPath,
		project.Author,
		project.Name,
		config.C.ApiGeneration.GithubOrganisation,
		config.C.ApiGeneration.GithubUser,
		config.C.ApiGeneration.GithubPersonalToken,
	).CombinedOutput()
	if err != nil {
		return string(output), err
	}
	generators := apipack.Generators(project)
	repoPath := reposPath + "/" + project.Author
	for _, gen := range generators {
		fileTuples, err := gen.FilesToBuild(*context)
		if err != nil {
			return "", err
		}
		for _, v := range fileTuples {
			fileName := v[0]
			fileContents := v[1]
			fPath := repoPath + "/" + gen.FolderName() + "/" + fileName
			err = os.MkdirAll(filePath.Dir(fPath), 0700)
			if err != nil {
				return "", err
			}
			f, err := os.Create(fPath)
			if err != nil {
				return "", err
			}
			defer f.Close()
			_, err = f.Write([]byte(fileContents))
			if err != nil {
				return "", err
			}
		}
	}
	output, err = exec.Command("/bin/bash", config.C.Path+"/bash/git-commit-api.sh",
		reposPath,
		project.Author,
		project.Name,
		config.C.ApiGeneration.GithubOrganisation,
		config.C.ApiGeneration.GithubUser,
		config.C.ApiGeneration.GithubPersonalToken,
		buildId,
	).CombinedOutput()
	if err != nil {
		return string(output), err
	}
	return "", nil
}
