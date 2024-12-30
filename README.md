# fee_demo
- Demonstration of tool `fileless-elf-exec` (https://github.com/nnsee/fileless-elf-exec)
- Code for shell taken from https://gist.github.com/magisterquis/ce8b99b42188f4517efb3e1038a483de
### Usage
- Tutorial for usage on fedora Linux
- Install golang and fee: `bash install.sh`
- Compile countdown code: `bash compile_odb.sh`
- Compile shell code: `bash compile_shell.sh`
- Start countdown executable "normaly": `./odbrojavanje/odbrojavanje`
- Start countdown executable with `fee`: `python fee_odb.py`
- Start shell executable "normaly": `./shell/demoshell`
- Start shell executable with `fee`: `python fee_shell.py`
### Shell demo
- After starting the program, in the second terminal execute the following command: `rlwrap -S "$(printf '\033[95mds>\033[m ')" nc -nvlp 4444`
- You should enter the interactive shell in which you can execute any command
- The point of the attack is to execute demoshell on victim's computer and configure it to send a callback to attacker's computer
- In this scenario, victim and attacker are on the same host for simplicity
- When executed with `fee` (i.e. without using reading source file from disk) it is harder to trace and detect the malware (in this case demoshell)
### Difference
