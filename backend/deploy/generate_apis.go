package deploy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	filePath "path/filepath"
	"strconv"
	"strings"

	"github.com/1backend/1backend/backend/client-plugins"
	apiTypes "github.com/1backend/1backend/backend/client-plugins/types"
	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
)

// @todo revisit this output return: likely output should go into error
// because the error is only something like "exit status 1"
func (d Deployer) GenerateAPIs(project *domain.Project, build *domain.Build) error {
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
		return nil
	}
	projects := []domain.Project{}
	err := d.db.Where("author = ? AND name <> ?", project.Author, project.Name).Find(&projects).Error
	if err != nil {
		return err
	}
	projectNames := []string{}
	for _, v := range projects {
		projectNames = append(projectNames, v.Name)
	}
	reposPath := config.C.Path + "/repos"
	err = os.MkdirAll(reposPath, 0700)
	if err != nil {
		return err
	}
	if !config.IsTestUser(project.Author) {
		build.Output += "### Creating GitHub repo\n"
		createOutput, err := exec.Command("/bin/bash", config.C.Path+"/bash/create-git-repo.sh",
			reposPath,
			project.Author,
			project.Name,
			config.C.ApiGeneration.GithubOrganisation,
			config.C.ApiGeneration.GithubUser,
			config.C.ApiGeneration.GithubPersonalToken,
		).CombinedOutput()
		build.Output += "```sh\n" + string(createOutput) + "\n```\n"
		if err != nil {
			return err
		}
	}
	generators := clientplugins.Plugins(project)
	repoPath := reposPath + "/" + project.Author
	context, err := apiTypes.GetContext(project, projectNames)
	if err != nil {
		return err
	}
	build.Output += "### Generating client libraries\n"
	for _, gen := range generators {
		build.Output += "#### Generating " + gen.Name() + " clients\n"
		files, err := gen.ClientFiles(*context)
		if err != nil {
			return err
		}
		for _, v := range files.Files {
			fileName := v[0]
			fileContents := v[1]
			fPath := repoPath + "/" + files.FolderName + "/" + project.Name + "/" + fileName
			err = os.MkdirAll(filePath.Dir(fPath), 0700)
			if err != nil {
				return err
			}
			f, err := os.Create(fPath)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = f.Write([]byte(fileContents))
			if err != nil {
				return err
			}
		}
	}
	if config.IsTestUser(project.Author) {
		return nil
	}
	npmPublishOutput := []byte{}
	if config.C.NpmPublication.Enabled {
		build.Output += "### Publishing Angular client to NPM\n"
		npmPublishOutput, err = exec.Command("/bin/bash", config.C.Path+"/bash/npm-publish.sh",
			reposPath,
			project.Author,
			project.Name,
			config.C.NpmPublication.NpmToken,
		).CombinedOutput()
		build.Output += "```sh\n" + string(npmPublishOutput) + "\n```\n"
		if err != nil {
			return err
		}
	}
	build.Output += "### Pushing client sources to GitHub\n"
	commitOutput, err := exec.Command("/bin/bash", config.C.Path+"/bash/git-commit-api.sh",
		reposPath,
		project.Author,
		project.Name,
		config.C.ApiGeneration.GithubOrganisation,
		config.C.ApiGeneration.GithubUser,
		config.C.ApiGeneration.GithubPersonalToken,
		build.Id,
	).CombinedOutput()
	build.Output += "```sh\n" + string(commitOutput) + "\n```\n"
	if err != nil {
		return err
	}
	return nil
}

func bumpPackageJsonVersion(filePath string) error {
	pjsonContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(pjsonContents, &m)
	if err != nil {
		return err
	}
	semVerParts := strings.Split(m["version"].(string), ".")
	patchNumber, _ := strconv.ParseInt(semVerParts[2], 10, 64)
	newSemVer := strings.Join([]string{semVerParts[0], semVerParts[1], fmt.Sprintf("%v", patchNumber)}, ".")
	m["version"] = newSemVer
	newPackageJson, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, newPackageJson, 0700)
}
