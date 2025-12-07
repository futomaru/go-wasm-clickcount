//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")
	countEl := doc.Call("getElementById", "count")
	btn := doc.Call("getElementById", "click-btn")
	count := 0

	update := func() { countEl.Set("textContent", count) }

	click := js.FuncOf(func(js.Value, []js.Value) interface{} {
		count++
		update()
		return nil
	})
	defer click.Release()

	update()
	btn.Call("addEventListener", "click", click)
	select {}
}
