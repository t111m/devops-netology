package main

import "fmt"

func main() {

    for{
        fmt.Println("\nВведите команду: ")
        fmt.Println("1 - Конвертировать метры в футы")
        fmt.Println("2 - Найти наименьший элемент в заданном списке")
        fmt.Println("3 - Вывести число от 1 до 100 которые делятся на 3")

        var command int
        fmt.Scanf("%d", &command)
        if command == 1 {
                var meters float64
    fmt.Print("Введите количество метров: ")
    fmt.Scanf("%f", &meters)
                    fmt.Println("Результат: ",ConvertMetrToFeet(meters))
            } else if command == 2 {
                fmt.Println("Результат: ",FindMinInList())
            } else if command == 3 {
            fmt.Println("Результат:", From1to100div3())
            } else {
                fmt.Println("Неправильная команда")
            }
        }
}

func ConvertMetrToFeet(meters float64) float64{
    return meters * 3.048
}

func FindMinInList() int{

    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}

    for{
        fmt.Println("\nНайти в исходном списке или вновом?: ")
        fmt.Println("1 - Создать новый список")
        fmt.Println("2 - Найти минимальное в исходном списке")
        fmt.Println("Исходный список: ",x)


        var command int
        fmt.Scanf("%d",&command)

        switch command{
            case 1:
                new_x := SetList()
                x = new_x
            case 2:
                var min = x[0]
                for i := 0; i < len(x); i++ {
                    if x[i] < min {
                        min = x[i]
                    }
                }
                return min
            default:
                fmt.Println("Неправильная команда")
        }
    }
}

func SetList()[]int{

    fmt.Println("\nЗадать список: ")
    fmt.Println("Задать количество элементов: ")

    var count_list int
    fmt.Scanf("%d",&count_list)

    x := make([]int, count_list)
    for i := 0; i < count_list; i++ {
        fmt.Scanf("%d",&x[i])
    }

    return x
}

func From1to100div3()[]int{

    var k int
    k = 0
    x := make([]int,33)

    for i := 1; i < 100; i++{
        if i%3 == 0{
            x[k] = i
            k++
        }
    }

    return x
}
