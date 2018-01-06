package sitemap

import (
	"fmt"
	"time"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"github.com/joeguo/sitemap"
)

// GenerateSitemap grabs all entries (now only ones added with borg) and saves a sitemap.xml.gz file in `sitemapPath`
func GenerateSitemap(db *gorm.DB) {
	defer func() {
		if r := recover(); r != nil {
			log.Warnf("Sitemap generation failed: %v", r)
		}
	}()
	start := time.Now()

	projects := []domain.Project{}
	err := db.Preload("Issues").Preload("Issues.Comments").Find(&projects).Error
	if err != nil {
		panic(err)
	}

	users := []domain.User{}
	err = db.Preload("Posts").Find(&users).Error
	if err != nil {
		panic(err)
	}

	items := []*sitemap.Item{}
	for _, project := range projects {
		if config.IsTestUser(project.Author) {
			continue
		}
		item := &sitemap.Item{
			Loc:        fmt.Sprintf("%v/%v/%v", config.C.SiteUrl, project.Author, project.Name),
			LastMod:    project.UpdatedAt,
			Priority:   0.5,
			Changefreq: "daily",
		}
		items = append(items, item)
		for _, issue := range project.Issues {
			issueLastUpdated := issue.UpdatedAt
			for _, comment := range issue.Comments {
				if comment.UpdatedAt.After(issueLastUpdated) {
					issueLastUpdated = comment.UpdatedAt
				}
			}
			item := &sitemap.Item{
				Loc:        fmt.Sprintf("%v/%v/%v/issue/%v", config.C.SiteUrl, project.Author, project.Name, issue.Id),
				LastMod:    issueLastUpdated,
				Priority:   0.5,
				Changefreq: "daily",
			}
			items = append(items, item)
		}
	}
	for _, user := range users {
		if config.IsTestUser(user.Nick) {
			continue
		}
		item := &sitemap.Item{
			Loc:        fmt.Sprintf("%v/%v", config.C.SiteUrl, user.Nick),
			LastMod:    time.Now(),
			Priority:   0.5,
			Changefreq: "daily",
		}
		items = append(items, item)
		for _, post := range user.Posts {
			item := &sitemap.Item{
				Loc:        fmt.Sprintf("%v/%v/p/%v", config.C.SiteUrl, user.Nick, post.Id),
				LastMod:    post.UpdatedAt,
				Priority:   0.5,
				Changefreq: "daily",
			}
			items = append(items, item)
		}
	}
	sitemapPath := "/var/sitemap.xml.gz"
	if config.C.Sitemap.Path != "" {
		sitemapPath = config.C.Sitemap.Path
	}
	err = sitemap.SiteMap(sitemapPath, items)
	if err != nil {
		panic(err)
	}
	log.Infof("Generated sitemap successfully, took: %v", time.Now().Sub(start))
}
