# One Site Crawler

This is a client-server tool that crawls a single website.

## Usage

First, deploy the server. The easiest way is to run the Docker container:

```
docker run -p 8080:8080 wbrefvem/osc:0.0.21
```

Then, download and run the binary from the release page:

```
$ wget https://github.com/wbrefvem/osc/releases/download/0.0.21/osc
...
$ chmod +x osc
...
$ ./osc
A tool that allows you to crawl a single domain.

Usage:
  osc [command]

Available Commands:
  crawl       Use the specified server to crawl a domain
  get-domain  Get information about a crawled domain
  help        Help about any command

Flags:
  -h, --help   help for osc

Use "osc [command] --help" for more information about a command.
```

Optionally add it to your PATH:
```
export PATH=$PATH:$(pwd)/osc
```


The first command you should run is `osc crawl`, i.e.,
```
$ osc crawl -d http://quotes.toscrape.com -s http://127.0.0.1:8080
2020/03/29 22:33:42 started crawl for domain http://quotes.toscrape.com

--

The spiders are now crawling your domain!
Depending on the size of your site, this could take awhile.
Please wait a bit and check back with

osc get-domain <your domain>
```

How long you should wait before running `osc get-domain` depends on the size of your site, but when ready, do, e.g.:

```
$ osc get-domain quotes.toscrape.com -s http://127.0.0.1:8080
```

## Building from source


