//go:generate mockgen -destination mock_nodeconf/mock_nodeconf.go github.com/anytypeio/any-sync/nodeconf Service,Configuration
package nodeconf

import "github.com/anytypeio/go-chash"

type Configuration interface {
	// Id returns current nodeconf id
	Id() string
	// NodeIds returns list of peerId for given spaceId
	NodeIds(spaceId string) []string
	// IsResponsible checks if current account responsible for given spaceId
	IsResponsible(spaceId string) bool
	// FilePeers returns list of filenodes
	FilePeers() []string
	// ConsensusPeers returns list of consensusnodes
	ConsensusPeers() []string
	// Addresses returns map[peerId][]addr with connection addresses for all known nodes
	Addresses() map[string][]string
}

type configuration struct {
	id             string
	accountId      string
	filePeers      []string
	consensusPeers []string
	chash          chash.CHash
	allMembers     []NodeConfig
}

func (c *configuration) Id() string {
	return c.id
}

func (c *configuration) NodeIds(spaceId string) []string {
	members := c.chash.GetMembers(spaceId)
	res := make([]string, 0, len(members))
	for _, m := range members {
		if m.Id() != c.accountId {
			res = append(res, m.Id())
		}
	}
	return res
}

func (c *configuration) IsResponsible(spaceId string) bool {
	for _, m := range c.chash.GetMembers(spaceId) {
		if m.Id() == c.accountId {
			return true
		}
	}
	return false
}

func (c *configuration) FilePeers() []string {
	return c.filePeers
}

func (c *configuration) ConsensusPeers() []string {
	return c.consensusPeers
}

func (c *configuration) Addresses() map[string][]string {
	res := make(map[string][]string)
	for _, m := range c.allMembers {
		res[m.PeerId] = m.Addresses
	}
	return res
}