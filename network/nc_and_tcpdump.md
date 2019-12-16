# network tools
server: 192.168.1.100:1234<br>
client: 192.168.1.200<br>

### nc ###
##### scan port #####
TCP scanner: `nc -v -z -w2 192.168.1.100 1234`<br>
UDP scanner: `nc -v -z -w2 192.168.1.100 1234`<br>
UDP scanner can't fail.<br>

##### listen #####
TCP listener: `nc -l 1234`<br>
UDP listener: `nc -u -l 1234`<br>

##### send #####
TCP sender: `nc 192.168.1.100 1234`<br>
UDP sender: `nc -u 192.168.1.100 1234`<br>


### tcpdump ###
server: ` tcpdump -i eth0 port 1234`<br>

##### simulate TCP connection #####
server: `nc -l 1234`<br>
client: `nc 192.168.1.100 1234`<br>

output of `tcpdump`:<br>
`three handshake: `12:22:36.551563 IP 192.168.1.200.51516 > 192.168.1.100.1234: Flags [S], seq 2524926410, win 29200, options [mss 1460,sackOK,TS val 2061376385 ecr 0,nop,wscale 7], length 0<br>
`three handshake: `12:22:36.551575 IP 192.168.1.100.1234 > 192.168.1.200.51516: Flags [S.], seq 3399548678, ack 2524926411, win 28960, options [mss 1460,sackOK,TS val 1582320715 ecr 2061376385,nop,wscale 7], length 0<br>
`three handshake: `12:22:36.551639 IP 192.168.1.200.51516 > 192.168.1.100.1234: Flags [.], ack 1, win 229, options [nop,nop,TS val 2061376385 ecr 1582320715], length 0<br>
`sending: `12:22:36.551692 IP 192.168.1.200.51516 > 192.168.1.100.1234: Flags [P.], seq 1:7, ack 1, win 229, options [nop,nop,TS val 2061376385 ecr 1582320715], length 6<br>
`ack: `12:22:36.551699 IP 192.168.1.100.1234 > 192.168.1.200.51516: Flags [.], ack 7, win 227, options [nop,nop,TS val 1582320715 ecr 2061376385], length 0<br>
`four handshake: `12:22:36.551701 IP 192.168.1.200.51516 > 192.168.1.100.1234: Flags [F.], seq 7, ack 1, win 229, options [nop,nop,TS val 2061376385 ecr 1582320715], length 0<br>
`four handshake: `12:22:36.551716 IP 192.168.1.100.1234 > 192.168.1.200.51516: Flags [F.], seq 1, ack 8, win 227, options [nop,nop,TS val 1582320715 ecr 2061376385], length 0<br>
`four handshake: `12:22:36.551780 IP 192.168.1.200.51516 > 192.168.1.100.1234: Flags [.], ack 2, win 229, options [nop,nop,TS val 2061376385 ecr 1582320715], length 0<br>

Why does the last `four handshake` miss?<br>


##### simulate UDP connection #####
server: `nc -u -l 1234`<br>
client: `echo abcde | nc -u 192.168.1.100 1234`<br>

output of `tcpdump`:<br>
12:20:46.833832 IP 192.168.1.200.38490 > 192.168.1.100.1234: UDP, length 6
