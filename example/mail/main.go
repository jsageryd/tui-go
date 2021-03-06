package main

import "github.com/marcusolsson/tui-go"

type mail struct {
	from    string
	subject string
	date    string
	body    string
}

var mails = []mail{
	{
		from:    "John Doe <john@doe.com>",
		subject: "Vacation pictures",
		date:    "Yesterday",
		body: `
Hey,

Where can I find the pictures from the diving trip?

Cheers,
John`,
	},
	{
		from:    "Jane Doe <jane@doe.com>",
		subject: "Meeting notes",
		date:    "Yesterday",
		body: `
Here are the notes from today's meeting.

/Jane`,
	},
}

func main() {
	var (
		from    = tui.NewLabel("")
		subject = tui.NewLabel("")
		date    = tui.NewLabel("")
		body    = tui.NewLabel("")
	)

	info := tui.NewGrid(0, 0)
	info.SetSizePolicy(tui.Expanding, tui.Minimum)
	info.AppendRow(tui.NewLabel("From:"), from)
	info.AppendRow(tui.NewLabel("Subject:"), subject)
	info.AppendRow(tui.NewLabel("Date:"), date)

	mail := tui.NewVBox(info, body)
	mail.SetSizePolicy(tui.Expanding, tui.Expanding)

	inbox := tui.NewTable(0, 0)
	inbox.SetSizePolicy(tui.Expanding, tui.Minimum)
	inbox.SetColumnStretch(0, 3)
	inbox.SetColumnStretch(1, 2)
	inbox.SetColumnStretch(2, 1)

	for _, m := range mails {
		inbox.AppendRow(
			tui.NewLabel(m.subject),
			tui.NewLabel(m.from),
			tui.NewLabel(m.date),
		)
	}

	inbox.OnSelectionChanged(func(t *tui.Table) {
		m := mails[t.Selected()]
		from.SetText(m.from)
		subject.SetText(m.subject)
		date.SetText(m.date)
		body.SetText(m.body)
	})

	// Select first mail on startup.
	inbox.Select(0)

	root := tui.NewVBox(inbox, tui.NewLabel(""), mail)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	if err := tui.New(root).Run(); err != nil {
		panic(err)
	}
}
