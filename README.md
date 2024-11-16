
# ADVENT CALENDAR API
## Для тестирования

Запустить `docker-compose`

```bash
  docker compose up -d
```

Запустить тесты

```bash
  make test
```


## Файлы, в которых прописана логика тестов

 - [E2E тестирование некоторых endpoints](https://github.com/vladvvlvlvdv/ADVENT-CALENDAR-API/blob/main/cmd/main_test.go)
 - [UNIT тестирование хеширования пароля](https://github.com/vladvvlvlvdv/ADVENT-CALENDAR-API/blob/main/pkg/utils/bcrypt_test.go)
 - [UNIT тестирование операций с датой](https://github.com/vladvvlvlvdv/ADVENT-CALENDAR-API/blob/main/pkg/utils/date_test.go)
 - [CI/CD для github actions](https://github.com/vladvvlvlvdv/ADVENT-CALENDAR-API/blob/main/.github/workflows/go.yml)
 - [Просмотреть историю github actions](https://github.com/vladvvlvlvdv/ADVENT-CALENDAR-API/actions)

