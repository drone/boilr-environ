An environment variable extension to {{Description}}. _Please note this project requires Drone server version 1.7 or higher._

## Installation

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the extension:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --restart=always \
  --name=drone-environ {{DockerRepository}}
```

Update your runner configuration to include the extension address and the shared secret.

```text
DRONE_ENV_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_ENV_PLUGIN_TOKEN=bea26a2221fd8090ea38720fc445eca6
```

## Testing

Use the command line tools to test your extension. _This extension uses http-signatures to authorize client access and will reject unverified requests. You will be unable to test this extension using a simple curl command._

```text
export DRONE_ENVIRON_ENDPOINT=http://1.2.3.4:3000
export DRONE_ENVIRON_SECRET=bea26a2221fd8090ea38720fc445eca6

drone plugins env
```

## License

This software is licensed under the [Blue Oak Model License 1.0.0](https://spdx.org/licenses/BlueOak-1.0.0.html).
