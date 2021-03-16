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

#### raspi-config

Login with the default login `pi/raspbery` and run `rasp-config`

```bash
raspi-config
```

* Change the password for the `pi` user
* Change the hostname to `pi1` and `pi2`
* Enable ssh under "Interface Options"

#### Static IP address

Configure the `eth0` interface statically by editing the `/etc/dhcpcd.conf` file.

For the hostname `pi1`:

```bash
interface eth0
static ip_address=192.168.2.160/24
static routers=192.168.2.1
static domain_name_servers=192.168.2.1 8.8.8.8
```

For the hostname `pi2`:

```bash
interface eth0
static ip_address=192.168.2.161/24
static routers=192.168.2.1
static domain_name_servers=192.168.2.1 8.8.8.8
```

#### Prometheus

The following is based on this [article](https://pimylifeup.com/raspberry-pi-prometheus/).

```bash
sudo apt update
sudo apt full-upgrade

wget https://github.com/prometheus/prometheus/releases/download/v2.22.0/prometheus-2.22.0.linux-armv7.tar.gz
tar xfz prometheus-2.22.0.linux-armv7.tar.gz
mv prometheus-2.22.0.linux-armv7/ prometheus/
rm prometheus-2.22.0.linux-armv7.tar.gz

sudo nano /etc/systemd/system/prometheus.service
```

Past in this text:


```text
[Unit]
Description=Prometheus Server
Documentation=https://prometheus.io/docs/introduction/overview/
After=network-online.target

[Service]
User=pi
Restart=on-failure

ExecStart=/home/pi/prometheus/prometheus \
  --config.file=/home/pi/prometheus/prometheus.yml \
  --storage.tsdb.path=/home/pi/prometheus/data

[Install]
WantedBy=multi-user.target
```

```
sudo systemctl enable prometheus
sudo systemctl start prometheus
sudo systemctl status prometheus

open http://192.168.2.161:9090
```

#### Prometheus Push Gateway

This section base on this [github page](https://github.com/prometheus/pushgateway/blob/master/README.md)

```bash
wget https://github.com/prometheus/pushgateway/releases/download/v1.4.0/pushgateway-1.4.0.linux-armv7.tar.gz


tar xfz pushgateway-1.4.0.linux-armv7.tar.gz
mv pushgateway-1.4.0.linux-armv7/ pushgateway/
rm pushgateway-1.4.0.linux-armv7.tar.gz

# test  it
echo "aeg_metric 3.14" | curl --data-binary @- http://192.168.2.161:9091/metrics/job/aeg_job

#of
cat <<EOF | curl --data-binary @- http://192.168.2.161:9091/metrics/job/some_job/instance/some_instance
# TYPE some_metric counter
some_metric{label="val1"} 42
# TYPE another_metric gauge
# HELP another_metric Just an example.
another_metric 2398.283
EOF
```

The query API allows accessing pushed metrics and build and runtime information.

```bash
curl -X GET http://192.168.2.161:9091/api/v1/status | jq
```

Configure pushgateway as a scrape target in prometheus,
**EDIT**: /home/pi/prometheus/prometheus.yml (must have correct indentation)

```yaml
...
scrape_configs:
  ...
  - job_name: 'pushgateway'
    honor_labels: true
    static_configs:
    - targets: ['localhost:9091']
...

# Then restart 
sudo systemctl restart prometheus
```

todo - need to make the scraper a service

### metrics

TODO 
echo "temp1f 59" | curl -v --data-binary @- http://192.168.2.161:9091/metrics/job/sensor_reading/location/greenhouse


```text
/metrics/job/sensor_reading/location/greenhouse

# TYPE temp1f gauge
# HELP temp1f Greenhouse Tempature
temp1f 2398.283
```

### Static Routs

This section based on [this](https://wiki.dd-wrt.com/wiki/index.php/Linking_Subnets_with_Static_Routes)


on 1
![](img/subnet-route.png)


on 2

```
# Allow everything to be forwarded through the router (simple but do not use on routers directly connected to the internet)
iptables -I FORWARD -j ACCEPT
```