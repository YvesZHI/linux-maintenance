# Context Switching
### definition
Context is a stuff on CPU, which needs to be saved so that CPU can restart the execution at the current point at some later time (usually after an interrupt).

Switching means switching between runnable tasks (processes, threads etc). It removes the running task in CPU and replaces it with another runnable task.

### steps
1. Enters the privileged state (内核态);
2. Copies the register values to a safe place (stack on RAM);
3. Loads the registers for new context;
4. Re-enters the user state (用户态).

### Process Context Switching vs Thread Context Switching
When a PCS happens, the CPU state of one process will be saved into PCB (Process control block, a data structure used by OS to store all the information about a process), and the CPU state of the other process will be restored from its PCB.

Furthermore, the virtual memory space will become invalid so that TLB has to be flushed, which spends lots of time. But this won't happen while TCS occuring. For TCS, it's only necessary to store the registers and the stack information.
