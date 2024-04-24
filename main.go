package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(name, class string) (string, error) {
	damage := 5
	switch class {
	case "warrior":
		damage += randint(3, 5)
	case "mage":
		damage += randint(5, 10)
	case "healer":
		damage += randint(-3, -1)
	default:
		return "", fmt.Errorf("неизвестный класс персонажа")
	}
	return fmt.Sprintf("%s нанес урон противнику равный %d.", name, damage), nil
}

func defence(name, class string) (string, error) {
	block := 10
	switch class {
	case "warrior":
		block += randint(5, 10)
	case "mage":
		block += randint(-2, 2)
	case "healer":
		block += randint(2, 5)
	default:
		return "", fmt.Errorf("неизвестный класс персонажа")
	}
	return fmt.Sprintf("%s блокировал %d урона.", name, block), nil
}

func special(name, class string) (string, error) {
	var skillName string
	var skillPower int
	switch class {
	case "warrior":
		skillName = "Выносливость"
		skillPower = 80 + 25
	case "mage":
		skillName = "Атака"
		skillPower = 5 + 40
	case "healer":
		skillName = "Защита"
		skillPower = 10 + 30
	default:
		return "", fmt.Errorf("неизвестный класс персонажа")
	}
	return fmt.Sprintf("%s применил специальное умение `%s %d`", name, skillName, skillPower), nil
}

func startTraining(name, class string) string {
	fmt.Printf("Привет, %s!\n", name)

	switch class {
	case "warrior":
		fmt.Println("Ты Воин - отличный боец ближнего боя.")
	case "mage":
		fmt.Println("Ты Маг - превосходный укротитель стихий.")
	case "healer":
		fmt.Println("Ты Лекарь - чародей, способный исцелять раны.")
	default:
		return "Неизвестный класс персонажа."
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack, defence, special.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for {
		fmt.Print("Введи команду: ")
		_, err := fmt.Scanf("%s\n", &cmd)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch cmd {
		case "attack":
			fmt.Println(attack(name, class))
		case "defence":
			fmt.Println(defence(name, class))
		case "special":
			fmt.Println(special(name, class))
		case "skip":
			return "Тренировка окончена."
		default:
			fmt.Println("Некорректная команда. Попробуй еще раз.")
		}
	}
}

func chooseCharClass() string {
	var class string

	for {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		_, err := fmt.Scanf("%s\n", &class)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		class = strings.ToLower(class)

		if class == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if class == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if class == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		} else {
			fmt.Println("Некорректный класс персонажа. Попробуй еще раз.")
			continue
		}

		var approveChoice string
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		_, err = fmt.Scanf("%s\n", &approveChoice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		approveChoice = strings.ToLower(approveChoice)
		if approveChoice == "y" {
			break
		}
	}

	return class
}

func main() {

	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := chooseCharClass()

	fmt.Println(startTraining(name, class))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
