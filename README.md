# NibbleClassic Nest

The universal desktop GUI wallet for NibbleClassic

![Logo](/nibbleclassicnestlogo.png)

## Installation

[Windows](#windows) - [Mac](#mac) - [Linux](#linux)

### Windows

1. Go [here](https://github.com/Seimax/nibbleclassic-wallet-go/releases) and download the latest release called **NibbleClassic-Nest-Winxx-vx.xx.zip**
2. Unzip the folder and launch **NibbleClassic-Nest.exe**. (Make sure you leave everything as is in the folder)

Important notes:

* Make sure *nibble-service.exe* is not running before you start *NibbleClassic-Nest*

### Mac

No Build right now!

### Linux

1. Go [here](https://github.com/Seimax/nibbleclassic-wallet-go/releases) and download the latest release called **NibbleClassic-Nest-Linux-vx.xx.tar.gz**
2. Unzip(tar -xzf NibbleClassic-Nest-Linux-vx.xx.tar.gz) the folder and launch **./nibbleclassic-wallet-go**. (Make sure you leave everything as is in the folder)

## Upgrade

Just download the new release and follow the same steps as [Installation](#installation).
If you are on Windows or Linux, move your wallets (.wallet) and settings.db files from the old Nest folder to the new. Then you can delete the old folder. (on Mac, you do not need to move the settings.db file as it stays in ~/Library/Application Support/NibbleClassic-Nest/).

## Checkpoints

NibbleClassic Nest Supports the loading of a Checkpoints file. This is recommended when using a local node to shorten the waiting time for loading the chain. You find the actual checkpoints.csv at https://github.com/NibbleClassic/Checkpoints

## Screenshots

![Main Screen](/Screenshots/MainScreen.png)

![Open Wallet](/Screenshots/OpenWallet.png)

## Donations

Seimax:
Nib1bW1CQPHJqRQjHEG96t9im5BUN3iQ243kt45NeU8YTKV3fmhTYmaLqqG9p1HQkk769JYR3Y8YCPYNWvoN9nxe5nTXHLW3GH

TurtleCoin Nest Devs:
TRTLv3jzutiQwqHL3qFwsu5EVLWesxZr1AFQ4AuMR3SD56n3rkHDkwj79eKwvaiU1nYQWGydKoXM6fXyiiGKsPDnVCNXzNdusxx

gfx by [thomsane](https://thomsane.de)

## Build - (for developers only)

### Linux

1. Download Go from [here](https://golang.org/dl/)

2. Use `tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz` to extract the downloaded go package.

3. Add the following lines to `.bashrc` file, save the file and then execute the command `source .bashrc` in a terminal.
    ```
    export GOPATH=$HOME/go

    export GOBIN=$GOPATH/bin

    export GOROOT=/usr/local/go

    export PATH=$HOME/bin:$HOME/.local/bin:$PATH:$GOROOT/bin:$GOBIN
    ```
4. Similarly add the following lines to `.profile` file, save the file and then execute the command `source .profile` in a terminal.
    ```
    CGO_CXXFLAGS_ALLOW=".*" 
    CGO_LDFLAGS_ALLOW=".*" 
    CGO_CFLAGS_ALLOW=".*" 
    ```
5. Follow the instructions present [here](https://github.com/therecipe/qt/wiki/Installation-on-Linux) till **Run the setup** to install Qt which is the most important binding required to build Nest.
6. Type the following commands to clone the Nest wallet, install dependencies and build the wallet.
    ```
    $ cd $HOME/go/src
    $ git clone https://github.com/NibbleClassic/turtle-wallet-go.git NibbleClassic-Nest
    $ go get -v github.com/atotto/clipboard github.com/dustin/go-humanize github.com/mattn/go-sqlite3 github.com/mcuadros/go-version github.com/mitchellh/go-ps github.com/pkg/errors
    $ cd NibbleClassic-Nest
    $ qtdeploy build desktop
    ```

1. The app folder is in deploy/linux/
1. Include the latest _nibble-service_ and _Nibbled_ builds in the app folder

### Windows - Mac

1. Install Go (https://golang.org/doc/install)

1. Install this binding: https://github.com/therecipe/qt (installation instructions at https://github.com/therecipe/qt/wiki/Installation)

1. Insall Go libraries (in console or terminal):
    ```
    $ go get github.com/atotto/clipboard github.com/dustin/go-humanize github.com/mattn/go-sqlite3 github.com/mcuadros/go-version github.com/mitchellh/go-ps github.com/pkg/errors
    ```

1. Run `qtdeploy build desktop`

1. The app folder is in deploy/*your os*/

1. Include the latest _nibble-service_ and _Nibbled_ builds in:
    * Windows: in the app folder
    * Mac: in NibbleClassic-Nest.app/Contents/
