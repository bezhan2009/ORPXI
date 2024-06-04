package goCmd

import (
	"bufio"
	"fmt"
	"goCmd/commands/Create"
	"goCmd/commands/Read"
	"goCmd/commands/Remove"
	"goCmd/commands/Rename"
	"goCmd/commands/Write"
	"goCmd/debug"
	"goCmd/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func changeDirectory(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("не удалось сменить директорию: %v", err)
	}
	return nil
}

func runExternalCommand(command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func GoCmd() {
	utils.SystemInformation()
	isWorking := true
	reader := bufio.NewReader(os.Stdin)

	for isWorking {
		dir, _ := os.Getwd()
		fmt.Printf("\n%s>", dir)
		commandLine, _ := reader.ReadString('\n')

		commandLine = strings.TrimSpace(commandLine)
		commandParts := strings.Fields(commandLine)

		if len(commandParts) == 0 {
			continue
		}

		command := commandParts[0]
		commandArgs := commandParts[1:]
		commandLower := strings.ToLower(command)

		if commandLower == "gohelp" {
			fmt.Println("Для получения сведений об командах наберите GOHELP")
			fmt.Println("CREATE             создает новый файл")
			fmt.Println("CLEAN              очистка экрана")
			fmt.Println("REMOVE             удаляет файл")
			fmt.Println("READ               выводит на экран содержимое файла")
			fmt.Println("WRITE              записывает данные в файл")
			fmt.Println("GOCMD              запускает ещё одну GoCMD")
			fmt.Println("CD                 смена текущего каталога")
			fmt.Println("SYSTEMINFO         вывод информации о системе")
			fmt.Println("SYSTEMGOCMD          вывод информации о GoCMD")
			fmt.Println("EXIT               Выход")
			debug.Commands(command, true)
			continue
		}

		commands := []string{"systemgocmd", "rename", "remove", "read", "write", "create", "gohelp", "exit", "gocmd", "clean", "cd"}

		isValid := utils.ValidCommand(commandLower, commands)

		if !isValid {
			fullCommand := append([]string{command}, commandArgs...)
			err := runExternalCommand(fullCommand)
			if err != nil {
				fullPath := filepath.Join(dir, command)
				fullCommand[0] = fullPath
				err = runExternalCommand(fullCommand)
				if err != nil {
					fmt.Printf("Ошибка при запуске команды '%s': %v\n", commandLine, err)
				}
			}
			continue
		}

		switch commandLower {
		case "systemgocmd":
			utils.SystemInformation()
		case "gocmd":
			GoCmd()

		case "exit":
			isWorking = false

		case "create":
			err, name := Create.File()
			if err != nil {
				fmt.Println(err)
				debug.Commands(command, false)
			} else {
				fmt.Printf("Файл %s успешно создан!!!\n", name)
				fmt.Printf("Директория нового файла: %s\n", filepath.Join(dir, name))
				debug.Commands(command, true)
			}

		case "write":
			if len(commandArgs) < 2 {
				fmt.Println("Использование: write <файл> <данные>")
				continue
			}
			nameFileForWrite := commandArgs[0]
			data := strings.Join(commandArgs[1:], " ")

			if nameFileForWrite == "debug.txt" {
				debug.Commands(command, false)
				fmt.Println("PermissionDenied: You cannot write, delete or create a debug.txt file")
				continue
			}

			errWriting := Write.File(nameFileForWrite, data+"\n")
			if errWriting != nil {
				debug.Commands(command, false)
				fmt.Println(errWriting)
			} else {
				debug.Commands(command, true)
				fmt.Printf("Мы успешно записали данные в файл %s\n", nameFileForWrite)
			}

		case "read":
			if len(commandArgs) < 1 {
				fmt.Println("Использование: read <файл>")
				continue
			}
			nameFileForRead := commandArgs[0]

			dataRead, errReading := Read.File(nameFileForRead)
			if errReading != nil {
				debug.Commands(command, false)
				fmt.Println(errReading)
			} else {
				debug.Commands(command, true)
				_, errWrite := os.Stdout.Write(dataRead)
				if errWrite != nil {
					fmt.Println(errWrite)
				}
			}

		case "remove":
			err, name := Remove.File()
			if err != nil {
				debug.Commands(command, false)
				fmt.Println(err)
			} else {
				debug.Commands(command, true)
				fmt.Printf("Файл %s успешно удален!!!\n", name)
			}

		case "rename":
			errRename := Rename.Rename()
			if errRename != nil {
				debug.Commands(command, false)
				fmt.Println(errRename)
			} else {
				debug.Commands(command, true)
			}

		case "clean":
			clearScreen()

		case "cd":
			if len(commandArgs) == 0 {
				fmt.Println("Введите путь")
			} else {
				err := changeDirectory(commandArgs[0])
				if err != nil {
					fmt.Println(err)
				}
			}

		default:
			validCommand := utils.ValidCommand(commandLower, commands)
			if !validCommand {
				fmt.Printf("'%s' не является внутренней или внешней командой,\nисполняемой программой или пакетным файлом.\n", commandLine)
			}
		}
	}
}
