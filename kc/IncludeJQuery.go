package kc

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func IncludeJquery() chromedp.ActionFunc {
	script := `
function loadJs(url, callback) {
    var script = document.createElement('script');
    script.type = "text/javascript";
    if (typeof (callback) != "undefined") {
        if (script.readyState) {
            script.onreadystatechange = function () {
                if (script.readyState == "loaded" || script.readyState == "complete") {
                    script.onreadystatechange = null;
                    callback();
                }
            }
        } else {
            script.onload = function () {
                callback();
            }
        }
    }
    script.src = url;
    document.body.appendChild(script);
}

window.onload = function () {
    loadJs("https://lib.baomitu.com/jquery/3.5.0/jquery.min.js", function () {
    });
};



`
	var scriptID page.ScriptIdentifier
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		scriptID, err = page.AddScriptToEvaluateOnNewDocument(script).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
