# Rule Of Thumb

When you get a bug, you should first do the following steps:
1) Google its bug information;
2) Read its logs;
3) Read the system logs and the kernel log to check if there is OOM/cgroups information;
4) GDB coredump if a coredump file has been generated.

If you still get blocked, classify the bug as below and debug at source code level:


### bohrbug ###
##### Description #####
It is a good, solid bug.
##### Solution #####
print & bisection method


### heisenbug ###
##### Description #####
When you try to debug, it will disappear or phenomenon will change.
##### Solution #####
1) Check what will happen if compiler optimization is enabled/disabled;
2) Check if it is caused by float precision issue;
3) Breakpoint may change the execution time so that some lifetime of variable/resource may be changed, so replace breakpoint with `sleep` and check what will happen.
##### Deeper... #####
1) Print & Bisection method to locate the position of the bug;
2) Analyze which kinds of inputs have a higher probability to reproduce the bug;
3) Read the lists of call stacks to check if there is some race condition issues: deadlock? variable/resource shared by two threads is overwritten? constructor is called once but destructor is called twice? destructor is called too early?


### hindenbug ###
##### Description #####
It is a fatal bug.
##### Solution #####
apologize to your client/your boss/your teammates and debug.


### mandelbug ###
##### Description #####
When you try to debug, more complex bugs appear.
##### Solution #####
Take a coffee...


### schroedinbug ###
##### Description #####
When you haven't read the source code, the program works fine. Whereas when you read the source code, you find a bug. After that, the bug appear immediately while executing the program.
##### Solution #####
Confess to God and debug.


### higgs-bugson ###
##### Description #####
It is a bug that is predicted to exist based upon other observed conditions (most commonly, vaguely related log entries and anecdotal user reports) but is difficult, if not impossible, to artificially reproduce in a development or test environment.
##### Solution #####
FUZZ testing?
