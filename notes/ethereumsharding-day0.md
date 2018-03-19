# Sharding

Blurry idea

# Introduction

### Phase 1
 - No state execution
 - Because there is no execution how Proposer can get paid ?
 - Incentive Execution Engine (Out of blockchain incentive)
 - Transaction
 - /!\ Multiple same transaction into different blobs /!\ // processing ? => See state execution
 - Phase ~0 dumping data (anykind of data) into blob for phase ~1
 - Private tx pool before becoming blobs
 - Collators choose the shard depending if they(shard) are available(witness? data available to download) to build on
 - Zone or galaxy bunch of shards

### Phase 2 State execution
 - (Dapp / User) Batching TX to proposer => Phase 1 ?

### Phase 4 Cross-shard TX

### Phase 5
 - Shard manager (Collate/Propose/Execute)
 - -> Get notify w/ a listener to SMC
 - -> Communication between shards & SMC
 - -> new p2p subprotocol

### Blurry phase
 - Head choosing based on availability
 - Proposer should know the state to maximize profitability
 - Proposer also executors
 - Archival node => Proposer (because they need to know the state) ?
 - Zone clusters of contract with synchronous internal TX and asynchronous TX w/ contract outside of zone (w/ receipt ?)
 - 
 - VM for shard may not be the same EVM as the main shard VM

# Shard Manager Contract
 - Proposer
 - Collator
 - Executor (Phase 3?)

Basically what's in https://ethresear.ch/t/sharding-phase-1-spec/1407

# Proposer / Collator Separation, Ghost and JMRS
Proposer can challenge collator that included their proposal, or up to 25th-degree descendants
The transaction ID is w/o Collator sig
Challenge = index in collation Merkle tree
Response = merkle branch of that piece of data
If responder can't answer, deposit transferred to proposer
/!\ DOS by asking too many reponsder /!\

=> Encourage collator to pre-co-signs all available proposals
Could be made Out Of Band /!\
 1. Proposers create proposals, reveal headers only
 2. Collator pre-co-signs all available propoposals
Cryptoeconomically commits to not co-signing anything outside the list or expanding the list.
(Order (somehow) list of signed proposals)
Proposer of winning proposal reveals body Collator so-signs

Side note: Broadcast relay scheme (relay top % price for example)
Demystifying Incentives in the Consensus Computer => https://eprint.iacr.org/2015/702.pdf

# Execution

## Ideas

No cross shard communication
Clients download chain & compute state

### V1
Execution Manager Contract
 1. Executor register w/ deposit
 2. Executor claims a state that can be challenged if incorrect (slashing)

### V2
Execution Manager Contract
Truebit kind-of verification mecanism
If substantial disagreement exist, localy calculate state [collation_hash+prev_state_root] => new state root
State roots are basically claims calculated either by light clients or executors
state is tree

Account abstraction (w/ PAYGAS?) is basically removing execution from the consensus, it rely on executors (which can be light client)
Proposer should make claim when they propose a collation

# Rent or self pruning scheme

https://projectchicago.io
https://gastoken.io
https://eprint.iacr.org/2018/078.pdf
