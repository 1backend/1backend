package endpoints

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/1backend/1backend/backend/deploy"
	"github.com/1backend/1backend/backend/domain"
	tp "github.com/1backend/1backend/backend/tech-pack"
)

func (e Endpoints) CreateProject(proj *domain.Project) error {
	if !reg.Match([]byte(proj.Name)) {
		return errors.New("Allowed project name characters: lowercase letters, numbers, dash")
	}
	proj.Version = "0.0.1"
	proj.Description = "Change this to something sensible"
	proj.Id = domain.Sid.MustGenerate()
	proj.CallerId = domain.Sid.MustGenerate()
	proj.InfraPassword = domain.Sid.MustGenerate()
	proj.CreatedAt = time.Now()
	proj.UpdatedAt = time.Now()
	proj.CallerId = domain.Sid.MustGenerate()
	err := e.state.SetCallerIdToNameSpace(proj.CallerId, proj.Namespace)
	if err != nil {
		return err
	}
	tp.GetProvider(proj).CreateProjectPlugin()
	err = e.db.Save(proj).Error
	if err != nil {
		return err
	}
	for _, v := range proj.Endpoints {
		v.Id = domain.Sid.MustGenerate()
		v.ProjectId = proj.Id
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	for _, v := range proj.Dependencies {
		v.Id = domain.Sid.MustGenerate()
		v.ProjectId = proj.Id
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	go deploy.NewDeployer(e.db, e.state).Deploy(proj)
	return nil
}

func (e Endpoints) UpdateProject(proj *domain.Project) error {
	if !reg.Match([]byte(proj.Name)) {
		return errors.New("Allowed project name characters: lowercase letters, numbers, dash")
	}
	if proj.Version == "" {
		proj.Version = "0.0.2"
	} else {
		semVerParts := strings.Split(proj.Version, ".")
		patchNumber, _ := strconv.ParseInt(semVerParts[2], 10, 64)
		patchNumber++
		proj.Version = strings.Join([]string{semVerParts[0], semVerParts[1], fmt.Sprintf("%v", patchNumber)}, ".")
	}
	proj.CallerId = domain.Sid.MustGenerate()
	err := e.state.SetCallerIdToNameSpace(proj.CallerId, proj.Namespace)
	if err != nil {
		return err
	}
	// @todo this continous regeneration of infra password is likely a bad idea.
	// it requires a user credential flush in SQL (see build script in bash/build.sh)
	proj.InfraPassword = domain.Sid.MustGenerate()
	if proj.CallerId == "" {
		proj.CallerId = domain.Sid.MustGenerate()
	}
	err = e.db.Save(proj).Error
	if err != nil {
		return err
	}
	for _, v := range proj.Endpoints {
		if v.Id == "" {
			v.Id = domain.Sid.MustGenerate()
		}
		v.ProjectId = proj.Id
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	for _, v := range proj.Dependencies {
		if v.Id == "" {
			v.Id = domain.Sid.MustGenerate()
		}
		v.ProjectId = proj.Id
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	go deploy.NewDeployer(e.db, e.state).Deploy(proj)
	return nil
}
