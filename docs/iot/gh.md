# Greenhouse IOT 

## Temperature Monitoring

![](img/gh-iot.png)

The drawing above shows how the sensor data emanating from *ObserverIP* is intercepted via `iptables` running on *WRT54G* and sent to a Node-RED service running on `pi1`

### iptables

An *iptables* entry is created to intercept traffic from *ObserverIP* to *ambientweather.net*. and redirect it Node-RED running on a Raspberry Pi. The *gh-flow* running on the Pi will act as a tee, sending the data to Prometheus before sending it along to its original target *ambientweather.net*.

```bash
#
# ssh to WRT54G
#
ssh root@192.168.2.1 

#
# While on pi1
# redirect traffic from ObserverIP to pi1
#
iptables -t nat -A PREROUTING -s 192.168.2.151 -p tcp --dport 80 -j DNAT --to-destination 192.168.2.160:1880

# Verify
iptables -t nat -L PREROUTING

```

### Pi Setup

Load raspbian lite OS and then set a static IP. For more information see the official [doc](https://www.raspberrypi.org/documentation/configuration/tcpip/)

#### Static IP address

If you wish to disable automatic configuration for an interface and instead configure it statically, add the details to /etc/dhcpcd.conf. For example:

```bash
interface eth0
static ip_address=192.168.2.160/24
static routers=192.168.2.254
static domain_name_servers=192.168.2.254 8.8.8.8
```
