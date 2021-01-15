/* Node.go: Provides the necessary additional functionality to make the Node a lib.Message
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2018, Triad National Security, LLC
 * See LICENSE file for details.
 */

package proto

import (
	"github.com/hpc/kraken/lib/json"
)

// MarshalJSON creats a JSON version of Node
func (n *Node) MarshalJSON() ([]byte, error) {
	return json.MarshalJSON(n)
}

// UnmarshalJSON populates a node from JSON
func (n *Node) UnmarshalJSON(j []byte) error {
	return json.UnmarshalJSON(j, n)
}