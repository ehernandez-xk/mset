# mset

Maven Settings cli

## Description

Is a simple cli tool that helps to easy save, set and switch the settings.xml file inside the ~/.m2 directory

## Usage

```sh
$mset init
mset initialized

$mset add projectA settings1.xml
projectA added

$mset add projectB settings2.xml
projectB added

$mset current
projectB

$mset set projectB
added projectB to current

$mset list
projectA
projectB

```

## Install

- See the Makefile to build the binary and then copy it to a bin directory inside $PATH

## Commands

`mset init`
Create an empty catalog /Users/ehernandez/.m2/.mset to store the settings.xml files

`mset set <name>`
Set the `<name>` as current, copying the settings.xml in ~.m2/

`mset add <name> <file_location>`
Add a new entry in the catalog, copying <file> in a new settings.xml

`mset current`
Show the name of the current settings.xml file

`mset list`
List files available in the catalog

`remove <name>`
Remove the `<name>` from the catalog

## Notice

- By default the catalog is created in `$HOME/.m2/.mset` but this can be changed using an EnvVar `MSET_CATALOG_PATH`
- New entries with `add` command are stored with a suffix `-settings.xml` in the catalog
- Name entries added with `add` must be alphanumerica and with `-` regexp: (`^[a-zA-Z0-9-]*$`)
- The command `list` only sees files with the suffix `-settings.xml` other are ignored
- Set a new file with `set` command copies the related settings.xml file to `$HOME/.m2/settings.xml`
- The default Maven directory is `$HOME/.m2/` but this can be changed using an EnvVar `MSET_MAVEN_PATH`
