# RPC Claymore

Package to take information about claymore status

## Installation

```go
go get github.com/ivandelabeldad/rpc-claymore
```

## Usage

```go
miner := claymore.Miner{Address: "localhost:3333"}
info, err := miner.GetInfo()
if err != nil {
  log.Fatal(err)
}
fmt.Printf("%v", info)
```
Output:
```
Version:   11.8
Up Time:   23 min

Main Crypto
HashRate:           119162 Mh/s
Shares:                 16
Rejected Shares:         0
Invalid Shares:          0

Alt Crypto
Disabled

Main Pool
Address:   eth-eu1.nanopool.org:9999
Switches:  0

Alt Pool
Disabled

GPU 0
Hash Rate:        29779 Mh/s
Alt Hash Rate:        0 Mh/s
Temperature:         47 º
Fan Speed:           60 %

GPU 1
Hash Rate:        29798 Mh/s
Alt Hash Rate:        0 Mh/s
Temperature:         49 º
Fan Speed:           60 %

```

You can access either each field on its own:

```go
info.MainCrypto.HashRate \\ int 119313
...
```

## Warning

Because of claymore bad implementation of the json rpc protocol there is no way to
keep it working with a password. So this option is useless until claymore received an update
to support it (May be never).

## License

RPC Claymore is open-sourced software licensed under
the [MIT license](https://github.com/ivandelabeldad/rpc-claymore/blob/master/LICENSE).
