/* Client.go: extension adds special fields tracking generic Client/iClient state
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2018, Triad National Security, LLC
 * See LICENSE file for details.
 */

package pxe

import (
	"github.com/hpc/kraken/core"
	"github.com/hpc/kraken/lib/json"
	"github.com/hpc/kraken/lib/types"
)

//go:generate protoc -I ../../core/proto -I . --go_out=plugins=grpc:. pxe.proto

const Name = "type.googleapis.com/PXE.Client"

/////////////////
// Client Object /
///////////////

var _ types.Extension = (*Client)(nil)

func (*Client) New() types.Message {
	return &Client{}
}

func (*Client) Name() string {
	return Name
}

// MarshalJSON creats a JSON version of Node
func (c *Client) MarshalJSON() ([]byte, error) {
	return json.MarshalJSON(c)
}

// UnmarshalJSON populates a node from JSON
func (c *Client) UnmarshalJSON(j []byte) error {
	return json.UnmarshalJSON(j, c)
}

func init() {
	core.Registry.RegisterExtension(&Client{})
}