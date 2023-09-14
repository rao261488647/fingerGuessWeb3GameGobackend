package main

import "fingerGuessWeb3GameGoBackend/config"

func init() {
	config.SetupConfig()
	config.SetupDbEngine()
	config.SetupTable()
	config.SetupEthClient()
}

func main() {

}
