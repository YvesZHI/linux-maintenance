# DeadLock

### DeadLock Detection ###

#### Method 1: strace ####
`strace -ttT ./a.out` blocks at a line, which contains `FUTEX_WAIT`.

#### Method 2: gdb ####
##### sample code #####
test.cpp:<br>
```
#include <iostream>
#include <thread>
#include <mutex>

std::mutex gMutex;

int Reenter(){
   std::lock_guard<std::mutex> lLock(gMutex);
   return 10;	
}

int Callback()
{
   std::lock_guard<std::mutex> lLock(gMutex); 
   return Reenter();
}
int main(int argc, char**argv) {
    std::cout << "Hello CMake World!" << std::endl;
    std::thread lThread(Callback);
    lThread.join();
    std::cout << "sub thread is gone!" << std::endl;
    return 0;
}
```
##### Compile and run #####
```
g++ test.cpp -ggdb -lpthread -std=c++11 -o test
./test
```
##### Attach it #####
```
sudo gdb -p `pgrep test`
```
##### Find where the deadlock is #####
`bt`
> #0 0x00007f65c7d3e9cd in pthread_join (threadid=140075107923712,<br>
> thread_return=0x0) at pthread_join.c:90<br>
> #1 0x00007f65c7a6cb97 in std::thread::join() ()<br>
> from /usr/lib/x86_64-linux-gnu/libstdc++.so.6<br>
> #2 0x00000000004011ea in main (argc=1, argv=0x7fff7a26af28) at test.cpp:27

`info threads`
> 1 Thread 0x7f65c815a740 (LWP 4094) “test” 0x00007f65c7d3e9cd in pthread_join (threadid=140075107923712, thread_return=0x0) at pthread_join.c:90<br>
> 2 Thread 0x7f65c70cb700 (LWP 4095) “test” __lll_lock_wait ()<br>
> at ../sysdeps/unix/sysv/linux/x86_64/lowlevellock.S:135

`p gMutex`
> $1 = { = {_M_mutex = {__data = {__lock = 2, __count = 0, ***__owner = 4095,*** __nusers = 1, __kind = 0, __spins = 0, __elision = 0, __list = {__prev = 0x0, __next = 0x0}}, __size = "\002\000\000\000\000\000\000\000\377\017\000\000\001", '\000' <repeats 26 times>, __align = 2}}, }

`thread 2`
> [Switching to thread 2 (Thread 0x7f65c70cb700 (LWP 4095))]<br>
> #0 __lll_lock_wait () at ../sysdeps/unix/sysv/linux/x86_64/lowlevellock.S:135<br>
> 135	../sysdeps/unix/sysv/linux/x86_64/lowlevellock.S: No such file or directory.

`bt`
> #0 __lll_lock_wait () at ../sysdeps/unix/sysv/linux/x86_64/lowlevellock.S:135<br>
> #1 0x00007f65c7d3fdfd in GI_pthread_mutex_lock (mutex=0x6052e0 ) at ../nptl/pthread_mutex_lock.c:80 #2 0x0000000000401007 in __gthread_mutex_lock (__mutex=0x6052e0 ) at /usr/include/x86_64-linux-gnu/c++/5/bits/gthr-default.h:748 #3 0x0000000000401498 in std::mutex::lock (this=0x6052e0 ) at /usr/include/c++/5/mutex:135 #4 0x000000000040151e in std::lock_guard::lock_guard ( this=0x7f65c70cae20, __m=...) at /usr/include/c++/5/mutex:386 #5 0x00000000004010ef in Reenter () at test.cpp:11 #6 0x000000000040114b in Callback () at test.cpp:18 #7 0x0000000000402801 in std::_Bind_simple<int (*())()>::_M_invoke<>(std::_Index_tuple<>) (this=0x2085058) at /usr/include/c++/5/functional:1531 #8 0x000000000040275a in std::_Bind_simple<int (*())()>::operator()() ( this=0x2085058) at /usr/include/c++/5/functional:1520 #9 0x00000000004026ea in std::thread::_Impl<std::_Bind_simple<int (*())()> >::_M_run() (this=0x2085040) at /usr/include/c++/5/thread:115 #10 0x00007f65c7a6cc80 in ?? () from /usr/lib/x86_64-linux-gnu/libstdc++.so.6 #11 0x00007f65c7d3d6fa in start_thread (arg=0x7f65c70cb700) at pthread_create.c:333 #12 0x00007f65c74dbb5d in clone () at ../sysdeps/unix/sysv/linux/x86_64/clone.S:109

##### Reference #####
https://ethanhao.github.io/c++11,/gdb,/multithread,/2017/03/03/Deadlock-detecting-using-GDB-Copy.html
