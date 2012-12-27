About
==========

Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.

Installation
==========

Installation consists of several simple steps. They may be a bit different on your target system (e.g. require more permissions) so adapt them to the parameters of your system.

### Get libtextcat C library code

* Download original libtextcat archive from [libtextcat download section](http://software.wise-guys.nl/libtextcat/). 
* Unarchive it.

NOTE: If this link is not working or there are some problems with downloading, there is a stable version 2.2 snapshot saved in [Downloads](https://github.com/downloads/goodsign/libtextcat/libtextcat-2.2.tar.gz).

### Build and install libtextcat C library

From the directory, where you unarchived libtextcat, run:

```
./configure
make
sudo make install
sudo ldconfig
```

### Install Go wrapper

```
go get github.com/goodsign/libtextcat
go test github.com/goodsign/libtextcat (must PASS)
```

Installation notes
==========

Make sure that you have your local library paths set correctly and that installation was successful. Otherwise, **go build** or **go test** may fail.

libtextcat is installed in your local library directory (e.g. **/usr/local/lib**) and puts its libraries there. This path should be registered in your system (using ldconfig or exporting LD_LIBRARY_PATH, etc.) or the linker would fail.

Usage
==========

```go
cat, err := NewTextCat(ConfigPath) // See 'Usage notes' section

if nil != err {
    // ... Handle error ...
}
defer cat.Close()

matches, err := cat.Classify(text)

if nil != err {
    // ... Handle error ...
}

// Use matches. 
// NOTE: matches[0] is the best match.

```

Usage notes
==========

libtextcat library needs to load language models to start guessing languages. These models are set using a configuration file and a number of language model (.lm) files.

Configuration file maps .lm files to identifiers used in the library. See [example](https://github.com/goodsign/libtextcat/blob/master/defaultcfg/conf.txt). Path to this file is specified in the **NewTextCat** call.

.lm files contain language patterns and frequencies for a specified language. See [example](https://github.com/goodsign/libtextcat/blob/master/defaultcfg/english.lm). Paths to these files are specified in the config file above. They can be absolute or relative (to the caller).

Quickstart
----------

To immediately get started, copy **/defaultcfg** folder contents to the directory of your target project and use:

```go
cat, err := NewTextCat("defaultcfg/conf.txt")
```

This will give you a standard set of languages described in the **Default configuration** section below.

Default configuration
----------

This package contains a default configuration (/defaultcfg) which is created to work in following conditions:

* Utf-8 only languages
* Language list is taken from [snowball](github.com/goodsign/snowball) package
* Language identifiers are the same as in [snowball](github.com/goodsign/snowball) package

This configuration is meant to be used in pair with the [snowball](github.com/goodsign/snowball) package.

More info
----------

For more information on libtextcat refer to the original [website](http://software.wise-guys.nl/libtextcat/), which contains links on theory and other details.

libtextcat Licence
==========

The libtextcat library is released under the [BSD Licence](http://opensource.org/licenses/bsd-license.php)

[LICENCE file](https://github.com/goodsign/libtextcat/blob/master/LICENCE_libtextcat)

Licence
==========

The goodsign/libtextcat binding is released under the [BSD Licence](http://opensource.org/licenses/bsd-license.php)

[LICENCE file](https://github.com/goodsign/libtextcat/blob/master/LICENCE)