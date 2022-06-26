package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

// make a data structures
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

//  make a slices of the same data structures
type Todos []item

// add method it's add the task
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

// complete method

func (t *Todos) Complete(index int) error {
	ls := *t
	// if index is invalid
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil
}

// delete method
func (t *Todos) Delete(index int) error {
	ls := *t
	// if index is invalid
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Load(fileName string) error {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(fileName string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

func (t *Todos) Print() {
	// for i, item := range *t {
	// 	i++
	// 	fmt.Printf("%d - %s\n", i, item.Task)
	// }

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
