package parser

import (
	"fmt"
	"github.com/aliffatulmf/tsch-compose/internal/model"
	"github.com/aliffatulmf/tsch-compose/internal/schedule"
	"os/exec"
	"os/user"
	"slices"
	"strings"
)

type Args []string

type Parse struct {
	compose *model.Compose
	CMD     map[string]Args
}

func NewParser(compose *model.Compose) *Parse {
	return &Parse{compose: compose}
}

func (p *Parse) parseModifier(task model.Task) Args {
	switch task.ScheduleType {
	case schedule.Minute:
	case schedule.Hourly:
	case schedule.Daily:
	case schedule.Weekly:
	case schedule.Monthly:
		switch task.Modifier.(type) {
		case string:
			if task.Modifier.(string) == "LASTDAY" && task.Month != "" {
				return Args{"/MO", task.Modifier.(string), "/M", task.Month}
			}

			if slices.Contains([]string{"FIRST", "SECOND", "THIRD", "FOURTH"}, task.Modifier.(string)) && task.Day != nil {
				return Args{"/MO", task.Modifier.(string), "/D", task.Day.(string)}
			}
		case int:
			if task.Modifier.(int) != 0 {
				return Args{"/MO", fmt.Sprintf("%d", task.Modifier.(int))}
			}
		}
	}
	return nil
}

func (p *Parse) parseTime(task model.Task) Args {
	if task.StartTime != (model.Time{}) {
		return Args{"/ST", task.StartTime.Clock()}
	}

	if task.EndTime != (model.Time{}) {
		return Args{"/ET", task.EndTime.Clock()}
	}

	return nil
}

func (p *Parse) parse() {
	p.CMD = make(map[string]Args)

	for name, task := range p.compose.Tasks {
		p.CMD[name] = Args{
			"/SC", string(task.ScheduleType),
			"/TN", fmt.Sprintf(`"%s\%s"`, p.compose.Namespace, task.Name),
			"/TR", fmt.Sprintf(`"%s"`, task.Run),
		}

		p.parseTime(task)

		if task.Modifier != nil {
			p.CMD[name] = append(p.CMD[name], p.parseModifier(task)...)
		}
	}
}

func (p *Parse) Parse(action string) map[string]Args {
	p.parse()

	curUser, _ := user.Current()

	for name, args := range p.CMD {
		switch action {
		case "create":
			args = slices.Insert(args, 0, "SCHTASKS", "/CREATE", "/F", "/RU", strings.Split(curUser.Username, "\\")[1])
			//args = slices.Insert(args, 0, "powershell.exe", "-Command", "SCHTASKS", "/CREATE")
		case "delete":
			args = slices.Insert(args, 0, "SCHTASKS", "/DELETE")
			//args = slices.Insert(args, 0, "powershell.exe", "-Command", "SCHTASKS", "/DELETE")
		}
		p.CMD[name] = args
	}

	return p.CMD
}

func (p *Parse) ToCMD() map[string]*exec.Cmd {
	var cmds = make(map[string]*exec.Cmd)

	for name, args := range p.CMD {
		arg := fmt.Sprintf(`%s`, strings.Join(args, " "))
		cmds[name] = exec.Command("powershell.exe", "-Command", arg)
	}

	return cmds
}
