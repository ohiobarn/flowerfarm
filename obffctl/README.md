# obffctl

The OBFF CLI

## Development

1. Use cobra to add a new obffctl command and/or subcommand. Below are a few examples:

```bash
# First time
mkdir obffctl
cd obffctl

go mod init github.com/ohiobarn/flowerfarm/obffctl
cobra init --pkg-name github.com/ohiobarn/flowerfarm/obffctl

# Add new "foo" command
# Usage: $ opsctl foo 
$ cobra add foo

# Add new "bar" subcommand to the existing "create" command
# Usage: $ opsctl create bar 
$ cobra add bar -p createCmd

# to install
$ go install github.com/ohiobarn/flowerfarm/obffctl

```

1. Cobra will add the source files needed for the command. You will need to edit the source files and add specific logic for your command.