# Rick and Morty Challenge

## Instalación

Para instalarlo (Windows, Linux, Mac) basta con seguir los siguientes pasos de la pagina oficial de Go. 

https://golang.org/doc/install

---
### Iniciar el servicio
Para arrancar el servicio solo es necesario estar posicionado en la raíz del proyecto y ejecutar el siguiente comando (Mac):
```
./main
```
Si es Windows solo es necesario ejecutar el .exe

Otra manera de arrancar el servicio es ejecutando el comando go run.
```
go run src/github.com/main.go 
```

Si solamente se desea crear el compilado se tendrá que ejecutar el comando go build. El compilado se ejecuta con los comandos antes mencionados.
```
go build src/github.com/main.go 
```

---
### Ejecutar tests
Para ejecutar los test es necesario moverse a la carpeta donde se encuentra el main y ejecutar el comando go test.
```
cd src
cd github.com
go test
```

En el caso de los test es de preferencia que no deberian tener conexión a base de datos ni ninguna conexión a servicios externos.
En el caso de este servicio se dejaron las conexiones a la API de Rick and Morty para ejemplificar los puntos que se piden. Funcionan como test de integración.

---
### Deuda Tecnica
* Agregar tambien test unitarios para todas las funciones que se utilizan y con mocks en los resultados de las llamadas HTTP.
* Agregar un caché para los resultados para no tener que re calcular todo en cada llamada (Se podría usar Redis o Memcaché)

Rick and Morty Challenge
