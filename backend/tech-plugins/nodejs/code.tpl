'use strict';

const mysql = require('mysql');
const express = require('express');
const bodyParser = require('body-parser');
{{.Project.Imports}}

var db = mysql.createConnection({
  host     : process.env['MYSQL_IP'],
  user     : process.env['MYSQL_USER'],
  password : process.env['MYSQL_PASSWORD'],
  database : process.env['MYSQL_DB']
});

const app = express();
app.use(bodyParser.json());

app.get('/ping', (req, res) => {
    res.send({"pong": true});
});

{{ range $key, $value := .Project.Endpoints }}
app.get('{{ $value.Url }}', {{ $value.Code }});
{{ end }}

const port = 8883;
app.listen(port, () => {
    console.log(`Server is running`);
});
