# OS tools


### Monitor mem usage of a process ###
```
top -d 1 -b | grep a.out
```

### Get the resource usage of a process ###
```
/usr/bin/time -v /absolutePathOfExe/a.out
```

### Sort mem usage by VIRT ###
```
ps -e -o pid,vsz,comm= | sort -n -k 2
```

### Calculate all CPUs usages ###
```
res=`ps -e -o %cpu | sed '1d'`
num=0.0
for i in $res; do
    num=`awk "BEGIN {print $num+$i}"`
done
echo $num
```

### Find out what is using your swap ###
```
SUM=0
OVERALL=0
for DIR in `find /proc/ -maxdepth 1 -type d | egrep "^/proc/[0-9]"` ; do
PID=`echo $DIR | cut -d / -f 3`
PROGNAME=`ps -p $PID -o comm --no-headers`
for SWAP in `grep Swap $DIR/smaps 2>/dev/null| awk '{ print $2 }'`
do
let SUM=$SUM+$SWAP
done
echo "PID=$PID - Swap used: $SUM - ($PROGNAME )"
let OVERALL=$OVERALL+$SUM
SUM=0
done
echo "Overall swap used: $OVERALL"
```

### Where all your disk space goes ###
```
sudo du -hsx /* | sort -rh
sudo du -h `find / -type f -size +500M`
```

### Arithmetic ###
```
awk "BEGIN {print 2.0/3}"
```

### find ###
##### find all executable binary files #####
```
find -type f -executable -exec file -i '{}' \; | grep 'x-executable; charset=binary'
```

### grep ###
##### grep multi lines #####
```
grep -RlZ lineA | xargs -0 grep -l lineB
```

### sed ###
##### insert a file content into another file after a specific line number #####
```
sed -e "${line}r FILE1" FILE2
```

##### replace the nth line of a file #####
```
sed '${line}s/.*/NEW_CONTENT/' FILE
```

##### insert a line after match #####
```
sed '/PATTERN/a NEW_CONTENT' FILE
```

##### insert a line before match #####
```
sed '/PATTERN/i NEW_CONTENT' FILE
```

##### search and replace #####
```
sed 's/PATTERN/NEW_CONTENT/g' FILE
```

### Perf ###
```
g++ -O3 main.cpp -g
perf record --call-graph dwarf -- yourapp
perf report -g graph --no-children 
```

### Non-interactive ssh ###
```
sshpass -p[passwd] ssh -oStrictHostKeyChecking=no [user]@[server] [cmd]
```

### List listened ports by a process
```
lsof -aPi [-Fn] -p <PID>
lsof -aPi4 [-Fn] -p <PID>  # IPv4
lsof -aPi6 [-Fn] -p <PID>  # IPv6
```

### Check which process is using a port
```
lsof -i :<PORT>
```

### User management
##### Send message to all logged-in users
```
wall -n <MESSAGE>
```
##### List all logged-in users
```
who -u  # You can kill -9 <PID> to kick off some logged-in user
```
##### Send message to a specific terminal
```
echo '<MESSAGE>' > /dev/pts/<ID>
```

### objdump
```
objdump -dj <section_name> <binary_name>
```
