package streamreader

import "poe.market/config"

func LaunchReader(config config.ConfigFile) {
	pageChan := make(chan StashPage)
	defer close(pageChan)

	go readRoutine(pageChan)
	writeToDbRoutine(pageChan, config)
}
