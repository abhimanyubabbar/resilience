package main

func togglerEnable() error {
	stateState.hostsHash = [32]byte{}
	err := updateHosts(false)
	if err == nil {
		stateState.enabled = true
	}
	return err
}

func togglerDisable() error {
	err := denierUpdate([]byte("# Block list currently cleared."))
	if err == nil {
		stateState.enabled = false
	}
	return err
}
