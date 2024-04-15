# BDD test using Godog

## Шаг 1: Установка Godog
```bash
go get github.com/cucumber/godog/cmd/godog
```

## Шаг 2: Создание теста 
Создайте каталог `features` в корне проекта.

```bash
mkdir features
```

Создайте файл с именем `disk.feature` в каталоге `features`.

```gherkin
Feature: create disk
  As a user
  I want to create a disk

  Scenario: create disk
      Given an image with family "ubuntu-2204-lts" in the "standard-images" folder
      When I create a disk from it with name "my-disk" in the "b1gbpo1c8qkicn81mfok" folder
      Then I should see the disk in the folder
      Then I want to delete the disk
```

## Шаг 3: Создание тестового файла

Создайте каталог `steps` в каталоге `features`.

```bash
mkdir features/steps
```

Создайте файл с именем `disk_test.go` в каталоге `features/steps`.

```go
package steps

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		// ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

```

## Шаг 4: Запуск теста

Запустите тест.

```bash
go test -v ./features/steps
```
Вы увидите следующий вывод:


```
=== RUN   TestFeatures
Feature: create disk
  As a user
  I want to create a disk
=== RUN   TestFeatures/create_disk

  Scenario: create disk                                                                   # ../disk.feature:5
    Given an image with family "ubuntu-2204-lts" in the "standard-images" folder
    When I create a disk from it with name "my-disk" in the "b1gbpo1c8qkicn81mfok" folder
    Then I should see the disk in the folder

1 scenarios (1 undefined)
3 steps (3 undefined)
323.792µs

You can implement step definitions for undefined steps with these snippets:

func anImageWithFamilyInTheFolder(arg1, arg2 string) error {
        return godog.ErrPending
}

func iCreateADiskFromItWithNameInTheFolder(arg1, arg2 string) error {
        return godog.ErrPending
}

func iShouldSeeTheDiskInTheFolder() error {
        return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
        ctx.Step(`^an image with family "([^"]*)" in the "([^"]*)" folder$`, anImageWithFamilyInTheFolder)
        ctx.Step(`^I create a disk from it with name "([^"]*)" in the "([^"]*)" folder$`, iCreateADiskFromItWithNameInTheFolder)
        ctx.Step(`^I should see the disk in the folder$`, iShouldSeeTheDiskInTheFolder)
}

--- PASS: TestFeatures (0.00s)
    --- PASS: TestFeatures/create_disk (0.00s)
PASS
ok      bdd-godog/features/steps        (cached)
```

## Шаг 5: Реализация шагов

Теперь нам нужно реализовать предложенные шаги.
Полный код шагов можно найти в файле `features/steps/disk_test.go`.

## Шаг 6: Запуск теста

Запустите тест снова.

```bash
go test -v ./features/steps
```