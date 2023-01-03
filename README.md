# Awesome Go project

## Objetivo
Servicio REST en el que acepte ficheros en Excel. Se entiende que es en 
formato *xls* a raíz del fichero facilitado.

El serivcio debe procesar el fichero y mostrar su contenido en formato tabla
en la traza de registros *(logs)*.

## Bitácora
Normalmente cuando se inicia un proyecto, se empieza por el modelo y los detalles
como el *framework* o las vistas hay que retrasarlo todo lo posible.

En este caso, al carecer de base de datos, se empieza por los manejadores de las 
peticiones http y las rutas. Se usa el router [go-chi/chi](https://github.com/go-chi/chi).

El router go-chi/chi es muy ligero, con muy buena documentación y fácil de empezar
a trabajar con él.

La arquitectura de la aplicación se trata de separar en dos paquetes principales,
lo que es la aplicación *(internal)* y la inicialización *(cmd)*. Dentro de *internal*
se separa todo por funciones:

- *Handlers*: Manejadores de las peticiones http.
- *Router*: Rutas de la aplicación.
- *Configuration*: Ahí van los detalles como versión, puerto, log y detalles que son
configurables mediante variables de entorno.
- *Core*: Lógica de negocio, en este caso su única funcion es imprimir un fichero Excel.

## Problemas encontrados
### Fichero XLS
El fichero proporcionado no es *xlsx*, extensión más reciente de Excel, sino *xls*,
extensión que no es soportado por las dos librerías más actualizadas para el
tratamiento de dichos ficheros: [qax-os/excelize](https://github.com/qax-os/excelize) y
[tealeg/xlsx](https://github.com/tealeg/xlsx).

En esta [discusión](https://github.com/qax-os/excelize/issues/44) comenta el autor que
usa un formato totalmente diferente, eso en el 2017, y a [día de hoy sigue sin haber
soporte](https://github.com/qax-os/excelize/issues/1381#issuecomment-1294496456).

Con lo cual se plantea usar esta [librería](https://github.com/extrame/xls), 
[gracias al aporte del usuario *lunny*](https://github.com/qax-os/excelize/issues/44#issuecomment-386771170),
el único inconveniente es que lleva dos años sin actualizar, corriendo el riesgo de
que pueda tener vulnerabilidades, que no se analizarán para este proyecto.

Es más, este proyecto tiene otro problema, y es que a la hora de [importar el proyecto](https://github.com/extrame/xls/issues/77)
está mal etiquetada la rama *master* y al importar el módulo, coge el *commit* que
no es.

### En caso de que el fichero sea XLSX
La librería usada no soporta la extensión más moderna, en caso de querer usarse ambas,
habría que pasar una comprobación de la extensión del fichero y usar una u otra librería.

Como buenas prácticas de código habría que cambiar el método *PrintXlsFile* por *PrintFile*
y elevarlo hacia una interfaz, y de ahí hacer ambas implementaciones para cada caso.
## Otras consideraciones
### El paquete Slog
Es un paquete experimental, novedad hace unos meses para realizar una traza del programa
de forma estructurada. Slog viene de *structured logging*. Puedes ver más información
en el [siguiente vídeo](https://www.youtube.com/watch?v=gd_Vyb5vEw0).

## Conclusiones
Soy plenamente consciente de cómo está construida la aplicación, en multicapas con lo que
se extiende *ligeramente* el tiempo de desarrollo, pero dando ventajas como abierto a la
extensión y cerrado a la modificación en la medida de lo posible. Podría haber estado todo
en el paquete *main* y hubiese funcionado sin problema.

Como sólo se disponía de dos horas para realizarse, en mi caso he necesitado un poco más, 
he tenido que realizar algunos sacrificios para ganar tiempo como prescindir el hecho de 
hacer pruebas unitarias, separar el *front-end* del *back-end*, pues se ha usado el motor
de plantillas que integra la librería estándar en su formato más básico, sin caché ni
fragmentos.

La mayor dificultad que he tenido que afrontar han sido cómo tratar un fichero Excel,
pues nunca he dado con el caso, pero además de ser un formato antiguo se han complicado un
poco más las cosas (XLS en lugar de XLSX) y la documentación inexistente, pero 
afortunadamente gran parte de la magia está escondida en *issues y pull requests* y usando 
el buscador. Navegando por esa zona he dado con las soluciones.

En definitiva, una gran prueba, que a priori parece sencilla (subir fichero, rutas, http...), 
pero con un pequeño desafío con el tratamiento de ficheros, desconzoco si ha sido de forma
intencionada, pero ha estado muy interesante.
