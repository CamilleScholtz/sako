package main

func daemonHeight() (int, error) {
	v, err := daemon.GetBlockCount()
	if err != nil {
		return 0, err
	}

	return v.Count, nil
}
