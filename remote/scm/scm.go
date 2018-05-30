package scm

import (
	"net"
	"net/url"

	"github.com/drone/drone/remote"
)

const DefaultScope = "api"

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

type SCM struct {
	URL        string
	Username   string
	Password   string
	SkipVerify bool
}
