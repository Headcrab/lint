# Интеграция с golangci-lint

## Module Plugin System (Windows, Linux, macOS)

### 1. Опубликуй модуль
```bash
git remote add origin https://github.com/Headcrab/unusedintf.git
git tag v1.0.0
git push origin main --tags
```

### 2. Собери кастомный golangci-lint
```bash
golangci-lint custom --custom-config .custom-gcl.yml
```

### 3. Запусти
```bash
# Windows
.\custom-gcl.exe run

# Linux/macOS  
./custom-gcl run
```

## Локальная разработка

Для тестирования без публикации используй `path: "./"` в `.custom-gcl.yml`

## Параметры

- `skipGenerics: true` - пропускать generic интерфейсы

## Альтернатива: форк golangci-lint

1. Форкни https://github.com/golangci/golangci-lint
2. Добавь в `cmd/golangci-lint/plugins.go`:
   ```go
   import _ "github.com/yourusername/unusedintf/pkg/unusedintf"
   ```
3. Зарегистрируй в `pkg/lint/lintersdb/builder_linter.go`:
   ```go
   b.Add(&unusedintf.Analyzer).WithSince("v1.XX.0")
   ``` 