## linstor-gateway check-health

Check if all requirements and dependencies are met on the current system

### Synopsis

Check if all requirements and dependencies are met on the current system

```
linstor-gateway check-health [flags]
```

### Options

```
  -h, --help   help for check-health
```

### Options inherited from parent commands

```
      --config string         Config file to load (default "/etc/linstor-gateway/linstor-gateway.toml")
      --controllers strings   List of LINSTOR controllers to try to connect to (default from $LS_CONTROLLERS, or localhost:3370)
      --loglevel string       Set the log level (as defined by logrus) (default "info")
```

### SEE ALSO

* [linstor-gateway](linstor-gateway.md)	 - Manage linstor-gateway targets and exports

###### Auto generated by spf13/cobra on 1-Dec-2021