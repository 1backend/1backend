import * as mysql from 'mysql';
import * as express from 'express';
{{.Project.Imports}}

const db = mysql.createConnection({
  host     : process.env['MYSQLIP'],
  user     : '{{ .Project.Author }}_{{ .Project.Name }}',
  password : process.env['INFRAPASS'],
  database : '{{ .Project.Author }}_{{ .Project.Name }}'
});

const app = express();

app.get('/ping', (req: express.Request, rsp: express.Response) => {
    rsp.send({"pong": true});
})

{{ range $key, $value := .Project.Endpoints }}
app.get('{{ $value.Url }}', {{ $value.Code }});
{{ end }}

const port = 8883;
app.listen(port, () => {
    console.log(`Server is running`);
});
