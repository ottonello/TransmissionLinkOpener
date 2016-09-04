# TransmissionLinkOpener

Sends a magnet/torrent link to a Transmission Daemon running somewhere else, uses https://trac.transmissionbt.com/wiki/rpc.   


Useful for starting downloads with a single click from the browser. In Windows this requires modifying/adding some registry keys as detailed in https://msdn.microsoft.com/en-us/library/aa767914(v=vs.85).aspx.
It also supports sending Basic auth but it's not very safe as you'll need to store your username/password plain in the registry.

Sample Windows registry configuration for Magnet links:
```
Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\Magnet]
@="Magnet URI"
"URL Protocol"=""
"Content Type"="application/x-magnet"

[HKEY_CLASSES_ROOT\dgh\DefaultIcon]
@="C:\Users\Test\AppData\Roaming\uTorrent\maindoc.ico"

[HKEY_CLASSES_ROOT\Magnet\shell]
@="open"

[HKEY_CLASSES_ROOT\Magnet\shell\open]

[HKEY_CLASSES_ROOT\Magnet\shell\open\command]
@="\"C:\\go_workspace\\bin\\transmissionLinkOpener.exe\" -t \"http://192.168.1.1:9091/transmission/rpc\" -l \"%1\""
```
