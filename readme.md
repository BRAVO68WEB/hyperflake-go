<h1 align="center" style="border-bottom: none;">❄️ Hyperflake </h1>

<h3 align="center">A simple and lightweight Go module to generate unique snowflake-like IDs starting from a custom epoch. </h3>

<br />

<p align="center">
<a href="https://github.com/chirag3003/hyperflake-go/actions/workflows/test.yml">
<img alt="Build states" src="https://github.com/chirag3003/hyperflake-go/actions/workflows/test.yml/badge.svg?branch=main">
</a>
</p>

<p align="center">
<a href="https://github.com/chirag3003/hyperflake-go/issues/new">Bug report</a>
</p>

<hr />

`hyperflake-go` is a Go library for generating unique and distributed IDs that are suitable for use as primary keys in distributed systems.

It generates 64-bit IDs that are composed of a timestamp, a datacenter ID, a machine ID, and a sequence number. These IDs are based on [Twitter's Snowflake ID](https://github.com/twitter-archive/snowflake/tree/snowflake-2010) generation algorithm.

## Installation 🚀

You can install `hyperflake-go` by adding it to your project:

```bash
go get github.com/chirag3003/hyperflake-go
```

## Methods 🧮

The `Config` instance has the following methods:

- `GenerateHyperflakeID()`: Generates a unique ID.
- `DecodeID(id int64)`: Decodes a given Hyperflake ID into its components.
- `SetMachineID(machineID int)`: Sets the machine ID in the configuration.
- `SetDatacenterID(datacenterID int)`: Sets the datacenter ID in the configuration.
- `GetMachineID()`: Returns the machine ID from the configuration.
- `GetDatacenterID()`: Returns the datacenter ID from the configuration.

## Usage 💻

Here's an example of how to use `hyperflake-go`:

```go
package main

import (
    "fmt"
    "github.com/chirag3003/hyperflake-go"
)

func main() {
    config := hyperflake.NewHyperflakeConfig(5, 5)
    id, err := config.GenerateHyperflakeID()
    if err != nil {
        fmt.Println("Error generating ID:", err)
        return
    }
    fmt.Println("ID:", id)  // ID: 3273974649684016911
}
```

This will generate a unique ID.

## Validating IDs ✅

You can decode IDs generated by `hyperflake-go` using the `DecodeID` method:

```go
package main

import (
    "fmt"
    "github.com/chirag3003/hyperflake-go"
)

func main() {
    config := hyperflake.NewHyperflakeConfig(5, 5)
    id, err := config.GenerateHyperflakeID()
    if err != nil {
        fmt.Println("Error generating ID:", err)
        return
    }
    fmt.Println("ID:", id)  // ID: 3273974649684016911

    decodedID, err := hyperflake.DecodeID(id)
    if err != nil {
        fmt.Println("Error decoding ID:", err)
        return
    }
    fmt.Printf("Decoded ID: %+v\n", decodedID)
}
```

This will return the components of the generated ID.

## Error Handling 😱

The `Config` instance throws an error if the clock moves backwards, i.e., if the current timestamp is less than the last timestamp.

This can happen if the system clock is adjusted manually or if the system clock drifts significantly.

If this happens, the library throws an error with the message `Clock is moving backwards!`.

## Examples 🔠

Here's an example of how to generate 10 IDs:

```go
package main

import (
    "fmt"
    "github.com/chirag3003/hyperflake-go"
)

func main() {
    config := hyperflake.NewHyperflakeConfig(5, 5)

    for i := 0; i < 10; i++ {
        id, err := config.GenerateHyperflakeID()
        if err != nil {
            fmt.Println("Error generating ID:", err)
            continue
        }
        fmt.Println("ID:", id)
        decodedID, err := hyperflake.DecodeID(id)
        if err != nil {
            fmt.Println("Error decoding ID:", err)
            continue
        }
        fmt.Printf("Decoded ID: %+v\n", decodedID)
    }
}
```

## Bugs or Requests 🐛

If you encounter any problems feel free to open an [issue](https://github.com/chirag3003/hyperflake-go/issues/new) on GitHub.