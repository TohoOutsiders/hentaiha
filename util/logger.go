/**
 * @Author: Tomonori
 * @Date: 2019/6/3 10:57
 * @File: logger
 * @Desc:
 */
package util

import (
	"fmt"
	"github.com/kataras/golog"
	"time"
)

func init() {
	golog.SetTimeFormat("")
	golog.SetLevel("debug")
}

type ILogger interface {
	Normal(opt ...interface{})
	Info(opt ...interface{})
	Underline(opt ...interface{})
	Complate(str string)
}

type Logger struct {
}

func (l *Logger) Normal(opt ...interface{}) {
	//fmt.Println(chalk.Bold.TextStyle(str), opt)
	fmt.Printf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	golog.Println(opt...)
}

func (l *Logger) Info(opt ...interface{}) {
	//blueOnWhite := chalk.Blue.NewStyle().WithBackground(chalk.White)
	//fmt.Println(blueOnWhite.WithTextStyle(chalk.Bold).Style(str), opt)
	fmt.Printf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	golog.Info(opt...)
}

func (l *Logger) Underline(opt ...interface{}) {
	//fmt.Println(chalk.Underline.TextStyle(str), opt)
	fmt.Printf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	golog.Warn(opt...)
}

func (l *Logger) Complate(str string) {
	//lime := chalk.Green.NewStyle().WithBackground(chalk.Black).WithTextStyle(chalk.Bold).Style
	//fmt.Printf("%v\n\n\n\n", lime(str))
	fmt.Printf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	golog.Warn(str)
}
