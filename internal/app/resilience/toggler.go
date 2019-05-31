package main

func togglerEnable() error {
	stateState.enabled = true
	return nil
}

func togglerDisable() error {
	stateState.enabled = false
	return nil
}
