# REST API сервис на Go
***
#### ФИО
Хабибуллин Руслан Марсович
***
#### Тестовое задание GO
Сущность "Мебель"
:   Название мебели
    Производитель
    Высота
    Ширина
    Длина

***
#### Описание проекта
##### 1. Был создан REST API серивис с следующими HTTP запросами

  HTTP запрос | Принимоемое значение | Возвращаемое значение | Совершаемое действие
 :------------|:---------------------|:-----------------------:|----
 GET    | Ничего не принимает | Все созданные объекты сущности хранимые в БД| Отображаем все сущности из БД в формате JSON
 GET    | ID объекта сущности | Сущность с данным ID | Отображаем сущность с данным ID в формате JSON
 POST | Объект сущности | ID объекта сущности | Создаем новый объект сущности 
 PUT | ID объекта сущности и объект сущности на который хотим заменить объект с данным ID | Статус выполненной операции(OK/Error) | Совершаем изменения в объекте сущности, если введены все поля для изменений
 PATCH | ID объекта сущности и объект сущности на который хотим заменить объект с данным ID | ID объекта сущности и объект сущности на который хотим заменить объект с данным ID | Совершаем изменения в объекте сущности, если введено хотя бы одно поле для изменений
 DELETE | ID объекта сущности | Статус выполненной операции(OK/Error) | Удаляем объект сущности с данным ID

##### 2. Основные элементы проекта и их содержание

cmd
 : основная папка, в которой происходит запуск проекта

internal
 : папка в которой содержатся внутренний процессы сервиса (работа с клиентом, сервером и базой данных)

 configs
 : папка в которой содержится конфигурация для подключения к docker и базе данных postgres

 schema
 : папка, в которой содержится описание базы данных и функционал для создания и удаления содержимого всей базы данных

 furnitures.go
 : файл в котором содержатся необходимые структуры для работы с базой данных (в данном случае сущность "Мебель" и структура для обновления этой сущности)

 server.go
 : файл в котором содержатся необходимые методы для создания сервера

##### 3. Проект создан с помощью фреймворка gin.
##### 4. Для создания базы данных использовались docker и postgres(встроенный в docker)
***
#### Подготовительные действия (установки, настройки и т.д.) для успешной работы

Для успешной работы сервиса на вашем компьютере необходимо следующее:
1. Поменять названия всех подключаемых пакетов приложения(те которые непосредственно содержатся в приложении) для своего компьютера. Например в файле main.go:
 ```
    "log"
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/handler"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/repository"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/service"

	"github.com/spf13/viper"
 ```
 Поменять на: 
 ```
    "log"
	restapimediasoft "<your_path>/rest_api_mediasoft/rest_api-Mediasoft"
	"<your_path>/rest_api_mediasoft/rest_api-Mediasoft/internal/handler"
	"<your_path>/rest_api_mediasoft/rest_api-Mediasoft/internal/repository"
	"<your_path>/rest_api_mediasoft/rest_api-Mediasoft/internal/service"

	"github.com/spf13/viper"
 ```
2. Необходимо подключить внешние(сторонние) библиотеки. Список команд для их подключения:
```
go get init github.com/spf13/viper
go get init github.com/gin-gonic/gin
go get init github.com/sirupsen/logrus 
go get init github.com/jmoiron/sqlx

```
3. Для создания базы данных использовался docker, поэтому необходимо, чтобы он был установлен, также вот список команд для создания контейнера и миграций: 

-Создание контейнера и базы данных
```
docker run --name=restapi_db -e POSTGRES_PASSWORD=12345 -p 5436:5432 -d --rm postgres
```
Если введенные настройки вам не удовлетворяют(например порт 5436 уже занят), то можете ввести свои настройки и поменять их в файле конфигураций config/configs.yaml

-Создание миграций (для контроля версий БД и в принципе создания файла в котором будем писать)
```
migrate create -ext sql -dir <путь до проекта>/schema -seq init
```
В случае, если у вас не установлена утилита migrate, то введите в терминал следующее:
```
scoop install migrate
```
-После вышеперечисленных действий появятся файлы(000001_init.down.sql и 000001_init.up.sql) в папке schema в них необходимо ввести соответствующий код из проекта.
-Как только вы ввели соответсвующие изменения в проект необходимо вернуться в терминал и ввести следующее(применить миграцию):
```
migrate -path <ваш путь до проекта>/schema 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up
```
***
#### Информация о доступах
Помимо вышеперечисленных махинаций с созданием базы данных (присвоение названия и пароля базе данных) других нет
***
#### Описание как запустить ваш проект
Необходимо зайти на сайт [postman](https://www.postman.com/).
 Войти и зарегестрироваться там.
 Перейти на страницу "Send an API request".
 Вводить соответствующие команды для запросов:
 localhost:8000/api/lists для запросов GET, POST
 localhost:8000/api/lists/<нужный id> для запросов GET, PUT, PATCH, DELETE

***
#### *В случае, если не получится запустить проект
Вот [видео](https://www.youtube.com/watch?v=egnFzrOKO6A) с демонстрацией работы сервиса.
