# tabledumper


## Описание
Модуль для многопоточного дампа таблиц в указанной базе данных mysql


## Использование
Ниже фрагмент кода с помощью которого можно использовать данный модуль

```go
package main

import (
    "github.com/riasloff/tabledumper"
)

func main() {
    cfg := tabledumper.Config{
        Host:     "11.22.33.44",
        User:     "golang",
        Password: "golang-password",
        DbName:   "data",
        Tables:   []string{"myTable", "dog_data", "love_data", "mock_data", "post_data"},
    }

    tabledumper.Start(cfg)
}
```