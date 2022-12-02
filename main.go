import (
	"fmt"

	cfg "github.com/tendermint/tendermint/config"
	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/store"
)

func main() {
	bs, err = loadBlockStore(config)				// TODO: create config
	if err != nil {
		return -1, nil, err
	}
	defer func() {
		_ = bs.Close()
	}()
	maxHeight := bs.Height()
	for height := maxHeight;;i-- {
		block := bs.LoadBlock(height)
		if block == nil {
			return -1, nil, fmt.Errorf("block at height %d not found", height)
		}
		// Parse Block Parameters
		fmt.Printf("%+v\n", block)
	}
}

func loadBlockStore(config *cgf.BaseConfig) (*store.BlockStore, error) {
	dbType := dbm.BackendType(config.DBBackend)

	if !os.FileExists(filepath.Join(config.DBDir(), "blockstore.db")) {
		return nil, nil, fmt.Errorf("no blockstore found in %v", config.DBDir())
	}

	// Get BlockStore
	blockStoreDB, err := dbm.NewDB("blockstore", dbType, config.DBDir())
	if err != nil {
		return nil, nil, err
	}
	blockStore := store.NewBlockStore(blockStoreDB)

	return blockStore, nil
}

func loadStateStore(config *cgf.BaseConfig) (*store.StateStore, error) {
	dbType := dbm.BackendType(config.DBBackend)

	if !os.FileExists(filepath.Join(config.DBDir(), "statestore.db")) {
		return nil, nil, fmt.Errorf("no statestore found in %v", config.DBDir())
	}

	// Get BlockStore
	stateStoreDB, err := dbm.NewDB("statestore", dbType, config.DBDir())
	if err != nil {
		return nil, nil, err
	}
	stateStore := store.NewStateStore(stateStoreDB)

	return stateStore, nil
}