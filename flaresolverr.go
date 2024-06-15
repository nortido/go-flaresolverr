package flaresolverr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	commandGet  = "request.get"
	commandPost = "request.post"
)

type Config struct {
	Host       string
	MaxTimeout int
}

type IClient interface {
	Send(r *http.Request, opts *Payload) (result *Response, err error)
}

type client struct {
	hc   *http.Client
	conf *Config
}

func (f *client) Send(r *http.Request, opts *Payload) (result *Response, err error) {
	var (
		b    []byte
		data []byte
		resp *http.Response
	)

	switch r.Method {
	case http.MethodGet:
		opts.Cmd = commandGet
	case http.MethodPost:
		if b, err = io.ReadAll(r.Body); err != nil {
			err = fmt.Errorf("error reading proxy request body")
			return
		}
		opts.PostData = string(b)
		opts.Cmd = commandPost
	default:
		err = fmt.Errorf("unknown http method: %s", r.Method)
		return
	}
	opts.MaxTimeout = f.conf.MaxTimeout
	if data, err = json.Marshal(opts); err != nil {
		return
	}
	if resp, err = f.hc.Post(f.conf.Host, "application/json", bytes.NewBuffer(data)); err != nil {
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}
	return
}

func New(hc *http.Client, conf *Config) IClient {
	return &client{
		hc:   hc,
		conf: conf,
	}
}
