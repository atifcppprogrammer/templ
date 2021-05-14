// Code generated by templ DO NOT EDIT.

package ifelse

import "github.com/a-h/templ"
import "context"
import "io"

func render(d data) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		if d.IsTrue() {
			_, err = io.WriteString(w, templ.EscapeString("True"))
			if err != nil {
				return err
			}
		} else {
			_, err = io.WriteString(w, templ.EscapeString("False"))
			if err != nil {
				return err
			}
		}
		return err
	})
}

