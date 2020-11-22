
**qr-code** Docker image
========================

This Docker image contains an example API for the project "ln-paywall". For more info about the project visit https://github.com/philippgille/ln-paywall.

Prerequisites
-------------

- A running lnd node, either on a remote host and accessible from outside, or on the same host, in which case you can either start this container in "host network" mode, or use the container's gateway IP address to reach the host's localhost

Usage
-----

1. Create a data directory on the host: `mkdir data`
2. Copy the `tls.cert` and `invoice.macaroon` from your lnd to the `data` directory