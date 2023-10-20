# Тестовое задание.

## Запуск проекта (backend)

### Способ 1: С использованием Docker (с Docker hub)

1. Скачайте контейнер Docker командой:

    ```
    docker pull happydayaway/task_img:latest
    ```

2. Запустите контейнер командой:

    ```
    docker run -d -p 80:8080 --rm --name task_cont happydayaway/task_img
    ```

### Способ 2: С использованием Docker (с Github)

1. Скачайте проект с GitHub командой:

    ```
    git clone https://github.com/karrrakotov/test_go.git
    ```

2. Откройте консоль, перейдите в корневую папку проекта и выполните команду:

    ```
    make build
    ```

3. Затем запустите проект командой:

    ```
    make run
    ```
4. Если необходимо остановить работу проекта, впишите команду:

    ```
    make stop-cont
    ```
### Способ 3: С использованием исходного кода
1. Скачайте проект с GitHub командой:

    ```
    git clone https://github.com/karrrakotov/test_go.git
    ```

2. Откройте консоль, перейдите в корневую папку проекта и выполните команду:

    ```
    go build cmd/main/main.go
    ```
3. Затем запустите проект командой:

    ```
    ./main.go
    ```
