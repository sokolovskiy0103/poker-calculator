# Poker Score Calculator

Це кросплатформенний GUI-додаток на базі бібліотеки [Fyne](https://fyne.io/), який допомагає відслідковувати та
обчислювати результати карткової гри в покер (насправді це не покер, але в дитинстві це так називали 😀).

## Вимоги

- Go >= 1.16
- Бібліотека Fyne v2: `go get fyne.io/fyne/v2`
- Fyne SDK: для кросплатформеної компіляції (встановлюється за допомогою інструкцій Fyne)
    - Fyne для мобільних платформ потребує Android SDK або Xcode для iOS.

## Встановлення

1. Клонуйте репозиторій або скопіюйте код до своєї локальної машини:
    ```bash
    git clone https://github.com/username/poker-score-calculator.git
    cd poker-score-calculator
    ```

2. Завантажте та встановіть залежності:
    ```bash
    go mod tidy
    ```

3. Запустіть додаток:
    ```bash
    go run main.go
    ```

## Компіляція за допомогою Fyne

### Загальні вимоги

Для компіляції з використанням Fyne вам потрібно встановити інструмент Fyne CLI. Це можна зробити за допомогою команди:

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```

### Windows (64-біт)

```bash
fyne package -os windows
```

### Android
```bash
fyne package -os android -appID {appID}
```


