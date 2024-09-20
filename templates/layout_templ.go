// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func MainLayout() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@2.0.2\"></script><script src=\"https://cdn.jsdelivr.net/gh/gnat/css-scope-inline/script.js\"></script><script src=\"https://cdn.jsdelivr.net/gh/gnat/surreal@main/surreal.js\"></script><title>OCR Web UI</title></head><body><div><canvas></canvas><div id=\"img-container\"><img id=\"img\" src=\"https://i.pinimg.com/736x/72/eb/36/72eb365f16469ea0e6093d29e42c5924.jpg\"></div><style>\n\t\t\t\t\tme {\n\t\t\t\t\t\tbackground: red;\n\t\t\t\t\t}\n\n\t\t\t\t\tme img {\n\t\t\t\t\t\tz-index: 0;\n\t\t\t\t\t\tposition: relative;\n\t\t\t\t\t}\n\n\t\t\t\t\tme canvas {\n\t\t\t\t\t\tz-index: 20;\n\t\t\t\t\t\tposition: absolute;\n\t\t\t\t\t}\n\t\t\t\t</style><script>\n\t\t\t\t\tconst pos = { start: { x: null, y: null }, end: { x: null, y: null } }\n\t\t\t\t\tlet canvas;\n\t\t\t\t\tlet ctx;\n\t\t\t\t\twindow.onload = (e) => {\n\t\t\t\t\t\tcanvas = me(\"canvas\")\n\t\t\t\t\t\tctx = canvas.getContext(\"2d\")\n\t\t\t\t\t\tcalculateCanvasBound()\n\t\t\t\t\t\t\n\t\t\t\t\t\tcanvas.onmousedown = function (e) {\n\t\t\t\t\t\t\tclearPosForm()\n\t\t\t\t\t\t\tpos.start = getPosition(e)\n\t\t\t\t\t\t}\n\t\t\t\t\t\tcanvas.onmousemove = function (e) {\n\t\t\t\t\t\t\tconst { x, y } = getPosition(e)\n\t\t\t\t\t\t\tdrawRect(ctx, pos.start.x, pos.start.y, x, y)\n\t\t\t\t\t\t}\n\t\t\t\t\t\tcanvas.onmouseup = function (e) {\n\t\t\t\t\t\t\tpos.end = getPosition(e)\n\t\t\t\t\t\t\tdrawRect(ctx, pos.start.x, pos.start.y, pos.end.x, pos.end.y)\n\t\t\t\t\t\t\tfillPosForm(pos.start.x, pos.start.y, pos.end.x, pos.end.y)\n\t\t\t\t\t\t\tclearPos()\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t\tme(\"#img-container\").on('htmx:afterSettle', function(evt) {\n\t\t\t\t\t\t\tcalculateCanvasBound()\n\t\t\t\t\t\t\tclearPosForm()\n\t\t\t\t\t});\n\n\t\t\t\t\n\t\t\t\t\tconst calculateCanvasBound = function (e) {\n\t\t\t\t\t\tconst img = document.getElementById(\"img\")\n\t\t\t\t\t\tcanvas.width = img.width\n\t\t\t\t\t\tcanvas.height = img.height\n\t\t\t\t\t}\n\n\t\t\t\t\tconst clearPos = function () {\n\t\t\t\t\t\tpos.start.x = null\n\t\t\t\t\t\tpos.start.y = null\n\t\t\t\t\t\tpos.end.x = null\n\t\t\t\t\t\tpos.end.y = null\n\t\t\t\t\t}\n\n\t\t\t\t\tconst getPosition = function (e) {\n\t\t\t\t\t\tconst rect = e.target.getBoundingClientRect();\n\t\t\t\t\t\tconst x = e.clientX - rect.left;\n\t\t\t\t\t\tconst y = e.clientY - rect.top;\n\t\t\t\t\t\treturn { x: x, y: y }\n\t\t\t\t\t}\n\n\t\t\t\t\tconst drawRect = function (ctx, a, b, x, y) {\n\t\t\t\t\t\tif (a !== null && b !== null && x !== null && y !== null) {\n\t\t\t\t\t\t\tif (a > x) {\n\t\t\t\t\t\t\t\t[a, x] = [x, a]\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\tif (b > y) {\n\t\t\t\t\t\t\t\t[b, y] = [y, b]\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\tctx.clearRect(0, 0, canvas.width, canvas.height);\n\t\t\t\t\t\t\tctx.strokeRect(a, b, x - a, y - b)\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t</script></div><button hx-post=\"/clipboard\" hx-target=\"#img-container\" hx-swap=\"innerHTML\">Clipboard</button><form hx-post=\"/upload\" hx-target=\"#img-container\" hx-swap=\"innerHTML\" enctype=\"multipart/form-data\"><input type=\"file\" name=\"image\" accept=\"image/*\" required> <button type=\"submit\" class=\"upload-button\">Upload</button><div class=\"loading-indicator\">Loading...</div><style>\n\t\t\t\t\t.loading-indicator{\n\t\t\t\t\t\tdisplay:none;\n\t\t\t\t\t}\n\t\t\t\t\t.htmx-request .loading-indicator{\n\t\t\t\t\t\tdisplay:inline;\n\t\t\t\t    }\n\t\t\t\t\t.htmx-request .upload-button{\n\t\t\t\t\t\tdisplay:none;\n\t\t\t\t\t}\n\t\t\t\t</style></form><form><input type=\"number\" class=\"pos-input\" name=\"a\"> <input type=\"number\" class=\"pos-input\" name=\"b\"> <input type=\"number\" class=\"pos-input\" name=\"x\"> <input type=\"number\" class=\"pos-input\" name=\"y\"><script>\n\t\t\t\t\tconst clearPosForm = function () {\n\t\t\t\t\t\tconst inputs = any(\".pos-input\")\n\t\t\t\t\t\tinputs.forEach(i => i.value = \"\")\n\t\t\t\t\t}\n\t\t\t\t\tconst fillPosForm = function (a, b, x, y) {\n\t\t\t\t\t\tconst inputs = any(\".pos-input\")\n\t\t\t\t\t\tif (a > x) {\n\t\t\t\t\t\t\t[a, x] = [x, a]\n\t\t\t\t\t\t}\n\t\t\t\t\t\tif (b > y) {\n\t\t\t\t\t\t\t[b, y] = [y, b]\n\t\t\t\t\t\t}\n\t\t\t\t\t\tinputs[0].value = a\n\t\t\t\t\t\tinputs[1].value = b\n\t\t\t\t\t\tinputs[2].value = x\n\t\t\t\t\t\tinputs[3].value = y\n\t\t\t\t\t}\n\t\t\t\t</script></form></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
