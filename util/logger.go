/**
 * @Author: Tomonori
 * @Date: 2019/6/3 10:57
 * @File: logger
 * @Desc:
 */
package util

import (
	"fmt"
	"github.com/ttacon/chalk"
)

type ILogger interface {
	Normal(str string, opt ...interface{})
	Info(str string, opt ...interface{})
	Underline(str string, opt ...interface{})
	Complate(str string)
}

type Logger struct {
}

func (l *Logger) Normal(str string, opt ...interface{}) {
	fmt.Println(chalk.Bold.TextStyle(str), opt)
}

func (l *Logger) Info(str string, opt ...interface{}) {
	blueOnWhite := chalk.Blue.NewStyle().WithBackground(chalk.White)
	fmt.Println(blueOnWhite.WithTextStyle(chalk.Bold).Style(str), opt)
}

func (l *Logger) Underline(str string, opt ...interface{}) {
	fmt.Println(chalk.Underline.TextStyle(str), opt)
}

func (l *Logger) Complate(str string) {
	lime := chalk.Green.NewStyle().WithBackground(chalk.Black).WithTextStyle(chalk.Bold).Style
	fmt.Printf("%v\n\n\n\n", lime(str))
}
