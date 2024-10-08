// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/purylte/ocr-webui/types"

func CanvasImageContainer(containerId string, imageData *types.ImageData) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(containerId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/canvasImage.templ`, Line: 7, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CanvasImage(imageData).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><style>\n        me {\n            flex: 0 1 auto;\n            background-color: red;\n            height: 100%;\n            align-items: center;\n            justify-content: center;\n            display: flex;\n        }\n\n        me div {\n            align-items: center;\n            justify-content: center;\n            display: flex;\n        }\n    </style><script>\n        const pos = { start: { x: null, y: null }, end: { x: null, y: null } }\n\n        const clearPos = function () {\n            pos.start.x = null\n            pos.start.y = null\n            pos.end.x = null\n            pos.end.y = null\n        }\n\n        const getPosition = function (e) {\n            const rect = e.target.getBoundingClientRect();\n            const x = Math.round(e.clientX - rect.left);\n            const y = Math.round(e.clientY - rect.top);\n            return { x: x, y: y }\n        }\n\n        const drawRect = function (ctx, a, b, x, y) {\n            if (a !== null && b !== null && x !== null && y !== null) {\n                if (a > x) {\n                    [a, x] = [x, a]\n                }\n                if (b > y) {\n                    [b, y] = [y, b]\n                }\n                ctx.clearRect(0, 0, canvas.width, canvas.height);\n                ctx.strokeRect(a, b, x - a, y - b)\n            }\n        }\n\n        const getMaxSize = function (width, height, maxWidth, maxHeight) {\n            const aspectRatio = width / height;\n            let newWidth, newHeight;\n\n            if (width > maxWidth || height > maxHeight) {\n                if (maxWidth / aspectRatio <= maxHeight) {\n                    newWidth = maxWidth;\n                    newHeight = maxWidth / aspectRatio;\n                } else {\n                    newHeight = maxHeight;\n                    newWidth = maxHeight * aspectRatio;\n                }\n            } else {\n                newWidth = width;\n                newHeight = height;\n            }\n\n            return {\n                width: newWidth,\n                height: newHeight\n            };\n        }\n\n        const calculateCanvasSize = function () {\n            const canvas = document.getElementById(\"canvas\")\n            const img = canvas.nextElementSibling\n            const parentRect = canvas.parentElement.parentElement.getBoundingClientRect()\n            const originalWidth = JSON.parse(document.getElementById('originalWidth').textContent);\n            const originalHeight = JSON.parse(document.getElementById('originalHeight').textContent);\n            const { width, height } = getMaxSize(originalWidth, originalHeight, parentRect.width, parentRect.height)\n            img.width = width\n            img.height = height\n            canvas.width = width\n            canvas.height = height\n        }\n\n        me(\"div\", me()).on('htmx:afterSettle', function (evt) {\n            calculateCanvasSize()\n        })\n    </script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CanvasImage(imageData *types.ImageData) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if imageData == nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Upload or CTRL+V<style>\n</style>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templ.JSONScript("originalWidth", imageData.Width).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.JSONScript("originalHeight", imageData.Height).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <div><canvas id=\"canvas\"><script>\n            me().on('mousedown', e => {\n                clearPosForm()\n                pos.start = getPosition(e)\n                clearPosFormDebounce()\n            })\n            me().on('mousemove', e => {\n                const { x, y } = getPosition(e)\n                const canvas = document.getElementById(\"canvas\")\n                drawRect(canvas.getContext(\"2d\"), pos.start.x, pos.start.y, x, y)\n            })\n            me().on('mouseup', e => {\n                pos.end = getPosition(e)\n                const canvas = document.getElementById(\"canvas\")\n                drawRect(canvas.getContext(\"2d\"), pos.start.x, pos.start.y, pos.end.x, pos.end.y)\n                fillPosForm(pos.start.x, pos.start.y, pos.end.x, pos.end.y, canvas.width, canvas.height)\n                submitPosFormDebounce()\n                clearPos()\n            })\n        </script></canvas><img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(imageData.WebPath)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/canvasImage.templ`, Line: 130, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><script>\n        onloadAdd(_ => {\n            calculateCanvasSize()\n        })\n    </script><style>\n        me img {\n            z-index: 0;\n            position: relative;\n        }\n\n        me canvas {\n            z-index: 20;\n            position: absolute;\n        }\n    </style></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
