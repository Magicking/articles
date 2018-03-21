## eWASM

More users, more devs
For Go, see wagon (https://github.com/go-interpreter/wagon)
https://github.com/ewasm/design

Sentinel contract for deployment
Processes all contract
Benchmarks needed

Introducing complexity with 32 & 64 underlying architecture and the underlying infrastructure
https://github.com/ewasm/design/blob/master/determining_wasm_gas_costs.md

EVM-C seems to have a lot of cool features

## Execution & State minimization

Separation of data availability (bandwith), execution (computation) & state (storage & I/O)

To have a decoupling of witness data, use a non-dynamic accumulator

Definition availabilty:
 transaction: present in the gossip network
 data/state: historical state ?

Should incentivize the historical availability of blobs
Standardized receipt to cross-application communication (control ERC20 contracts w/ Multisig)

## Account abstraction (main chain)

 - Account abstraction for sharding is different /!\
 - Address re-use checking mecanism sort-of
 - Need standard/basic code

Problem w/ panic on certain conditions where gas usage is hidden to full nodes and only known to miner

Proposer might use a templates of opcodes to accept TX


## Stateless client witchcraft
witness: set of merkle branch you want to access
block producer => stateful
client => stateless
everyone can be stateless except users need to maintain Merkle branches for theirs accounts

## p2p Networking

See photos for notes
