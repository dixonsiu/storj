// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pb

import (
	"storj.io/storj/pkg/utils"
	"storj.io/storj/pkg/storj"
)

// NodeIDsToLookupRequests converts NodeIDs to LookupRequests
func NodeIDsToLookupRequests(nodeIDs storj.NodeIDList) *LookupRequests {
	var rq []*LookupRequest
	for _, v := range nodeIDs {
		r := &LookupRequest{NodeId: v}
		rq = append(rq, r)
	}
	return &LookupRequests{LookupRequest: rq}
}

// LookupResponsesToNodes converts LookupResponses to Nodes
func LookupResponsesToNodes(responses *LookupResponses) []*Node {
	var nodes []*Node
	for _, v := range responses.LookupResponse {
		n := v.GetNode()
		nodes = append(nodes, n)
	}
	return nodes
}

// NodesToIDs extracts Node-s into a list of ids
func NodesToIDs(nodes []*Node) (storj.NodeIDList, error) {
	ids := make(storj.NodeIDList, len(nodes))
	errs := make([]error, len(nodes))
	var err error
	for i, node := range nodes {
		if node != nil {
			ids[i], err = storj.NodeIDFromBytes(node.Id)
			if err != nil {
				errs[i] = err
			}
		}
	}
	return ids, utils.CombineErrors(errs...)
}

// CopyNode returns a deep copy of a node
// It would be better to use `proto.Clone` but it is curently incompatible
// with gogo's customtype extension.
// (see https://github.com/gogo/protobuf/issues/147)
func CopyNode(src *Node) (dst *Node) {
	node := Node{Id: []byte{}}
	copy(node.Id[:], src.Id[:])
	if src.Address != nil {
		node.Address = &NodeAddress{
			Transport: src.Address.Transport,
			Address:   src.Address.Address,
		}
	}
	if src.Metadata != nil {
		node.Metadata = &NodeMetadata{
			Email:  src.Metadata.Email,
			Wallet: src.Metadata.Wallet,
		}
	}
	if src.Restrictions != nil {
		node.Restrictions = &NodeRestrictions{
			FreeBandwidth: src.Restrictions.FreeBandwidth,
			FreeDisk:      src.Restrictions.FreeDisk,
		}
	}

	node.AuditSuccess = src.AuditSuccess
	node.IsUp = src.IsUp
	node.LatencyList = src.LatencyList
	node.Type = src.Type
	node.UpdateAuditSuccess = src.UpdateAuditSuccess
	node.UpdateLatency = src.UpdateLatency
	node.UpdateUptime = src.UpdateUptime

	return &node
}
