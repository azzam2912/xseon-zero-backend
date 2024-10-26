# [LOGGING](https://skorlife.atlassian.net/wiki/spaces/EF/pages/3276805/Logging+Philosophy)

## Installation

To install run the following command inside your project directory

```sh
go get github.com/skortech/st-kit/logging
```

## Vendoring to your project (recommended for all projects)

```sh
git clone ssh://git@github.com/skorlife/st-kit.git
```

## Usage Example

### Logger Initialization

```go
package main

import (
    "github.com/skortech/st-kit/logging"
)

func main() {
    logger := logging.New("sk-risk-module")
    logger.Log(
        Critical,
        "ERR.SK.CONNECTION",
        "error fetching abc from xyz",
        logging.WithDocument("https://abc.xyz/"),
        logging.WithIdentity("SL000141311"),
        logging.WithReference(map[string]interface{}{
             "host": "127.0.0.1",
        }),
        logging.WithSync(true),
 )
}
```
