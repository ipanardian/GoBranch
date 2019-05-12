# GoBranch
[![Go Report Card](https://goreportcard.com/badge/github.com/ipanardian/GoBranch)](https://goreportcard.com/report/github.com/ipanardian/GoBranch) 
[![Release](https://img.shields.io/badge/release-v0.0.0.2-orange.svg)](https://github.com/ipanardian/gobranch/releases)
[![Status](https://img.shields.io/badge/status-beta-green.svg)](https://github.com/ipanardian/gobranch/releases)
[![Go](https://img.shields.io/badge/go-v1.10.x-blue.svg)](https://gitter.im/ipanardian/gobranch)
[![GitHub license](https://img.shields.io/badge/license-MIT-red.svg)](https://github.com/ipanardian/GoBranch/blob/master/LICENSE)



A command line app to help you quickly creating git branch without hassle steps.

## Usage
```
$ GoBranch
```
![gobranch](https://user-images.githubusercontent.com/415225/44306664-cc81b880-a3bd-11e8-878f-73bc0551bfca.gif)

## Features
* You don't need to type a valid branch name
* You don't need to checkout to the base branch
* You don't need to git pull the base branch
* Automatic prefixes such as feature, enhance, bugfix, hotfix, release
* Prevent the selected base branch from being wrong
* Custom base branch
* Custom naming conventions

## Requirements
* Git

## Installation with Go
```
$ go get github.com/ipanardian/GoBranch
```

## Update GoBranch
```
$ go get -u github.com/ipanardian/GoBranch
```

## Installation executable file only
- Open the [release](https://github.com/ipanardian/GoBranch/releases) section
- Download binary files that match with your operating system (OS)
- Change the name to "GoBranch" or "GoBranch.exe" for windows
- add it to your ***$PATH*** environment variable, so you can run it from any location on the command line
- On mac copy to /usr/local/bin
- Then sudo chmod +x /usr/local/bin/GoBranch

## Naming Conventions
```
$ GoBranch --tc "/" --nc "snake"
//output: feature/abcd_efgh 

$ GoBranch --tc "-" --nc "kebab"
//output: feature-abcd-efgh

//Default: feature/abcd_efgh 
```

## Branch Tree
The following is a branch of the GoBranch tree, you must have a branch **development** and **hotfix**, unless you choose **custom** then there is no need to follow this diagram.
```
 -- master
    |-- development
    |   |-- feature/{name}
    |   |-- enhance/{name}
    |   |-- bugfix/{name}
    |   `-- test/{name}
    `-- hotfix
        |-- hotfix/{name}
-- release
```

## Flags
```
//show help
$ GoBranch --help or -h

//show version
$ GoBranch --version or -v

//Set type convention
$ GoBranch --tc /
//output: feature/{branch}

//Set naming convention
$ GoBranch --nc kebab
//output: branch-name
```

## License
The MIT License (MIT)
