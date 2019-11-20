# release-notes
[![Build Status](https://travis-ci.com/corneliusweig/release-notes.svg?branch=master)](https://travis-ci.com/corneliusweig/release-notes)
[![LICENSE](https://img.shields.io/github/license/corneliusweig/release-notes.svg)](https://github.com/corneliusweig/release-notes/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/corneliusweig/release-notes)](https://goreportcard.com/report/corneliusweig/release-notes)
<!--
[![Code Coverage](https://codecov.io/gh/corneliusweig/release-notes/branch/master/graph/badge.svg)](https://codecov.io/gh/corneliusweig/release-notes)
[![Releases](https://img.shields.io/github/release-pre/corneliusweig/release-notes.svg)](https://github.com/corneliusweig/release-notes/releases)
-->

Generates a markdown changelog of merged pull requests since last release.

The script uses the GitHub API to retrieve a list of all merged pull
requests since the last release. The found pull requests are then
printed as markdown changelog with their commit summary and a link
to the pull request on GitHub.  

The idea and original implementation of this script is due to Bálint Pató
([@balopat](https://github.com/balopat)) while working on
[minikube](https://github.com/kubernetes/minikube) and
[Skaffold](https://github.com/GoogleContainerTools/skaffold).

## Examples

The binary expects two parameters:

1. The GitHub organization which your repository is part of.
2. The repository name.

For example:
```sh
./release-notes GoogleContainerTools skaffold
```

which will output
```text
Collecting pull request that were merged since the last release: v0.38.0 (2019-09-12 22:56:07 +0000 UTC)
* add github pull request template [#2894](https://github.com/googlecontainertools/skaffold/pull/2894)
* Add Jib-Gradle support for Kotlin buildscripts [#2914](https://github.com/googlecontainertools/skaffold/pull/2914)
* Add missing T.Helper() in testutil.Check* as required [#2913](https://github.com/googlecontainertools/skaffold/pull/2913)
* Move buildpacks tutorial to buildpacks example README [#2908](https://github.com/googlecontainertools/skaffold/pull/2908)
...
```

## Options

##### `--token`

Specify a personal Github Token if you are hitting a rate limit anonymously (see https://github.com/settings/tokens).

##### `--since`

The tag of the last release up to which PRs should be collected (one of `any`, `patch`, `minor`, `major`, or a valid semver). Defaults to 'patch'. 

For example:

|  |`2.3.4-alpha.1+1234`|`2.3.4-alpha.1`|`2.3.4`|`2.3.0`|`2.0.0`|
|---|---|---|---|---|---|
|`any`|true|true|true|true|true|
|`patch`|false|false|true|true|true|
|`minor`|false|false|false|true|true|
|`major`|false|false|false|false|true|


## Installation

Currently, you need a working Go compiler to build this script:

```sh
go get github.com/corneliusweig/release-notes
```
