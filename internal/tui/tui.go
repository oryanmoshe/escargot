package tui

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/go-redis/redis/v8"
	"github.com/oryanmoshe/escargot/internal/tui/layout"
	splash "github.com/oryanmoshe/escargot/internal/tui/pages"
	"github.com/rivo/tview"
)

type UI interface {
	Start()
}

type ui struct {
	app    *tview.Application
	layout layout.Layout
	// root   *tview.Primitive
}

var ctx context.Context = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var updateChan = make(chan *redis.Message)

func subForChanges() {
	topic := rdb.PSubscribe(ctx, "__keyspace*__:*")
	// Get the Channel to use
	channel := topic.Channel()
	// Itterate any messages sent on the channel

	for {
		updateChan <- (<-channel)
	}
	//for msg := range channel {
	//// Unmarshal the data into the user
	////err := u.UnmarshalBinary([]byte(msg.Payload))
	////if err != nil {
	////panic(err)
	////}

	//updateChan <- msg.String()
	//// fmt.Println(msg.Payload)
	//}
}

func updateColor(s *tview.List, app *tview.Application, color, key, t string, i int) {
	name := fmt.Sprintf("[%s]%s", color, key)
	if color == "" {
		name = key
	}

	app.QueueUpdateDraw(func() {
		s.SetItemText(i, name, t)
	})
}

func handleUpdates(ctx context.Context, l layout.Layout, s *tview.List, m map[string]map[string]interface{}, app *tview.Application) {
	for {
		select {
		case <-ctx.Done():
			// If context expires, this case is selected
			// Free up resources that may no longer be needed because of aborting the work
			// Signal all the goroutines that should stop work (use channels)
			// Usually, you would send something on channel,
			// wait for goroutines to exit and then return
			// Or, use wait groups instead of channels for synchronization
			fmt.Println("Time to return")
			l.SetStatus("Time to return")
			return
		case msg := <-updateChan:
			// This case is selected when processing finishes before the context is cancelled
			// fmt.Println(key)
			// sp.SetContent([]byte(msg.String()))
			// l.SetStatus(msg.String())
			l.SetStatus("")
			now := time.Now().Format("2006-01-02 15:04:05")
			op := msg.Payload
			key := strings.Split(msg.Channel, "__:")[1]
			l.SetStatus(fmt.Sprintf("[%s]: Operation '%s' has been triggered on key '%s'", now, op, key))
			i := m[key]["Index"].(int)
			// l.SetStatus(fmt.Sprint(i))
			go func() {
				go updateColor(s, app, "red", key, fmt.Sprint(m[key]["Type"]), i)
				time.Sleep(time.Millisecond * 300)
				// go s.SetItemText(i, fmt.Sprintf("%s", key), fmt.Sprint(m[key]["Type"]))
				// go s.SetItemText(i, fmt.Sprintf("[red]%s", key), fmt.Sprint(m[key]["Type"]))
				go updateColor(s, app, "", key, fmt.Sprint(m[key]["Type"]), i)
				time.Sleep(time.Millisecond * 300)
				go updateColor(s, app, "red", key, fmt.Sprint(m[key]["Type"]), i)
				time.Sleep(time.Millisecond * 300)
				// go s.SetItemText(i, fmt.Sprintf("%s", key), fmt.Sprint(m[key]["Type"]))
				// go s.SetItemText(i, fmt.Sprintf("[red]%s", key), fmt.Sprint(m[key]["Type"]))
				go updateColor(s, app, "", key, fmt.Sprint(m[key]["Type"]), i)
				time.Sleep(time.Millisecond * 300)
				go updateColor(s, app, "red", key, fmt.Sprint(m[key]["Type"]), i)
				// go s.SetItemText(i, fmt.Sprintf("%s", key), fmt.Sprint(m[key]["Type"]))
				time.Sleep(time.Millisecond * 300)
				// go s.SetItemText(i, fmt.Sprintf("[red]%s", key), fmt.Sprint(m[key]["Type"]))
				go updateColor(s, app, "", key, fmt.Sprint(m[key]["Type"]), i)
				// go s.SetItemText(i, fmt.Sprintf("%s", key), fmt.Sprint(m[key]["Type"]))
			}()
		}
	}
}

