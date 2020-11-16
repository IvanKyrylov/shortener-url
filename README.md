# Create URL shortener server
Implement an HTTP server that can generate shortened URLs.

1.Ability to create new shortened URLs using API request.

2.The requests to shortened URLs should be redirected to their original URL (status 302) or return 404 for unknown URLs.

3.All of the implemented HTTP handlers should have unit tests.

4.(optional) All shortened URLs should be persisted locally to a storage (MySQL, PostgreSQL, Redis, CSV, etc).

5.(optional) The redirect requests should be cached in memory for a certain amount of time.


## Выполнено 

1.Ability to create new shortened URLs using API request.

2.The requests to shortened URLs should be redirected to their original URL (status 302) or return 404 for unknown URLs.

3.All of the implemented HTTP handlers should have unit tests.

4.(optional) All shortened URLs should be persisted locally to a storage (MySQL, PostgreSQL, Redis, CSV, etc).

## Работает некорректно 
5.(optional) The redirect requests should be cached in memory for a certain amount of time.
Хранит в памяти сервера любое указанное время. Для проверки код
`mux.HandleFunc("/", h.redirect)
// mux.HandleFunc("/", cached("60s", h.redirect))`
замениете на
`mux.HandleFunc("/", cached("60s", h.redirect))`
в файле handler.go, 27 строка


# Описание 
Основной url:
`http://localhost:8080/shortener`
Вот заготовка post запроса для получения сокрощенного url:
```json{
    "url":"https://github.com/IvanKyrylov/shortener-url"
}
```
Redirect:
`http://localhost:8080/{key}`
key получаеться из ответа `http://localhost:8080/shortener`
Url пернаправления 

info:
`http://localhost:8080/info/{key}`
Получение информации о ссылке


#Запуск
Была использована база PostgreSQL.
Для запуска 
`go run ./cmd/shortener/main.go`
или 
``go build ./cmd/shortener/
./shortener.exe``


# Эпитафия
Извиняюсь что без Docker или make файлов, за ужасный код, отсутствие комментариев, отвратительную архитектуру, за readme на славянском. Но могу пока что, только так. Спасибо за внимание. Слава Омниссии!
