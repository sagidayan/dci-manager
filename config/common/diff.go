package common

import (
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Diff struct {
	removed  []Test
	added    []Test
	modified []Test
}

func NewDiff(removed []Test, added []Test, modified []Test) *Diff {
	return &Diff{
		added:    added,
		removed:  removed,
		modified: modified,
	}
}

func (d *Diff) Added() []Test {
	return d.added
}

func (d *Diff) Removed() []Test {
	return d.removed
}

func (d *Diff) Modified() []Test {
	return d.modified
}

func (d *Diff) Len() int {
	return len(d.removed) + len(d.modified) + len(d.added)
}

func (d *Diff) String() string {
	if d.Len() == 0 {
		return "Everything up to date. Nothing to do..."
	}
	added := color.New(color.FgGreen).Sprint("+")
	modified := color.New(color.FgYellow).Sprint("~")
	removed := color.New(color.FgRed).Sprint("-")

	t := table.NewWriter()
	t.AppendHeader(table.Row{"", "Agent", "Component", "Target Version", "App Version", "Name"})
	for _, s := range d.added {
		t.AppendRow(table.Row{added, s.AgentName(), s.ComponentName(), s.TargetVersion(), s.AppVersion(), s.Name()})
	}

	for _, s := range d.removed {
		t.AppendRow(table.Row{removed, s.AgentName(), s.ComponentName(), s.TargetVersion(), s.AppVersion(), s.Name()})
	}

	for _, s := range d.modified {
		t.AppendRow(table.Row{modified, s.AgentName(), s.ComponentName(), s.TargetVersion(), s.AppVersion(), s.Name()})
	}

	return t.Render()
}
