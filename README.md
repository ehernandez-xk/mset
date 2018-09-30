# mset

Maven Settings cli

## Description

Is a simple cli tool that helps to easy save, set and switch the settings.xml file inside the ~/.m2 directory

## Usage

you can run mset command in any path

```sh
# initialize the catalog
$mset init

# add or save one settings.xml file to the catalog
$mset add projectA settings1.xml

# add another file to the catalog
$mset add projectB settings2.xml

# see what file is currently used
$mset current
projectB

# change to use the settings.xml called projectB
$mset set projectB
added projectB to current

# see all your files stored in the catalog
$mset list
projectA
projectB (current)

```

## Install

- See Makefile to build and then copy the binary to a /bin directory inside $PATH

## Commands

`mset init`
Initialize the catalog to store your settings.xml files, this is done under `$HOME/.m2/.mset`. Also stores the current `.m2/settings.xml`

`mset set <name>`
Set the `<name>` as current, this copies settings.xml file related to `<name>` to `~.m2/settings.xml`

`mset add <name> <file>`
Add a new settins.xml file in the catalog

`mset current`
Show the name of the current settings.xml file

`mset list`
List all available setting.xml files in the catalog

`remove <name>`
Remove the settings.xml fiile from the catalog

## Notice

- By default the catalog is created in `$HOME/.m2/.mset` but this can be changed using an EnvVar `MSET_CATALOG_PATH`
- New entries with `add` command are stored with a suffix `-settings.xml` in the catalog
- Name entries added with `add` must be alphanumerica and with `-` regexp: (`^[a-zA-Z0-9-]*$`)
- The command `list` only sees files with the suffix `-settings.xml` other are ignored
- Set a new file with `set` command copies the related settings.xml file to `$HOME/.m2/settings.xml`
- The default Maven directory is `$HOME/.m2/` but this can be changed using an EnvVar `MSET_MAVEN_PATH`
