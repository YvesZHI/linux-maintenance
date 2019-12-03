#include <cstdio>
#include <functional>
#include <sys/stat.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <unistd.h>
#include <signal.h>

struct Perf
{
    static void profile(const char *type, std::function<void()> body)
    {
        int pid= getpid();
        int cpid = fork();

        if(cpid == 0) {
            // child process .  Run your perf stat
            char buf[50];
            std::sprintf(buf, "perf %s -p %d > %s.data 2>&1", type, pid, type);
            execl("/bin/sh", "sh", "-c", buf, NULL);
        } else {
            // set the child the leader of its process group
            setpgid(cpid, 0);
            // part of program you wanted to perf stat
            body();
            // stop perf by killing child process and all its descendants(sh, perf stat etc )
            kill(-cpid, SIGINT);
            wait(nullptr);
            // rest of the program
        }
    }
};
