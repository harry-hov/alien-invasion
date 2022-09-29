# Alien Invasion

Mad aliens are about to invade the earth and this program is to simulate the invasion.

#### Usage

```
$ ./alien-invasion --help
Mad aliens are about to invade the earth and this program is to simulate the invasion.

Usage:
  alien-invasion [command]

Available Commands:
  help        Help about any command
  invade      Invade a World

Flags:
  -h, --help   help for alien-invasion

Use "alien-invasion [command] --help" for more information about a command.
```

#### Invade Command

  ```
  $ ./alien-invasion invade --help
  Invade a World

  Usage:
    alien-invasion invade [world-file] [flags]

  Flags:
    -a, --aliens uint   Alien Count
    -h, --help          help for invade
  ```

## Running Locally

```
$ go build
$ ./alien-invasion invade worlds/world-1 --aliens 8
```

Output: 
```
Foo has been destroyed by alien-4, alien-5, alien-0, alien-1, and alien-2!
Bar has been destroyed by alien-3 and alien-7!
Conclusion: alien (alien-6) won

Remaining World:
Baz:
Qu-ux:
Bee:
```

Note: Output can be different for you.
