# lock


### mechanism ###
The lock relies on two kinds of operations: `test-and-set` and memory fence. Both of them come from the CPU instruction, so they won't involve the kernel. Besides, some other hardware-level algorithms are relied too, such as Cache coherence, which won't involve the kernel either.

The lock can be regarded as an integer: 0 means unlock and 1 means lock. `test-and-set` makes sure that switching between 0 and 1 is atomic, and memory fence makes sure that the protected region by the lock won't be reordered outside of the region (without memory fence, compiler and CPU may do some optimisation which may reorder the execution order).

Normally, the atomic operation, such as `test-and-set`, only works for the fundamental types, like an integer, byte, long and double (depending on architecture), but with the help of some languages, such as C++11, the atomic operation on a struct (`std::atomic<CustomStruct>` in C++11) is possible too (depending on architecture).

### lock and kernel ###
Till now, the lock doesn't involve the kernel, so what is the role of the kernel in this issue?

The kernel provides a kind of lock, it is called mutex. The mutex is generated based on the lock above (this is why it is called low-level lock) and which can do two things:<br>
1) It can lock not noly a data, but also a piece of code;<br>
2) It allows the kernel to schedule the threads or the processes. For example, if a process or a thread needs to be put to sleep (to wait to acquire the lock) or woken (because it couldn't acquire the lock but now can), the kernel has to be involved to perform the scheduling operations.

In a word, the mutex wraps the atomic operation and memory fence together so that it can be used easily to lock a data, a struct and a piece of code, and it can also make the process or the thread be scheduled by the kernel.

One more thing: lock-free. What is lock-free?

Lock-free means "multithread but do NOT let the kernel schedules".

### futex ###
The futex is just this kind of mutex, coming from the Linux kernel. Furthermore, it does some optimization.

For the older-version mutex, if a process or a thread is trying to acquire the lock but fails, the process or the thread will be scheduled by the kernel immediately, meaning that it will be put to sleep. This cause two kinds of latency: 1) the context switch from the user space to the kernel space; 2) the kernel scheduling operation: the state switch between asleep and awake.

The rule of the futex is if a process or a thread fails on acquiring the lock, it will not involve the kernel but keep trying to acquire the lock in a short time. If the acquisition still fails, the futex will involve the kernel, meaning that the futex lets the kernel schedule the process or the thread.

In short, futex is a hybrid mutex (spinlock + custom mutex). In the period of spinlock, if the acquisition succeeds, the context switch and the state switch can be avoided.

### spinlock ###
The spinlock is implemented based on the atomic operation `test-and-set`. With the spinlock, if a process or a thread fails on the acquisition of lock, it will keep trying until it acquires the lock. This is called busy-waiting.

Today, the spinlock may (depending on architecture) become hybird too (spinlock + CPU instruction (such as HALT)). If a process or a thread keeps failing on the acquisition of lock for some time, the spinlock will decide to stop the process or the thread with some CPU instruction for a while and re-execute it later, instead of involving the kernel (because using spinlock means involving the kernel is not expected).

### performance ###
When to use lock, mutex/futex and spinlock? Which one has a better performance?<br>
The answer is: it depends.<br>
These things are implemented based on the architecture, so a different architecture may bring a different performance.

There are several rules of thumb:
1) When doubt, use mutex/futex;<br>
2) If the lock is needed because of some business logic, it normally means a piece of code needs to be locked. In this case, use mutex/futex;<br>
3) If a fundamental type needs to be locked for a short time, use spinlock.

In the real world, different architecture, different business logic and different design mixed up together. So if it is hard to make a convincing prediction on their performance by brain, testing is the only right way.


# volatile


### mechanism ###
`volatile` doesn't create any memory fence, doesn't make operations atomic, so it is orthogonal with thread safe, meaning that it can't be used as any kind of lock. It simply tells compiler and CPU that the variable specified by `volatile` may be changed from outside so do not do any optimisation on the variable, and the value of the variable cached by L1, L2 or L3 is not unreliable so accessing the variable must come from or go to the RAM.

`volatile` is normally used in three cases:
1) MMIO;
2) signal handlers;
3) `setjmp`, `longjmp` and `getjmp` sequences.

For example, 
```
int i = 1;
while (i == 1) {}
```
may be optimized as `while (true) {}`, but `volatile int i = 1;` will forbid this optimization.
