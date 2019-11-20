// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/plugin/environ"
)

// New returns a new environ plugin.
func New(param1, param2 string) environ.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) List(ctx context.Context, req *environ.Request) (map[string]string, error) {
	// TODO replace or remove
	// we could only expose environment variables to
	// specific repositories or organizations.
	if req.Repo.Namespace != p.param1 {
		return nil, nil
	}

	// TODO replace or remove
	// return a list of static environment variables
	// this is for demo purposes only.
	environ := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}

	return environ, nil
}