func New() UI {
	err := rdb.ConfigSet(ctx, "notify-keyspace-events", "KEA").Err()
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	lt := layout.New()
	splashScreen := splash.New("Splash screen")

	// listenCtx := context.Background()

	splashGrid := splashScreen.GetGrid()
	splashGrid.SetBorder(true).SetTitle("Content")
	lt.SetContent(splashGrid)

	sidebar := tview.NewList()
	sidebar.SetBorder(true).SetTitle("Sidbar")

	// files, _ := ioutil.ReadDir("./")
	// files := getRedisList()
	keys, _ := rdb.Keys(ctx, "*").Result()

	files := make(map[string]map[string]interface{})
	for i, k := range keys {
		t := rdb.Type(ctx, k).Val()
		files[k] = map[string]interface{}{
			"Type":  t,
			"Index": i,
		}
	}

	re := regexp.MustCompile("\\[\\w+\\]")

	for k, v := range files {
		// if !f.IsDir() && f.Sys {
		// sidebar.AddItem(f.Name(), "", rune(fmt.Sprintf("%d", sidebar.GetItemCount())[0]), func() {
		// sidebar.AddItem(f.Name(), fmt.Sprintf("%d", f.Size()), rune(fmt.Sprintf("%d", sidebar.GetItemCount())[0]), func() {
		sidebar.AddItem(k, fmt.Sprint(v["Type"]), rune(fmt.Sprintf("%d", sidebar.GetItemCount())[0]), func() {
			// sidebar.AddItem(f.Name(), fmt.Sprintf("%d", f.Size()), 0, func() {
			// fname := files[sidebar.GetCurrentItem()].Name()
			fname, _ := sidebar.GetItemText(sidebar.GetCurrentItem())
			fname = re.ReplaceAllString(fname, "")
			// text, _ := ioutil.ReadFile(fmt.Sprintf("./%s", fname))
			text, err := rdb.Get(ctx, fname).Result()
			splashScreen.SetContent([]byte(text))
			// rdb.Do(ctx, "JSON.SET", "omg", ".", "{\"a\": 2}")
			rdb.Set(ctx, "test", "value", 0).Err()
			if err != nil && strings.Contains(err.Error(), "WRONGTYPE") {
				text := rdb.Do(ctx, "JSON.GET", "omg").Val()
				splashScreen.SetContent([]byte(fmt.Sprint(text)))
				// panic(err)
			} else if err != nil {
				lt.SetStatus(fname)
			}
		})
		//}
		// sidebar.AddItem(f.Name(), fmt.Sprintf("%d", f.Size()), rune(strconv.Itoa(i)[0]), nil)
	}
	// firstItem, _ := sidebar.GetItemText(0)
	// text, _ := ioutil.ReadFile(fmt.Sprintf("./%s", firstItem))
	// text, _ := rdb.Get(ctx, firstItem).Result()
	// splashScreen.SetContent(text)
	// splashScreen.SetContent([]byte(text))
	sidebar.SetCurrentItem(0)
	// sidebar.ShowSecondaryText(false)

	lt.SetContent(splashGrid)
	lt.SetSidebar(sidebar)
	lt.SetFocused(sidebar)

	app := tview.NewApplication()
	go handleUpdates(ctx, lt, sidebar, files, app)
	go subForChanges()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
		}

		return event
	})

	sidebar.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			return tcell.NewEventKey(tcell.KeyLeft, 0, tcell.ModNone)
		case 'j':
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		case 'k':
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		case 'l':
			return tcell.NewEventKey(tcell.KeyRight, 0, tcell.ModNone)
		}

		return event
	})

	return ui{
		app: app,
		// root:   splashGrid,
		layout: lt,
	}
}

func (u ui) Start() {
	root := u.layout.GetLayout()
	root.SetBorder(true).SetTitle("Layout")
	u.app.SetRoot(root, true)
	u.app.SetFocus(*u.layout.GetFocused())
	if err := u.app.Run(); err != nil {
		panic(err)
	}
}
