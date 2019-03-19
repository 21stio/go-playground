package bdd

import (
	"fmt"
	"log"
	"strings"
	"regexp"
	"reflect"
	"runtime"
	"sync"

	"gopkg.in/yaml.v2"
	"github.com/fatih/color"
)

type Spec struct {
	descriptions []*Description
}

func (s Spec) init() {
	for _, d := range s.descriptions {
		d.init()
	}
}

func (s Spec) Run() {
	s.init()

	for _, d := range s.descriptions {
		d.run()
	}
}

func indentMultiLine(s string, c int) (string) {
	return strings.Repeat("  ", c) + strings.Replace(s, "\n", "\n"+strings.Repeat("  ", c), -1)
}

func indent(s string, c int) (string) {
	return strings.Repeat("  ", c) + s
}

func (s Spec) Print(v int) {
	for _, d := range s.descriptions {
		fmt.Printf(indent("describe %s\n", 1), d.name)

		for _, c := range d.contexts {
			fmt.Printf(indent("when %s\n", 2), c.name)
			if v > 0 {
				fmt.Printf("\n" + indentMultiLine(Sprint(c.Input), 3) + "\n")
			}

			for _, i := range c.its {
				var prefix string
				if i.failed {
					prefix = color.RedString("✘")
				} else {
					prefix = color.GreenString("✔")
				}

				fmt.Printf(indent(prefix+" it %s\n", 3), i.name)

				if v > 0 {
					fmt.Printf("\n" + indentMultiLine(Sprint(i.Expected), 4) + "\n")
				}
			}
		}
	}
}

func (s *Spec) Describe(name string, f func(*Description)) {
	d := &Description{
		name: name,
		f:    f,
	}

	s.descriptions = append(s.descriptions, d)
}

type Description struct {
	name     string
	f        func(*Description)
	contexts []*Context
}

func (d *Description) init() {
	d.f(d)

	for _, c := range d.contexts {
		c.init()
	}
}

func (d Description) run() {
	for _, c := range d.contexts {
		c.run()
	}
}

func (d *Description) Context(name string, f func(*Context)) {
	c := &Context{
		name: name,
		f:    f,
	}

	d.contexts = append(d.contexts, c)
}

type Context struct {
	name  string
	f     func(context *Context)
	its   []*It
	Input interface{}
}

func (c *Context) init() {
	c.f(c)
}

func (c Context) run() {
	for _, c := range c.its {
		c.run()
	}
}

func (d *Context) It(name string, f func(*It)) {
	i := &It{
		name: name,
		f:    f,
	}

	d.its = append(d.its, i)
}

type It struct {
	name        string
	f           func(*It)
	failed      bool
	errors      []string
	logs        []string
	Expected    interface{}
	So string
}

func (i *It) run() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		i.f(i)
	}()

	wg.Wait()
}

func (i *It) Errorf(format string, args ...interface{}) {
	i.failed = true
	i.errors = append(i.errors, fmt.Sprintf(format, args))
}

func (i *It) FailNow() {
	i.failed = true
	runtime.Goexit()
}

func (i *It) Log(args ...interface{}) {
	i.logs = append(i.logs, fmt.Sprint(args))
}

func Sprint(any interface{}) (string) {
	n := getName(any)

	b, err := yaml.Marshal(any)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("(?m)[\r\n]+^.*xxx_unrecognized.*$")
	res := re.ReplaceAll(b, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]*^.*: null*.*$")
	res = re.ReplaceAll(res, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: \\[].*$")
	res = re.ReplaceAll(res, []byte{})

	return fmt.Sprintf("%v%v", n, strings.Replace(string(res), "\n", "\n  ", -1))
}

func getName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
