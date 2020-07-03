package html_template

import (
	"html/template"
	"log"
	"os"
	"testing"
)

func Test_html_template(t *testing.T) {
	const tpl = `<html>
    <head>
        <meta charset="UTF-8" />
        </head>
        <body>
          <div id="ac">{{.AcNonce}}</div>
          <p id="sign"></p>
		</body>
        <script src='https://sf1-ttcdn-tos.pstatp.com/obj/rc-web-sdk/acrawler.js'></script>
        <script>
            var acEle = document.getElementById('ac')
            console.log(acEle.innerText)
            window.byted_acrawler.init({aid:99999999,dfp:!0});
            var f= acEle.innerText;
            console.log(f)
            var g=window.byted_acrawler.sign("",f);
            console.log(g)
            var signEle =  document.getElementById('sign')
            signEle.innerText=g
        </script>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	tp, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct{ AcNonce string }{
		AcNonce: "05ef884c100927e0fb59e",
	}
	//var buffer bytes.Buffer
	//writer := bufio.NewWriter(&buffer)
	err = tp.Execute(os.Stdout, data)
	check(err)
	//fmt.Println(string(buffer.Bytes()))

}
