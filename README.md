# What's this? 
I was curious about the difference of [GNU yes.c](https://github.com/coreutils/coreutils/blob/master/src/yes.c) and 
[OpenBSD yes.c](https://github.com/openbsd/src/blob/master/usr.bin/yes/yes.c). 

I implemented both variants in Go. Here's my fastest variant with around 8 GiB data throughput per second. 
For the sake of completeness, a simple for {} loop like this: 

```go
for {
    fmt.Println("y")
}
```

has a throughput of just 3-4 MiB per second. 

## How to use it
```sh
git clone https://github.com/mnlwldr/yes
```

## Usage and output
```sh
go run Main.go | pv -r > /dev/null
[8,10GiB/s]
```

You can install `pv` from [here](https://www.ivarch.com/programs/pv.shtml). 