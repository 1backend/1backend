package deploy

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
	"time"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/state"
	tp "github.com/1backend/1backend/backend/tech-pack"
)

type Deployer struct {
	db    *gorm.DB
	state *state.State
}

func NewDeployer(db *gorm.DB, state *state.State) Deployer {
	return Deployer{
		db:    db,
		state: state,
	}
}

func (d Deployer) Deploy(project *domain.Project) error {
	var err error
	id := domain.Sid.MustGenerate()
	build := &domain.Build{
		Id:         id,
		ProjectId:  project.Id,
		InProgress: true,
	}
	defer func() {
		if err != nil {
			log.Error(err)
			build.Output += "\n" + err.Error()
			build.Success = false
			build.InProgress = false
			d.db.Save(build)
		}
	}()
	build.CreatedAt = time.Now()
	build.UpdatedAt = time.Now()
	err = d.db.Create(build).Error
	if err != nil {
		return err
	}
	techPack := tp.GetProvider(project)
	recipePath := techPack.RecipePath()
	dat, err := ioutil.ReadFile(config.C.Path + "/tech-pack/" + recipePath + "/code.tpl")
	if err != nil {
		return err
	}
	// Create a new template and parse the letter into it.
	templFuncs := template.FuncMap{
		"trim": strings.TrimSpace,
	}
	techPack.AddTemplateFuncs(&templFuncs)
	t, err := template.New("code").Funcs(templFuncs).Parse(string(dat))
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, map[string]interface{}{
		"Project": project,
	})
	if err != nil {
		return err
	}
	buildPath := config.C.Path + "/builds/" + project.Author + "/" + project.Name + "/" + id
	err = os.MkdirAll(buildPath, 0700)
	if err != nil {
		return err
	}
	for _, v := range techPack.FilesToBuild() {
		f, err := os.Create(buildPath + "/" + v[0])
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write([]byte(v[1]))
		if err != nil {
			return err
		}
	}
	f, err := os.Create(buildPath + "/" + techPack.Outfile())
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}
	output, err := exec.Command("/bin/bash", config.C.Path+"/bash/build.sh", buildPath, project.Author, project.Name, project.InfraPassword, recipePath, config.C.Path).CombinedOutput()
	build.Output = string(output)
	build.Success = err == nil
	build.InProgress = false
	output, err = exec.Command("/bin/bash", config.C.Path+"/bash/get-port.sh", project.Author, project.Name).CombinedOutput()
	if err != nil {
		return err
	}
	port, err := strconv.ParseInt(string(output), 10, 64)
	if err != nil {
		return err
	}
	err = d.state.SetPort(project.Author, project.Name, int(port))
	if err != nil {
		return err
	}
	err = d.db.Table("projects").Where("id = ?", project.Id).Update(map[string]interface{}{
		"port": string(output),
	}).Error
	outp, err := d.GenerateAPIs(project, build.Id)
	if err != nil {
		build.Output += "\n" + outp + "\n" + err.Error()
		build.Success = false
	}
	err = d.db.Save(build).Error
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
