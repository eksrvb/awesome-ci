---
layout: default
title: Commands
nav_order: 2
has_children: true
has_toc: false
permalink: /docs/commands
---

# Commands

Overview of all available transfer parameters.

These commands are used to get information in a pipeline or local build.

Some calls can also create releases or more.

## Command structure

```bash
awesome-ci <subcommand> [subcommand-option]
```

| Option          | Description                                             | requiered |
| --------------- | ------------------------------------------------------- |:---------:|
| `-version`      | Returns current used version form awesome-ci            | false     |

### Subcommands

You can find out more about the subcommands by clicking on the relevant one in the navigation.

| Subcommand                                                                         | Description                                             |
| ---------------------------------------------------------------------------------- | ------------------------------------------------------- |
| [release](https://eksrvb.github.io/awesome-ci/commands/createRelease.html)         | creates a release at GitHub or GitLab                   |
| [pr](https://eksrvb.github.io/awesome-ci/commands/getBuildInfos.html)              | prints out any git information and can manipulate these |
| [parseJSON](https://eksrvb.github.io/awesome-ci/commands/parseJSON.html)           | can parse simple JSON files                             |
| [parseYAML](https://eksrvb.github.io/awesome-ci/commands/parseYAML.html)           | can parse simple YAML files                             |
