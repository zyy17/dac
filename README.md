# DAC

`dac`(dashboard as code) is a tool to transform native Grafana dashboard into a lightweight dashboard that can be treated as a intermediate representation for Grafana dashboard.

Compared with the relative closed Grafana dashboard JSON model, `dac` provides a more open and human-readable intermediate representation for Grafana dashboard.

## Usage

### Docker

```console
docker run --rm ghcr.io/zyy17/dac:latest -i dashboard.json
```

## Development

Use the following command to compile:

```console
make
```

And the `dac` binary will be in `bin/`.
