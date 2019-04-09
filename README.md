# aether

github.com/IMTTHOLDINGCORP/go-aether is an [Ethereum-compatible](https://github.com/ethereum/go-ethereum) project.

## Warning

We suggest that the GasPrice should not be less than 18Gwei, otherwise the transaction may not be packaged into the block.

## List of Chain ID's:
| Chain(s)    |  CHAIN_ID  | 
| ----------  | :-----------:| 
| testnet     | 1792            | 
| mainnet     | 1913     | 

## Build the source 

Building github.com/IMTTHOLDINGCORP/go-aether requires both a Go (version 1.9.2 or later) and a C compiler. You can install them using your favourite package manager.

### Ubuntu:

	1）Install git
		$ apt-get install git
		$ apt-get install golang
	2）Install golang in the directory of {GOHOME}
		$ cd {GOHOME}
		$ tar -zxvf go1.9.2.linux-amd64.tar
		$ mkdir {GOHOME}/gopath
		$ export GOAPTH={GOHOME}/gopath 
		$ export GOROOT=${GOHOME}/go 
		$ export PATH=$GOROOT/bin:$PATH
	3）download the source on the git
		$ cd {GOPATH}
		$ mkdirs src\github.com\IMTTHOLDINGCORP\
		$ cd src\github.com\IMTTHOLDINGCORP\
		$ git clone https://github.com/IMTTHOLDINGCORP/go-aether.git
	4）compile the source code
		$ cd github.com/IMTTHOLDINGCORP/go-aether
		$ make aether
	5）run the program
		$ ./build/bin/aether console

### Mac


	1）download the latest *github.com/IMTTHOLDINGCORP/go-aether* source code from the github
	
        	$ git clone https://github.com/IMTTHOLDINGCORP/go-aether.git

	2）Install by using source code

        	$ cd github.com/IMTTHOLDINGCORP/go-aether
        	$ make aether

		If any error happens during the compiling and prompt message shows " lacking of the header file that is related to the Mac OS", you may try to install the xcode command-line tool before executing above commands:
	
	        $ xcode-select --install

	3） startup
	
        	$ build/bin/aether console

### Windows

	1）First of all, you need to install a package management software for Windows named "chocolatey". Please refer to https://chocolatey.org for Installation method.
	2）Install git, golang, mingw by using chocolatey
		c:\Users\xxx> choco install git
		c:\Users\xxx> cholo install golang 
		c:\Users\xxx> cholo install mingw

	3）Set environment variables {GOPATH} of golang：
		c:\Users\xxx> mkdir {GOPATH}
		c:\Users\xxx> set "GOPATH={GOPATH}"
		c:\Users\xxx> set "Path={GOPATH}\bin;%Path%"
	4）download the source on the git： 
		c:\Users\xxx> cd {GOPATH}
		{GOPATH}> mkdir src\github.com\IMTTHOLDINGCORP\
		{GOPATH}> cd src\github.com\IMTTHOLDINGCORP\
		...\IMTTHOLDINGCORP> git clone https://github.com/IMTTHOLDINGCORP/go-aether.git
	5）compile source code：
		...\IMTTHOLDINGCORP> cd github.com/IMTTHOLDINGCORP/go-aether
		...\IMTTHOLDINGCORP> go install -v ./cmd/...
	6）run the program：
		c:\Users\xxx> aether console



## Encrypt your nodekey

     $ ./build/bin/aether security --passwd
## Decrypt your nodekey

     $ ./build/bin/aether security --unlock
     
## Run fast node to test aether

    $ ./build/bin/aether console
    
## Create new account

    > personal.newAccount()

## View the miner nodes

* tribe.getStatus()
* then you will see the following message:
<pre><code>{
  "epoch": 5760,
  "number": 39601,
  "signerLevel": "None",
  "signerLimit": 17,
  "signers": [
    {
      "address": "0x3a5fbac6ca913599c5fde8c1638db58d01de8a48",
      "score": 3
    },
    {
      "address": "0xad4c80164065a3c33dd2014908c7563eff88ab49",
      "score": 3
    },
    {
      "address": "0xc22d53456abd14da347517a4b47ea24866b8e3ae",
      "score": 3
    },
    {
      "address": "0x7b06dd132c089034157f1e1aacda75787df1e0c5",
      "score": 3
    },
    {
      "address": "0x00ab501f3fe4b2f71651764699ec5752598e679f",
      "score": 3
    }
  ],
  "totalSinner": 2,
  "volunteerLimit": 70,
  "volunteers": []
}</code></pre>

that tell you there are two miners in the testnet.

## View the block mining history in console

* tribe.getHistory(4)
    
* then you will see the following message (console format):
<pre><code>[{
    difficulty: 3,
    hash: "0xb54d1b36e324bfd0c0747edf93430a1b6c62d0a058fd7659eb97f51921e1f347",
    number: 2198473,
    signer: "0xf848f385fd21c6972264c777684940814a7d4792",
    timestamp: 1554792114
}, {
    difficulty: 3,
    hash: "0x27f7fff0011929c879b25b4fb1529bad0d5ef50ec66f4cfcd454ccfaebb311eb",
    number: 2198472,
    signer: "0x4110bd1ff0b73fa12c259acf39c950277f266787",
    timestamp: 1554792100
}, {
    difficulty: 3,
    hash: "0x6bfd6aa2c5dbf67529058d600b375c88f24b53c2631e31ae754e24607b18dfdf",
    number: 2198471,
    signer: "0xad9581fe7f9b640cc34915cd988965216e44a972",
    timestamp: 1554792086
}, {
    difficulty: 3,
    hash: "0x4fa3c0a910bf9378de8ab473ed27c955cb763c1315cbbf344c7ceb792046e029",
    number: 2198470,
    signer: "0xf848f385fd21c6972264c777684940814a7d4792",
    timestamp: 1554792072
}]</code></pre>

that tell the block number and miner's account that generate that block.
