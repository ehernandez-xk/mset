# mset

Maven Settings cli

## Description

Is a CLI for  developers that helps to switch, add and save the settings.xml file, allowing to use the correct settings.xml in the desired maven project.

## Usage

```
$mset init
mset initialized

$mset add projectA settings1.xml
projectA added

$mset add projectB settings2.xml
projectB added

$mset current
projectB

$mset change projectB
projectB changed

$mset list
projectA
projectB

```

## Commands
`mset change <name>`
Changes the 

`mset add <name> <file_location>`
permite agregar un nuevo settings al catálogo de settings

`mset current`
muestra el settings que está configurado

`mset list`
muestra lista de settings disponibles, estos fueron previamente configurados
también podría imprimir la ubicación de los archivos en backup


Persistence:
Aún no se donde seria bueno persistir los archivos settings o archivos de configuración que se puedan necesitar: (que esto pueda ser configurado)

dentro del ~/.m2/.mset (default location)
ventaja es que esté en un lugar que se le pueda hacer backups sin andar buscando la carpeta
saber si maven podría borrarlo o afectar su funcionamiento
dentro ~/.mset

No se necesita otro tipo de persistencia

Formato:
Los archivos settings.xml que se agreguen pueden ser persistidos de la siguiente forma: lifeway-settings.xml, danta-settings.xml
formato aceptado por el cli <name>-settings.xml
otros formatos diferentes deben de ser ignorados, de esta forma pueden ser agregados manualmente o por medio de command "add"


New ideas
generate basic settings.xml or basic files, like [empty, basic]
mset cmd -s danta mvn clean install
check if the file is a correct xml file
see elements in the file, like servers, plugins
add elements in the file, like a new server
