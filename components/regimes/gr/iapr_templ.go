// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package gr

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/invopop/gobl"
	"github.com/invopop/gobl.html/components/images"
	"github.com/invopop/gobl/regimes/gr"
)

func IAPR(env *gobl.Envelope) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if hasStamps(env) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\n\t\t\t.iapr {\n\t\t\t\tbreak-inside: avoid;\n\t\t\t\tfont-size: 7pt;\n\t\t\t\tline-height: 1.6em;\n\t\t\t\tdisplay: flex;\n\t\t\t\talign-items: top;\n\t\t\t}\n\t\t\t.iapr .image {\n\t\t\t\tmargin-right: 6mm;\n\t\t\t}\n\t\t\t.iapr img {\n\t\t\t\twidth: 30mm;\n\t\t\t\theight: 30mm;\n\t\t\t}\n\t\t\t.iapr .value {\n\t\t\t\tfont-family: monospace;\n\t\t\t}\n\t\t</style> <section class=\"iapr\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if st := env.Head.GetStamp(gr.StampIAPRQR); st != nil {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"image\"><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 templ.SafeURL = templ.URL(st.Value)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = images.QR(st.Value).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if st := env.Head.GetStamp(gr.StampIAPRHash); st != nil {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><b>Σήμανση:</b> <span class=\"value\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(st.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/regimes/gr/iapr.templ`, Line: 42, Col: 36}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if st := env.Head.GetStamp(gr.StampIAPRUID); st != nil {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><b>UNIQUE ID:</b> <span class=\"value\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(st.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/regimes/gr/iapr.templ`, Line: 48, Col: 36}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if st := env.Head.GetStamp(gr.StampIAPRMark); st != nil {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><b>ΜΑΡΚ:</b> <span class=\"value\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(st.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/regimes/gr/iapr.templ`, Line: 54, Col: 36}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if st := env.Head.GetStamp(gr.StampIAPRProvider); st != nil {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><b>Υ.ΠΑ.Η.Ε.Σ:</b> <span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(st.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/regimes/gr/iapr.templ`, Line: 60, Col: 22}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func hasStamps(env *gobl.Envelope) bool {
	return env.Head.GetStamp(gr.StampIAPRQR) != nil
}