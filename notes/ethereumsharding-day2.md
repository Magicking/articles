## Cross-shard / cross-contract communication

 - asynchronous call => atomic swap w/ timeout (e.g: timeout of X blocks) ?
 => gas reservation ?
 => still in research stage open to proposition
 gas model TBD

 => Different inter-process/contract models communication
 => TODO Experimentation

 - Ideas:
  => Loss-ok Messaging

## Data availability

Read => https://github.com/ethereum/research/wiki/A-note-on-data-availability-and-erasure-coding
Read => http://www.arijuels.com/wp-content/uploads/2013/09/BJO09b.pdf

Edge cases:
 Fool minority user => Use kind-of routing to avoid user(node) identity disclosure

Proof of custody:
https://distributedlab.com/whitepaper/ProofOfcustody.pdf
https://github.com/vladzamfir/PoC
