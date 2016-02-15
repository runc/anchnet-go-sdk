# Go client for anchnet API

Go client library for [anchnet](http://cloud.51idc.com/help/api/api_list.html)

## Overview

The library preserves all semantics from anchnet APIs.

`anchnet/` is the CLI tool implementation based on the client, see [README](anchnet/README.md)

## Authentication

Authentication is done reading a config file `~/.anchnet/config`. Example file:
```json
{
  "publickey":  "U9WGEXYO19ysz607rLXwOyC",
  "privatekey": "K4XX2OPKMA2VrMo4WjLFbRMMH3djEfW94LK4d1W"
}
```
