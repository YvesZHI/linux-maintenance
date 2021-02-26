# How does Kernel manage the filesystem
### Hard drive
There is a INode Table on the hard drive (partition), meaning that each filesystem has its own INode Table. This table contains the inodes of all of directories and files. Each inode in this table has a unique number, which is called INode Number.

An inode contains two sorts of data: metadata and a pointer.
- metadada: such as User, Group, Size, Ctime, Atime, Mtime, Permission etc
- pointer: points to the content which is located on the disk blocks
For a file, the content is simply its data. For a directory, the content is a directory structure, which contains the INode Number, filename and the length of the filename of the file in the directory.

### Kernel Space
Kernel maintains two things: INode Table and Open File Table.
The INode Table here comes from the INode Table on the hard drive (when booting up, the kernel will generate its INode Table by reading the INode Table on the hard drive), but a new inode will be generated only if RAM needs to access a file. If an inode is not needed anymore, it will be free.

Open File Table contains the metadata about opened file by processes. Each entry of Open File Table has the following format:
- INode Number
- Open Instance Count: the number of processes sharing the open instance of the file represented by this open file table entry
- Lseek: the offset position from which the next word is read from or written into in the file

`open` will increase the Open Instance Count, `close` will do the reverse.
`fork` and `dup` won't create a new entry of Open File Table but will increase the Open Instance Count.
When the Open Instance Count equals to 0, the resource of the entry of Open File Table can be free.

### User Space
Each process has its own file descriptor table. When a process `open` a file, a new file descriptor will be generated, which points to some entry in Open File Table.
`fork` and `dup` will generate a new file descriptor, old and new file descriptors point to the same entry of Open File Table.

### Hardlink vs Softlink
Hardlink: different entries in the INode Table containing the same Inode ID.<br>
Softlink: an entry in the INode Table, pointing a special file, whose data is a pointer, which points an entry in the INode Table of the real file.

### Graph show
```
  ProcA                   Open File             INode
|-------|                   Table               Table
|  fd1  |--------\       |--------|           |--------|
|-------|   dup   \----> | Cnt: 3 |---------> | Cnt: 1 |
|  fd2  |--------/ /     | Offset |           |filename|
|-------|         /      |--------|           |Inode ID|
                 /       | Cnt: 1 |           |--------|=========> Hard Drive
  ProcB    fork /   /--> | Offset |---------> | Cnt: 1 |
|-------|      /   /     |--------|           |filename|
|  fd1  |-----/   /                           |Inode ID|
|-------|        /                            |--------|
|  fd2  |-------/
|-------|

/proc/<PID>/fd/             lsof             ls, stat, debugfs...
```



# How does a file (data) is read and written
### Page Cache / Buffer
The Page Cache is located in RAM of Kernel space, which is a buffer between process (User space) and hard drive. It is used to avoid performing directory IO from CPU to hard drive.
There are two issues:
1. Dirty data. The data in the page cache and in the hard drive may be different.
2. Duplicated data. If the process is holding some data coming from the file, it means that the same data can be found in the process and in the buffer at the same time.

### Direct IO
Read and write data from process to hard drive directory, meaning that there is no Page Cache anymore.

### Standard IO
CPU copies data from hard drive to Page Cache, then CPU copies data from Page Cache to process. In a word, Page Cache is located between process and hard drive so data must be copied twice.

### mmap
CPU copies data from hard drive to Page Cache, then mmap maps the memory of Page Cache to process. In other words, mmap makes Page Cache and process share the same memory so that process can read and write immediately at User space. Comparing with the Standard IO, mmap reduces one data copy.

### sendfile
CPU copies data between one file descriptor and another. It happens in Kernel space so that it saves data copies between Page Cache to processes.

### splice
Splice is just like how sendfile works but there are two differences:
1. Splice can only be used to copy data from file to socket buffer (in Kernel space);
2. Sendfile is still a CPU copy but splice creates a pipeline between file (in Page Cache) and socket buffer to transfer data.

# References
https://www.programmersought.com/article/1218984627<br>
https://zhuanlan.zhihu.com/p/83398714
