# mutex


### How does mutex work ###
The mutex relies on two kinds of operations: `test-and-set` and memory fence. Both of them come from the CPU instruction, so they won't involve the kernel. Besides, some other hardware-level algorithms are relied too, such as Cache coherence.

The mutex can be regarded as an integer: 0 means unlock and 1 means lock. `test-and-set` makes sure that switching between 0 and 1 is atomic, and memory fence makes sure that the protected region by the mutex won't be reloaded outside of the region.

### futex ###
