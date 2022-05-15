# RegisterAndLogin Service

This is the RegisterAndLogin service

Generated with

```
micro new --namespace=go.micro --type=service registerAndLogin
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.registerAndLogin
- Type: service
- Alias: registerAndLogin

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./registerAndLogin-service
```

Build a docker image
```
make docker
```