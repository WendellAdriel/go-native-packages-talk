# Go Native Packages Talk

> Simple example created to use in a talk

## How to use

- Clone this repo and enter it

- Run the `build.sh` script, it will create three binaries for the application (**windows/amd64** **windows/386** **darwin/amd64**) inside a `builds/` directory

- If you pass a name to the `build` script, it will use it to create your binaries. Ex: If you run the script like this: `build.sh myapp` it will create: **myapp_windows_amd64.exe**, **myapp_windows_386.exe** and **myapp_darwin_amd64**, otherwise it will use **my-go-server** as the default name.

- While the **HTTP Server** is running you can stop it with `Ctrl + C` in the terminal. It will shutdown the server gracefully.

## Commands available

Below is a list of commands you can use. The examples will use the default app name: **my-go-server**, if you built it with a **custom name**, use it instead of the default to run the commands.

- To show the application info (version, name, etc) run:

```
builds/my-go-server -version
OR
builds/my-go-server -v
```

- To show the application options run:

```
builds/my-go-server -help
OR
builds/my-go-server -h
```

- To start the **HTTP Server** at the default port (defined at `config/config.go`) run:

```
builds/my-go-server -start
OR
builds/my-go-server -s
```

- To start the **HTTP Server** at a custom port run (like 5000):

```
builds/my-go-server -start -port 5000
OR
builds/my-go-server -s -p 5000
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/WendellAdriel/go-native-packages-talk/.
This project is intended to be a safe, welcoming space for collaboration.

## License

This project is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).