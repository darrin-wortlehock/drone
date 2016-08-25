package custom

import (
	"github.com/drone/drone/model"
	"github.com/drone/drone/remote"
	"net/http"
	"fmt"
)

type Opts struct {
	URL	string
}

type client struct {
	URL	string
}


func New(opts Opts) (remote.Remote, error) {
	remote := &client{
		URL:	opts.URL,
	}
	return remote, nil
}

func (c *client) Login(w http.ResponseWriter, r *http.Request) (*model.User, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) Auth(token, secret string) (string, error) {
	return "", fmt.Errorf("Not Implemented")
}

func (c *client) Teams(u *model.User) ([]*model.Team, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) Repo(u *model.User, owner, name string) (*model.Repo, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) Repos(u *model.User) ([]*model.RepoLite, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) Perm(u *model.User, owner, name string) (*model.Perm, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) File(u *model.User, r *model.Repo, b *model.Build, f string) ([]byte, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) Status(u *model.User, r *model.Repo, b *model.Build, link string) error {
	return fmt.Errorf("Not Implemented")
}

func (c *client) Netrc(u *model.User, r *model.Repo) (*model.Netrc, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (c *client) Activate(u *model.User, r *model.Repo, link string) error {
	return fmt.Errorf("Not Implemented")
}

func (c *client) Deactivate(u *model.User, r *model.Repo, link string) error {
	return fmt.Errorf("Not Implemented")
}

func (c *client) Hook(r *http.Request) (*model.Repo, *model.Build, error) {
	return nil, nil, fmt.Errorf("Not Implemented")
}

