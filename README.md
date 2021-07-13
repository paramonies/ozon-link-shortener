# ozon-link-shortener

[Тестовое задание](https://docs.google.com/document/d/1rrVDdDdLz_YV-B6kBsEz3N-nzqosuPvWkrk3iNG4_4I/edit) для backend-стажёра в Ozon МКК

## Содержание

1. [Общее описание](#ООписание)
1. [Стек технологий](#Стек)
1. [Установка](#Установка)
1. [Документация](#Документация)
1. [Архитектура](#Архитектура)
1. [Тестирование](#тесты)

## Описание: <a name="ООписание"></a>

Реализован сервис для сокращения и хранения http ccылок. Ссылки храняться в базе данных. Сервис предоставляет API, работающее поверх HTTP в формате JSON.

Сервис реализует следующие методы:

- `POST /short` Метод создает новую или получает существующую сокращенную ссылку
  - Request: POST /short {"url": "long-url-here"}
  - Response: {"url": "short-url-here"}
- `POST /long` Метод получения полной ссылки по сокращенной
  - Request: POST /long {"url": "short-url-here"}
  - Response: {"url": "long-url-here"}

Реализованы следующие усложнения:

- Написаны юнит тесты для уровней приложения handler, service, repository с покрытием больше 70%
- Возможность запуска приложения командой docker-compose up;
- Архитектура сервиса описана в виде [диаграммы](/docs/media/service-arc.png) и текста
- Настроена swagger документация: есть структурированное описание методов сервиса.

## Стек технологий: <a name="Стек"></a>

Golang, Gin-gonic, Postgres

## Установка: <a name="Установка"></a>

- Склонируйте проект с реппозитория GitHub
  ```
  git clone https://github.com/paramonies/ozon-link-shortener.git
  ```
- Перейдите в директорию ./ozon-link-shortener
  ```
  cd ./ozon-link-shortener
  ```
- Запустите docker-compose
  ```
  docker-compose up
  ```
- После установки сервис доступен по http на 8080 порту

## Документация: <a name="Документация"></a>

Документация к API http://localhost:8080/swagger/index.html
![docs](/docs/media/swagger_doc.png)

## Архитектура сервиса: <a name="Архитектура"></a>

![arc](/docs/media/service-arc.png)

## Тестирование: <a name="тесты"></a>

Запуск юнит-тестов:

```
docker exec -it api-server make test
```

Рассчет покрытия:

```
docker exec -it api-server make cover
```
