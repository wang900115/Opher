
# small data / index categories:
    block's header, tx's index, account's state

# big data / history categories:
    block's payload, NFT metadata

# immutable / verification site:
    merkle root, state tree 


```
        +---------------------+
        |   P2P Network       |
        +---------------------+
                   |
       +-----------------------+
       |   Consensus Layer     |
       | (PoW / PBFT / Raft)   |
       +-----------------------+
                   |
      -----------------------------
      |            |               |
+---------+  +-----------+   +-----------+
| LevelDB |  | Merkle    |   | IPFS /    |
| (State  |  | Trie /    |   | Arweave   |
| / Index)|  | Root)     |   | (Large    |
|         |  |           |   | Payload)  |
+---------+  +-----------+   +-----------+
```
