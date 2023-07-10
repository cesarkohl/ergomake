package github

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ergomake/ergomake/internal/api/auth"
	"github.com/ergomake/ergomake/internal/logger"
)

func (ghr *githubRouter) configureRepo(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")

	authData, ok := auth.GetAuthData(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	isAuthorized, err := auth.IsAuthorized(c, owner, authData)
	if err != nil {
		logger.Ctx(c).Err(err).
			Str("owner", owner).
			Str("repo", repo).
			Msg("fail to check if user is authorized to configure repo")
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if !isAuthorized {
		c.JSON(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		return
	}

	ergopack := `apps:
  app:
    path: ../
    publicPort: 3000
`
	changes := map[string]string{
		".ergomake/ergopack.yaml": ergopack,
	}
	title := "Configure ergomake"
	description := "Hi 👋"

	pr, err := ghr.ghApp.CreatePullRequest(
		context.Background(),
		owner, repo, "ergomake",
		changes, title, description,
	)
	if err != nil {
		logger.Ctx(c).Err(err).
			Str("owner", owner).
			Str("repo", repo).
			Msg("fail to create pull request")
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{"pullRequestURL": pr.GetURL()})
}
