# How does Kernel manage the filesystem
### Hard drive
There is a INode Table on the hard drive (partition), meaning that each filesystem has its own INode Table. This table contains the inodes of all of directories and files. Each inode in this table has a unique number, which is called INode Number.

An inode contains two sorts of data: metadata and a pointer.
- metadada: such as User, Group, Size, Ctime, Atime, Mtime, Permission etc
- pointer: points to the content which is located on the disk blocks
For a file, the content is simply its data. For a directory, the content is a directory structure, which contains the INode Number, filename and the length of the filename of the file in the directory.

### Kernel Space
Kernel maintains two things: INode Table and Open File Table.
The INode Table here comes from the INode Table on the hard drive, but a new inode will be generated only if RAM needs to access a file. If an inode is not needed anymore, it will be free.

Open File Table contains the metadata about opened file by processes. Each entry of Open File Table has the following format:
- INode Number
- Open Instance Count: the number of processes sharing the open instance of the file represented by this open file table entry
- Lseek: the offset position from which the next word is read from or written into in the file

`open` will increase the Open Instance Count, `close` will do the reverse.
`fork` and `dup` won't create a new entry of Open File Table but will increase the Open Instance Count.
When the Open Instance Count equals to 0, the resource of the entry of Open File Table can be free.

### User Space
Each process has its own file descriptor table. When a process `open` a file, a new file descriptor will be generated, which points to some entry in Open File Table.
`fork` and `dup` will generate a new file descriptor, old and new file descriptor points to the same entry of Open File Table.


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


# References
https://www.programmersought.com/article/1218984627/
