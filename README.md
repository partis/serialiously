# serialiously

Opens a connection to serial com port, sends commands to the port and reads in the response.

## Download and compile serialiously

To compile Serialiously you’ll need to have installed golang, you can download it from here https://golang.org/dl/ and a number of tutorials are available to help you to configure it. 

The examples below use a Linux host to compile and run Serialiously.

To download from GitHub use the following to do a clone of the repository into your golang workspace

`git clone git@github.com:partis/serialiously.git`

Once you have cloned the repository and your GOPATH is configured to your workspace then run the following command to get the dependencies from GitHub.

`go get ./…`

Then run the commands below to compile Serialiously.

`go install {workspace_path}/serialiously`

For example I run

`go install github.com/partis/serialiously`

Which creates a binary file in the bin directory of my workspace.

## Configuring Serialiously

Firstly we need somewhere for the logs to go, so in the bin directory of your workspace create a log directory

`mkdir log`

Create a commands list file and add your commands in to it

`vi commands.lst
command1
command2
reboot
exit`

Create a config file, the file is a JSON formatted file currently with 4 settings and must be named serialiously.cfg;

`vi serialiously.cfg
{
  "commandFile": "commands.lst",
  "comPort": "/dev/ttyS1",
  "prompt": ">",
  "delay": "2"
}`

commandFile – specifies the commands file we created above

comPort – is the location of the com port to connect to, in the example this is serial port 1 so /dev/ttyS1 is specified

prompt – the prompt that the device returns when it is waiting for input, in this case a > signifies the device is waiting.

delay – the time in seconds to wait between sending commands.

And finally change the permissions of the binary so that it can be executed.

`chmod u+x serialiously`
 
## Running Serialiously

Once Serialiously is compiled and configured to run it is simple

`./serialiously`

To output logs to the log directory use

`./serialiously -log_dir=log`

Set the logging level using the option

`-stderrthreshold=[INFO|WARN|FATAL]`

Set it to debug logging by using –v

`-v=1 `

For debug and

`-v=2`

For trace.

If you specify the log directory as above all logs will be output to that directory, on Linux a symlink is created to the latest run using the format,

> App_name.log_level

So for Serialiously @ info level that would be

`Serialiously.INFO`
