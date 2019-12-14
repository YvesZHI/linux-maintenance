# Limitation of network connection


### number of file ###
`ulimit -n`: the maximum number of opening files of a process<br>
`cat /proc/sys/fs/file-nr`: the maximum number of filehandles that the Linux kernel will allocate<br>
`cat /proc/sys/net/ipv4/ip_conntrack_max`: the maximum number of ip conntrack of Netfilter<br>

### number of port ###
0-1024 is reserved by the OS.<br>

### number of TCP socket ###
Each TCP connection is identified by four elements: {local ip, local port, remote ip, remote port}.<br>
Without counting the special IP addresses, each server can hold 2^48 TCP connections (in the case of local ip and local port fixed).<br>

### size of packet ###
As some physical reason, the size of each ethernet frame must be 46B-1500B. An ethernet frame is as below:<br>
|&nbsp;Ethernet head (22B)&nbsp;|&nbsp;IP head (20B)&nbsp;|&nbsp;TCP head (20B)&nbsp;|&nbsp;Application Data (1400B)&nbsp;|<br>
`cat /proc/sys/net/ipv4/tcp_rmem && cat /proc/sys/net/ipv4/tcp_wmem`: the size of sliding window (turn off TCP Window scaling will force all TCP connections to use a 64KB window)<br>


### other costs ###
https://netfilter.org/documentation/FAQ/netfilter-faq-3.html#ss3.6<br>
https://stackoverflow.com/questions/31378403/how-much-data-it-cost-to-set-up-a-tcp-connection<br>

### References ###
http://www.linuxvox.com/post/what-are-file-max-and-file-nr-linux-kernel-parameters/<br>
https://wiki.khnet.info/index.php/Conntrack_tuning<br>
