# GoBranch
[![Go Report Card](https://goreportcard.com/badge/github.com/ipanardian/GoBranch)](https://goreportcard.com/report/github.com/ipanardian/GoBranch) 
[![Release](https://img.shields.io/badge/release-v0.0.0.1-orange.svg)](https://github.com/ipanardian/gobranch/releases)
[![Status](https://img.shields.io/badge/status-beta-green.svg)](https://github.com/ipanardian/gobranch/releases)
[![Go](https://img.shields.io/badge/go-v1.10.x-blue.svg)](https://gitter.im/ipanardian/gobranch)
[![GitHub license](https://img.shields.io/badge/license-MIT-red.svg)](https://github.com/ipanardian/GoBranch/blob/master/LICENSE)



A command line app to help you quickly creating git branch without hassle step.

## Usage
```
$ gobranch
```
![gobranch](https://user-images.githubusercontent.com/415225/44306664-cc81b880-a3bd-11e8-878f-73bc0551bfca.gif)

## Custom
![gobranch-custom](https://user-images.githubusercontent.com/415225/44306690-49ad2d80-a3be-11e8-9c2c-f626618486cc.gif)

## Features
* You don't need to type a valid branch name
* You don't need to checkout to the base branch
* You don't need to git pull the base branch
* Automatic prefixes such as feature, enhance, bugfix, hotfix
* Prevent the selected base branch from being wrong
* Custom base branch
* Custom naming conventions

## Requirements
* Git

## Installation
```
$ go get github.com/ipanardian/GoBranch
```
or just download the binary file on the release tab and add it to your ***$PATH*** or save it to ***/usr/local/bin***.

## Naming Conventions
```
$ gobranch --tc "/" --nc "snake"
//output: feature/abcd_efgh

$ gobranch --tc "-" --nc "kebab"
//output: feature-abcd-efgh
```

## Branch Tree
The following is a branch of the gobranch diagram, you must have a branch **development** and **hotfix**, unless you choose **custom** then there is no need to follow this diagram.
```
 -- master
    |-- development
    |   |-- feature_{name}
    |   |-- enhance_{name}
    |   |-- bugfix_{name}
    |   `-- test_{name}
    `-- hotfix
        |-- hotfix_{name}
        `-- hotfeature_{name}
```

## Flags
```
//show help
$ gobranch --help or -h

//show version
$ gobranch --version or -v

//Set type convention
$ gobranch --tc /
//output: feature/{branch}

//Set naming convention
$ gobranch --nc kebab
//output: branch-name
```

## License
The MIT License (MIT)
