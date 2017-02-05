# Hover Cli

Hover Cli is a simple Command Line Interface that allows the user to interacts with Hover APIs.  
For Hover reference see http://github.com/iovisor/iomodules  

## Install
```bash
go get github.com/mbertrone/hovercli/
go install github.com/mbertrone/hovercli/
```

## Launch
```bash
sudo $GOPATH/bin/hovercli -hover http://127.0.0.1:5002
```

## Usage

```bash
cli@hovercli$ help

HoverCli Command Line Interface HELP

        interfaces, i    prints /external_interfaces/
        modules, m       prints /modules/
        links, l         prints /links/
        table, t         prints /tables/

        help, h          print help


Modules Usage

        modules get
        modules get <module-id>
        modules post <module-name>
        modules delete <module-id>

Links Usage

        links get
        links get <link-id>
        links post <from> <to>
        links delete <link-id>

Table Usage

        table get
        table get <module-id>
        table get <module-id> <table-id>
        table get <module-id> <table-id> <entry-key>
        table put <module-id> <table-id> <entry-key> <entry-value>
        table delete <module-id> <table-id> <entry-key> <entry-value>
```
