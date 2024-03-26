# Locknock

Locknock is a cli utility to generate port knocking iptables ruleset and execute knocking against it.
The port sequence is generated automatically based on a provided pre-shared password.

## Security preface
The port knocking in its essence is a *security by obscurity* aka **NOT A SECURITY**.

The technique is designed to *obfuscate* and **not** to protect.

Please contact your [physician](https://en.wikipedia.org/wiki/Port_knocking) to see if the port knoking is suitable for you.

## Basic usage

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

```
# cat ~/.ssh/config
Host myserver.example.com
	ProxyCommand locknock knock %h --port-proxy %p
```
