# All kinds of cache mechanism

https://www.geeksforgeeks.org/whats-difference-between-cpu-cache-and-tlb

1. First go to the cache memory and if it's a cache hit, then we are done.

2. If it's a cache miss, go to step 3.

3. First go to TLB and if it's a TLB hit, go to physical memory using physical address formed, we are done.

4. If it's a TLB miss, then go to page table to get the frame number of your page for forming the physical address.

5. If the page is not found, it's a page fault. Use one of the page replacement algorithms if all the frames are occupied by some page else just load the required page from secondary memory to physical memory frame.

