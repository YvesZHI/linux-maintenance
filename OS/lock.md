# lock


### How does lock work ###
The lock relies on two kinds of operations: `test-and-set` and memory fence. Both of them come from the CPU instruction, so they won't involve the kernel. Besides, some other hardware-level algorithms are relied too, such as Cache coherence, which won't involve the kernel either.

The lock can be regarded as an integer: 0 means unlock and 1 means lock. `test-and-set` makes sure that switching between 0 and 1 is atomic, and memory fence makes sure that the protected region by the lock won't be reloaded outside of the region.

Normally, the atomic operation, such as `test-and-set`, only works for the fundamental types, like an integer, byte, long and double (depending on architecture), but with the help of some languages, such as C++11, the atomic operation on a struct (`std::atomic<CustomStruct>` in C++11) is possible too (depending on architecture).

### lock and kernel ###
Till now, the lock doesn't involve the kernel, so what is the role of the kernel in this issue?

The kernel provides a kind of lock, it is called mutex. The mutex is generated based on the lock above (this is why it is called low-level lock) and which can do two things:<br>
1) It can lock not noly a data, but also a piece of code;<br>
2) It allows the kernel to schedule the threads or the processes. For example, if a process or a thread needs to be put to sleep (to wait to acquire the lock) or woken (because it couldn't acquire the lock but now can), the kernel has to be involved to perform the scheduling operations.

In a word, the mutex wraps the atomic operation and memory fence together so that it can be used easily to lock a data, a struct and a piece of code, and it can also make the process or the thread be scheduled by the kernel.

### futex ###
The futex is just this kind of mutex, coming from the Linux kernel. Furthermore, it does some optimization.

For the older-version mutex, if a process or a thread is trying to acquire the lock but fails, the process or the thread will be scheduled by the kernel immediately, meaning that it will be put to sleep. This cause two kinds of latency: 1) the context switch from the user space to the kernel space; 2) the state switch between asleep and awake.

The rule of the futex is if a process or a thread fails on acquiring the lock, it will not involve the kernel but keep trying to acquire the lock in a short time. If the acquisition still fails, the futex will involve the kernel, meaning that the futex lets the kernel schedule the process or the thread.

In short, futex = spinlock + mutex. In the period of spinlock, if the acquisition succeeds, the context switch and the state switch can be avoided.
