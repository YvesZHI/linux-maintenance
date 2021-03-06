# Limitation of network connection


### OS ###
##### number of file #####
`ulimit -n`: the maximum number of opening files of a process<br>
`cat /proc/sys/fs/file-nr`: the maximum number of filehandles that the Linux kernel will allocate<br>
`cat /proc/sys/net/ipv4/ip_conntrack_max`: the maximum number of ip conntrack of Netfilter<br>

##### number of port #####
0-1024 is reserved by the OS.<br>
`cat /proc/sys/net/ipv4/ip_local_port_range`: the usable ports<br>

##### number of TCP socket #####
Each TCP connection is identified by four elements: {local ip, local port, remote ip, remote port}.<br>
Without counting the special IP addresses, each server can hold 2^48 TCP connections (in the case of local ip and local port fixed).<br>

##### size of TCP/UDP segment #####
`cat /proc/sys/net/ipv4/tcp_rmem && cat /proc/sys/net/ipv4/tcp_wmem`: the size of sliding window (turn off TCP Window scaling will force all TCP connections to use a 64KB window)<br>

##### other costs #####
https://netfilter.org/documentation/FAQ/netfilter-faq-3.html#ss3.6<br>
https://stackoverflow.com/questions/31378403/how-much-data-it-cost-to-set-up-a-tcp-connection<br>

### Internet ###
As some physical reason, the size of each ethernet frame must be 46B-1500B (exclude 18B coming from the head and the tail of Data Link Layer). In other words, 1500B is the maximum size of IP packet (Network Layer). This 1500B is call MTU (Maximum transmission unit).<br>

An ethernet frame is as below:<br>
|&nbsp;Ethernet head (22B)&nbsp;|&nbsp;IP head (20B)&nbsp;|&nbsp;TCP head (20B) / UDP head (8B)&nbsp;|&nbsp;Application Data (1460B for TCP / 1472B for UDP)&nbsp;|<br>

In the case of UDP Fragmentation occurring, UDP datagram has a high probability to lose. Meanwhile, the minimum size of IPv4 datagram of any device has to be able to receive 576B. So it is strongly recommended to set the size of UDP datagram less than 548B (576 - 8(UDP head) - 20(IP head)) while Internet programming so that UDP Fragmentation can be avoided. For LAN programming, the recommended size of UDP datagram could be 1472B.<br>

### References ###
http://www.linuxvox.com/post/what-are-file-max-and-file-nr-linux-kernel-parameters/<br>
https://wiki.khnet.info/index.php/Conntrack_tuning<br>
https://www.cnblogs.com/duanxz/p/4464178.html<br>
