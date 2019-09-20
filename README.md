# release-notes

Generate a markdown changelog of merged pull requests since last release.

The script uses the public GitHub API to retrieve a list of all
closed pull requests since the last release. These pull requests
are then printed as markdown changelog with their commit summary
and a link to the pull request on GitHub.  

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

## Installation

Currently, you need a working Go compiler to build this script:

```sh
go get github.com/corneliusweig/release-notes
```
