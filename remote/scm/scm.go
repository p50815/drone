package scm

import (
	"net"
	"net/http"
	"net/url"

	"github.com/drone/drone/model"
	"github.com/drone/drone/remote"
)

// Opts defines configuration options.
type Opts struct {
	URL        string // Gogs server url.
	Username   string // Optional machine account username.
	Password   string // Optional machine account password.
	SkipVerify bool   // Skip ssl verification.
}

// New returns a Remote implementation that integrates with SCM
func New(opts Opts) (remote.Remote, error) {
	url, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}
	host, _, err := net.SplitHostPort(url.Host)
	if err == nil {
		url.Host = host
	}
	return &SCM{
		URL:        opts.URL,
		Username:   opts.Username,
		Password:   opts.Password,
		SkipVerify: opts.SkipVerify,
	}, nil
}

// SCM define configuration options.
type SCM struct {
	URL        string
	Username   string
	Password   string
	SkipVerify bool
}

// Login authenticates the session and returns the remote user details.
func (c *SCM) Login(res http.ResponseWriter, req *http.Request) (*model.User, error) {
	// TODO
	return nil, nil
}

// Auth returns the SCM user login for the given access token.
func (c *SCM) Auth(token, secret string) (string, error) {
	// TODO
	return "", nil
}

// Teams returns a list of all team membership for the SCM account.
func (c *SCM) Teams(u *model.User) ([]*model.Team, error) {
	// TODO
	return nil, nil
}

// Repo returns the named SCM repository.
func (c *SCM) Repo(u *model.User, owner, name string) (*model.Repo, error) {
	// TODO
	return nil, nil
}

// Repos returns a list of all repositories for SCM account, including
// organization repositories.
func (c *SCM) Repos(u *model.User) ([]*model.Repo, error) {
	// TODO
	return nil, nil
}

// Perm returns the user permissions for the named SCM repository.
func (c *SCM) Perm(u *model.User, owner, name string) (*model.Perm, error) {
	// TODO
	return nil, nil
}

// File fetches a file from the remote repository and returns in string format.
func (c *SCM) File(user *model.User, repo *model.Repo, build *model.Build, f string) ([]byte, error) {
	// TODO
	return nil, nil
}

// FileRef fetches the file from the SCM repository and returns its contents.
func (c *SCM) FileRef(u *model.User, r *model.Repo, ref, f string) ([]byte, error) {
	// TODO
	return nil, nil
}

// Netrc returns a netrc file capable of authenticating SCM requests and
// cloning Gitlab repositories. The netrc will use the global machine account
// when configured.
func (c *SCM) Netrc(u *model.User, r *model.Repo) (*model.Netrc, error) {
	// TODO
	return nil, nil
}

// Deactivate deactives the repository be removing registered push hooks from
// the GitHub repository.
func (c *SCM) Deactivate(u *model.User, r *model.Repo, link string) error {
	// TODO
	return nil
}

// Status sends the commit status to the remote system.
// An example would be the GitHub pull request status.
func (c *SCM) Status(u *model.User, r *model.Repo, b *model.Build, link string) error {
	// TODO
	return nil
}

// Activate activates a repository by creating the post-commit hook and
// adding the SSH deploy key, if applicable.
func (c *SCM) Activate(u *model.User, r *model.Repo, link string) error {
	// TODO
	return nil
}

// Hook parses the post-commit hook from the Request body
// and returns the required data in a standard format.
func (c *SCM) Hook(r *http.Request) (*model.Repo, *model.Build, error) {
	return nil, nil, nil
}
