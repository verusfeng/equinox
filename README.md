# equinox

0. all bin file will complie on local pc and then upload to server.
## Step

1. How it works
    - Equinox helps you sign, package and distribute self-updating Go programs. Equinox is made up of three parts:

    - The Equinox release tool, a small CLI tool that wraps go build
    - The Equinox SDK, a small go package that adds self-updating functionality to your app.
    - The Equinox service, that hosts your binaries, download pages and update patches
2. release tool 
    - https://bin.equinox.io/c/mBWdkfai63v/release-tool-stable-windows-386.zip
    - unzip file equinox.exe
3. equinox.exe genkey  生成非对称加密密钥
4. release
```example
equinox release \
  --version="0.1.0" \
  --platforms="darwin_amd64 linux_amd64" \
  --signing-key=equinox.key \
  --app="app_jTRogDU1urZ" \
  --token="4kj5ypgZj3QeGvGz7LckDGvcAcdmUozUNU8YhVhEg97r7dcmFgy" \
  github.com/example/tool

```

5. adding-update-code

## release
To build and release this program with Equinox, replace go build with equinox release. The Equinox-specific options are covered in the publishing section below.

$ equinox release [options] github.com/acme/rocket

1. other cmd 
	- equinox.exe release --config .\config.yaml --version [verson] --channel stable [github_packageName]
		- .config.yaml configure file .
		
```yaml
# saved in config.yaml
app: app_96H5xG6nZm1
signing-key: ./path/to/equinox.key
token: YOUR_TOKEN_HERE
platforms: [
  # darwin_amd64, darwin_386,
  # linux_amd64, linux_386,
  # windows_amd64, windows_386
]

```