package main

func walletBalance() (float64, float64, error) {
	v, err := wallet.GetBalance()
	if err != nil {
		return 0, 0, err
	}

	// TODO: Make this always use the same amount of decimals.
	return float64(v.Balance) / 1.e+12, float64(v.UnBalance) / 1.e+12, nil
}

func walletAddress() (string, error) {
	v, err := wallet.GetAddress()
	if err != nil {
		return "", err
	}

	return v.Address, nil
}

func walletHeight() (int64, error) {
	v, err := wallet.GetHeight()
	if err != nil {
		return 0, err
	}

	return v.Height, nil
}
