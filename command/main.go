// Golang ではすべてのものはパッケージに所属している必要があります．
// 最初に実行されるのは`main`パッケージです．
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// 関数は `func 「関数名」「戻り値の型名」の順で指定します．`
// 引数は「引数名」「型名」の順に記載します．
func add(args []int) int {
	var total int
	// 他言語での `forEach` に相当する書き方です．
	// `range` 句は後ろに指定したイテレータのインデックスと該当番目の要素を返すので
	// `for` 句の後の変数で受け取ります．
	// 値を使用しない場合は `_` で明示的に破棄する必要があります．
	for _, v := range args {
		total += v
	}
	return total
}

func subtruction(args []int) int {
	var result int = args[0]
	// 配列のスライスを用います．
	// python にある配列のスライスと同じ考えです．
	// Go では配は固定長なので，スライスを利用すると得られる値は
	// 「配列のスライス」となります．
	for _, v := range args[1:] {
		result -= v
	}
	return result
}

func multiplication(args []int) int {
	var result int = args[0]
	for _, v := range args[1:] {
		result *= v
	}
	return result
}

func divide(args []float64) float64 {
	var result float64 = args[0]
	for _, v := range args[1:] {
		result /= v
	}
	return result
}

func main() {
	flag.Usage = flagUsage
	addFlag := flag.NewFlagSet("add", flag.ExitOnError)
	subFlag := flag.NewFlagSet("sub", flag.ExitOnError)
	multiFlag := flag.NewFlagSet("multi", flag.ExitOnError)
	divFlag := flag.NewFlagSet("div", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "add":
		addFlag.Parse(os.Args[2:])
	case "sub":
		subFlag.Parse(os.Args[2:])
	case "mul":
		multiFlag.Parse(os.Args[2:])
	case "div":
		divFlag.Parse(os.Args[2:])
	default:
		fmt.Println("No matched subcommand.")
		return
	}

	if addFlag.Parsed() {
		nums := []int{}
		for _, v := range addFlag.Args() {
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}
		fmt.Println(add(nums))
	}

	if subFlag.Parsed() {
		nums := []int{}
		for _, v := range subFlag.Args() {
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}
		fmt.Println(subtruction(nums))

	}

	if multiFlag.Parsed() {
		nums := []int{}
		for _, v := range multiFlag.Args() {
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}
		fmt.Println(multiplication(nums))

	}

	if divFlag.Parsed() {
		nums := []float64{}
		for _, v := range divFlag.Args() {
			i, _ := strconv.ParseFloat(v, 64)
			nums = append(nums, i)
		}
		fmt.Println(divide(nums))

	}
}

func flagUsage() {
	usageText := `This is an example cli tool.

Usage:
main command [arguments]
The commands are:
add    Add all arguments.
sub    Subtructe 2nd ~ last arguments from first argument.
mul    Multiplicate all arguments.
div    Divide first argument with all the other arguments.
Use "Example [command] --help" for more infomation about a command`

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}
