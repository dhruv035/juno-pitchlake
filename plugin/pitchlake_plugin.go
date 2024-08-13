package pitchlake_plugin

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
)

type Plugin interface {
	// Init is called when the plugin is loaded.
	Init() error

	// Shutdown is called when the plugin is unloaded or the node is shutting down.
	Shutdown() error

	//  NewBlock is called whenever a new block (including state update and any new classes) is added to the blockchain.
	// This function is a hook called by Juno.
	NewBlock(block core.Block, stateUpdate core.StateUpdate, newClasses map[*felt.Felt]core.Class) error

	// The state is reverted by applying a write operation with the reverseStateDiff's StorageDiffs, Nonces, and ReplacedClasses,
	// and a delete option with its DeclaredV0Classes, DeclaredV1Classes, and ReplacedClasses.
	RevertBlock(from, to *BlockAndStateUpdate, reverseStateDiff *core.StateDiff) error

	// Note: We should only start working on the next two functions when
	// we implement and test an interface with the four functions above.

	// ContractStorage provides access to the contracts storage.
	// This function is not a Juno hook, the plugin calls it whenever needed.
	// Class returns the declared class associated with the provided classHash
	// This function is not a Juno hook, the plugin calls it whenever needed.
	Class(classHash *felt.Felt) (*core.DeclaredClass, error)
}

type BlockAndStateUpdate struct {
	Block       *core.Block
	StateUpdate *core.StateUpdate
}

type PitchlakePlugin struct{}

func (p *PitchlakePlugin) Init() error {
	return nil
}

func (p *PitchlakePlugin) Shutdown() error {
	return nil
}

func (p *PitchlakePlugin) NewBlock(block core.Block, stateUpdate core.StateUpdate, newClasses map[*felt.Felt]core.Class) error {
	return nil
}

func (p *PitchlakePlugin) RevertBlock(block core.Block, stateUpdate core.StateUpdate, newClasses map[*felt.Felt]core.Class) error {
	return nil
}
