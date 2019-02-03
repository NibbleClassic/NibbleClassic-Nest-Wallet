// Copyright (c) 2018, The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in NBX
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "nibble-service-session.log"
	logWalletdAllSessionsFilename        = "nibble-service.log"
	logNibbledCurrentSessionFilename = "Nibbled-session.log"
	logNibbledAllSessionsFilename    = "Nibbled.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "nibble-service"
	NibbledCommandName               = "Nibbled"
)
