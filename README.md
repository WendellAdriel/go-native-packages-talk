你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
