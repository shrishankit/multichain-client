# multichain-client
Golang client library for MultiChain blockchain

# Getting Started

Once you cloned the repo create a go module (go.mod) or directly in from git

- Create New Client
    mcObj = multichain.NewClient("Name_of_chain", 
                                "172.17.0.1", // address <br/>
                                "6482",       // port <br/>
                                "multichainrpc", // rpcuser <br/>
                                "JDmEbay3zQuE5XA2HpY9sPXmS2XsuFKGgCr8wZ9SzJQh" // rpcpassword <br/>
                                )
    <b>rpc user and password will present in the multichain config file</b>
