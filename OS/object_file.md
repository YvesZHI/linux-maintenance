# Object file


### Reading file without IO function ###
There is a file, named `custom.config`, its content as below:
```
name titi
password 123
```

1) Convert it into an obj file with the command: `objcopy -I binary -O elf64-x86-64 -B i386 custom.config custom.config.o`;<br>
2) Execute the command `objdump -x -s custom.config.o` to get the content of the object file:<br>
```
custom.config.o:     file format elf64-x86-64
custom.config.o
architecture: i386:x86-64, flags 0x00000010:
HAS_SYMS
start address 0x0000000000000000

Sections:
Idx Name          Size      VMA               LMA               File off  Algn
  0 .data         00000017  0000000000000000  0000000000000000  00000040  2**0
                  CONTENTS, ALLOC, LOAD, DATA
SYMBOL TABLE:
0000000000000000 l    d  .data  0000000000000000 .data
0000000000000000 g       .data  0000000000000000 _binary_custom_config_start
0000000000000017 g       .data  0000000000000000 _binary_custom_config_end
0000000000000017 g       *ABS*  0000000000000000 _binary_custom_config_size


Contents of section .data:
 0000 6e616d65 20746974 690a7061 7373776f  name titi.passwo
 0010 72642031 32330a                      rd 123.
```
3) Coding C as below:
```
// main.c
#include <stdio.h>
#include <string.h>

extern const char _binary_custom_config_start[];
extern const char _binary_custom_config_size;

int main()
{
    size_t size = (size_t)&_binary_custom_config_size;
    char config[size + 1];
    strncpy(config, _binary_custom_config_start, size);
    config[size] = '\0';
    printf("config = \"%s\"\n", config);
    return 0;
}
```
4) Compile and build: `gcc test.c custom.config.o`.

### Others ###
1) To move the content to another section (`.rodata` as an example), execute the command `objcopy -I binary -O elf64-x86-64 -B i386 custom.config custom.config.o --rename-section .data=.rodata,alloc,load,readonly,data,contents`.<br>
2) GCC allows specifying the section for data. For example, `__attribute__((section("FOO"))) int global_var = 42;` means put the global variable `global_var` into the section named `FOO`.
