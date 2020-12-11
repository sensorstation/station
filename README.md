# Sensor Station

Sensor Station is an Open Framework intended to help faciliate the
develop and deploy _Intelligent Online Techinologies (IoT)_.

Really, sensor station is Open Source Software as set of _Best
Practices_ built on _Open Protocols_ with reference applications. For
example, the only thing required of a "sensor station" is that it can
_publish_ to and _subscribe_ from MQTT channels, or _topics_.

To that regard, we have built two versions of _Sensor Station_, 1 st)
a Raspberry PI version and an pure embedded esp32 version.  Since they
all speak an opinionated formation of message channels, these
independent implementations can work to gether just fine.

## Build

1. Install Go 
2. go get ./...
3. cd ss; go build 

That should leave the execuable 'ss' in the 'ss' directory as so:

> ./station/ss/ss

## Fake Websocket Data

```bash
% ./ss -fake-ws
```

```bash
% ./ss -help
```

This will open the following URL for the fake websocket data:

> http://localhost:8011/ws

Replace localhost with a hostname or IP if needed. Have the websocket
connect to the URL and start spitting out fake data formatted like
this:

```json
{"year":2020,"month":12,"day":10,"hour":20,"minute":48,"second":8,"action":"setTime"}
{"K":"tempf","V":88}
{"K":"soil","V":0.49}
{"K":"light","V":0.62}
{"K":"humid","V":0.12}
```

Ignore the "time" data if it is a pain in the neck..

