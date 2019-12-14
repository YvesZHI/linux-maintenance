# Limitation of network connection


### limitation of file ###
`ulimit -n`: the maximum number of opening files of a process<br>
`cat /proc/sys/fs/file-nr`: the maximum number of filehandles that the Linux kernel will allocate<br>
`cat /proc/sys/net/ipv4/ip_conntrack_max`: the maximum number of ip conntrack of Netfilter<br>

### limitation of port ###
0-1024 is reserved by the OS.<br>

### TCP socket ###
Each TCP connection is identified by four elements: {local ip, local port, remote ip, remote port}.<br>
Without counting the special IP addresses, each server can hold 2^48 TCP connections (in the case of local ip and local port fixed).<br>

### Cost ###
https://netfilter.org/documentation/FAQ/netfilter-faq-3.html#ss3.6
https://stackoverflow.com/questions/31378403/how-much-data-it-cost-to-set-up-a-tcp-connection

### References ###
http://www.linuxvox.com/post/what-are-file-max-and-file-nr-linux-kernel-parameters/
https://wiki.khnet.info/index.php/Conntrack_tuning
