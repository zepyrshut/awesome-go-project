# Awesome Goproject

## Objetivo
Servicio REST en el que acepte ficheros en Excel. Se entiende que es en 
formato *xls* a raíz del fichero facilitado.

El serivcio debe procesar el fichero y mostrar su contenido en formato tabla
en la traza de registros *(logs)*.

## Bitácora

Normalmente cuando se inicia un proyecto, se empieza por el modelo y los detalles
como el *framework* o las vistas hay que retrasarlo todo lo posible.

En este caso, al carecer de modelo / base de datos, se empieza por los manejadores
de las peticiones http y las rutas. Se usa el router [go-chi/chi](https://github.com/go-chi/chi).

## Problemas encontrados
### Fichero XLS
El fichero proporcionado no es *xlsx*, extensión más reciente de Excel, sino *XLS*,
extensión que no es soportado por las dos librerías más actualizadas para el
tratamiento de dichos ficheros: [qax-os/excelize](https://github.com/qax-os/excelize) y
[tealeg/xlsx](https://github.com/tealeg/xlsx).

En esta [discusión](https://github.com/qax-os/excelize/issues/44) comenta el autor que
usa un formato totalmente diferente, eso en el 2017, y a [día de hoy sigue sin haber
soporte](https://github.com/qax-os/excelize/issues/1381#issuecomment-1294496456).

Con lo cual se plantea usar esta [librería](https://github.com/extrame/xls), 
[gracias al aporte del usuario *lunny*](https://github.com/qax-os/excelize/issues/44#issuecomment-386771170),
el único inconveniente es que lleva dos años sin actualizar, corriendo el riesgo de
que pueda tener vulnerabilidades, que no se analizará en para este proyecto.

Es más, este proyecto tiene otro problema, y es que a la hora de [importar el proyecto](https://github.com/extrame/xls/issues/77)
está mal etiquetada la rama *master* y al importar el módulo, coge el *commit* que
no es.
