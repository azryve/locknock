# Locknock

Locknock is a cli utility to generate port knocking iptables ruleset and execute knocking against it.
The port knocking actually accomplished with udp payload using the same one target port.
This way all packets can be sent without worrying about reordering. The matching is perfomed by u32 module.
The payload sequence is generated automatically based on a provided pre-shared password.

It is in no way is protected against the replay attacks or any eavesdropping.
The intended use is only as a precaution against mass internet scans and proxy detection.

## Security preface
The port knocking in its essence is a *security by obscurity* aka **NOT A SECURITY**.

The technique is designed to *obfuscate* and **not** to protect.

Please contact your [physician](https://en.wikipedia.org/wiki/Port_knocking) to see if the port knoking is suitable for you.

## Basic usage

### Download and install
```
curl https://raw.githubusercontent.com/azryve/locknock/main/download.py | python3
sudo install -m755 ./locknock /usr/local/bin/
```

### Server: generate and install iptables ruleset

```
export LOCKNOCK_PASSWORD=myverystrongpass
locknock ruleset | sudo bash
iptables -A INPUT -j LOCKNOCK
```

### Client: knock and open the port

```
locknock knock myserver.example.com
nc -vz myserver.example.com 22
```

## Ssh proxy command

Openssh client's ProxyCommand option allows to execute knocking before connecting to server automatically.

1) put `LOCKNOCK_PASSWORD` in your profile file
2) add following in the ~/.ssh/config:

```
Host myserver.example.com
	ProxyCommand locknock knock %h -P %p
```
