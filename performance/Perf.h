/*
 ******************* How to use this *******************
 *                                                     *
 *     #include "Perf.h"                               *
 *                                                     *
 *     // perf before                                  *
 *     Perf::profile("stat", [&](){                    *
 *              func_to_be_perf_stat();                *
 *          });                                        *
 *     Perf::profile("record", [&](){                  *
 *              func_to_be_perf_record();              *
 *          });                                        *
 *     // perf after                                   *
 *                                                     *
 *******************************************************
 */
 
#ifndef PERF_H_
#define PERF_H_

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
            // child process
            // run perf here
            char buf[50];
            std::sprintf(buf, "perf %s -p %d > %s.data 2>&1", type, pid, type);
            execl("/bin/sh", "sh", "-c", buf, nullptr);
        } else {
            // set the father process as the leader of its process group
            setpgid(cpid, 0);
            // part of program you wanted to perf
            body();
            // stop perf by killing child process and all its descendants(sh, perf stat etc)
            kill(-cpid, SIGINT);
            wait(nullptr);
            // rest of the program
        }
    }
};

#endif
