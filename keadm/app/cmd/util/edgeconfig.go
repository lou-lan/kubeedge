package util

var (
	EdgeYaml = []byte(`mqtt:
    server: tcp://127.0.0.1:1883 # external mqtt broker url.
    internal-server: tcp://127.0.0.1:1884 # internal mqtt broker url.
    mode: 0 # 0: internal mqtt broker enable only. 1: internal and external mqtt broker enable. 2: external mqtt broker enable only.
    qos: 0 # 0: QOSAtMostOnce, 1: QOSAtLeastOnce, 2: QOSExactlyOnce.
    retain: false # if the flag set true, server will store the message and can be delivered to future subscribers.
    session-queue-size: 100 # A size of how many sessions will be handled. default to 100.

edgehub:
    websocket:
        url: wss://0.0.0.0:10000/e632aba927ea4ac2b575ec1603d56f10/fb4ebb70-2783-42b8-b3ef-63e2fd6d242e/events
        certfile: /etc/kubeedge/certs/edge.crt
        keyfile: /etc/kubeedge/certs/edge.key
        handshake-timeout: 30 #second
        write-deadline: 15 # second
        read-deadline: 15 # second
    controller:
        placement: false
        heartbeat: 15  # second
        refresh-ak-sk-interval: 10 # minute
        auth-info-files-path: /var/IEF/secret
        placement-url: https://10.154.193.32:7444/v1/placement_external/message_queue
        project-id: e632aba927ea4ac2b575ec1603d56f10
        node-id: fb4ebb70-2783-42b8-b3ef-63e2fd6d242e

edged:
    register-node-namespace: default
    hostname-override: fb4ebb70-2783-42b8-b3ef-63e2fd6d242e
    interface-name: eth0
    node-status-update-frequency: 10 # second
    device-plugin-enabled: false
    gpu-plugin-enabled: false
    image-gc-high-threshold: 80 # percent
    image-gc-low-threshold: 40 # percent
    maximum-dead-containers-per-container: 1
    docker-address: unix:///var/run/docker.sock
    version: 2.0.0

`)

	EdgeLoggingYaml = []byte(`loggerLevel: "DEBUG"
enableRsyslog: false
logFormatText: true
writers: [stdout]
`)

	EdgeModulesYaml = []byte(`modules:
    enabled: [eventbus, servicebus, websocket, metaManager, edged, twin, dbTest]
`)

	EdgeNodeJSONContent = `{
"kind": "Node",
"apiVersion": "v1",
"metadata": {
    "name": "fb4ebb70-2783-42b8-b3ef-63e2fd6d242e",
        "labels": {
            "name": "edge-node"
        }
    }
}`
)
