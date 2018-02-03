import * as mysql from 'mysql';
import * as express from 'express';
import * as bodyParser from 'body-parser';
{{.Project.Imports}}

const sql = mysql.createConnection({
  host     : process.env['MYSQL_IP'],
  user     : process.env['MYSQL_USER'],
  password : process.env['MYSQL_PASSWORD'],
  database : process.env['MYSQL_DB']
});

const app = express();
app.use(bodyParser.json());

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
