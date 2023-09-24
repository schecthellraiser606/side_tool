from pwn import *

context.log_level = "debug"

binfile = './chall'
libcfile = ''
offset = b'A'*24

rhost = 'rop-2-35.seccon.games'
rport = 9999


# if you use system func
# set follow-fork-mode parent
gdb_script = '''
b main
'''

elf = ELF(binfile)
context.binary = elf
# libc =ELF(libcfile)

def conn():
    if args.REMOTE:
        p = remote(rhost, rport)
    elif args.GDB:
        p = process(elf.path)
        gdb.attach(p, gdbscript=gdb_script)
    else:
        p = process(elf.path)
    return p

p = conn()


# coding here
# p.sendline(b'/bin/sh')

p.recvrepeat()
p.interactive()

# After Pwned code
# print(p.recvrepeat().decode())