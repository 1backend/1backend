'use strict';

var mysql = require('mysql');
var db = mysql.createConnection({
  host     : process.env['MYSQLIP'],
  user     : '{{ .Project.Author }}_{{ .Project.Name }}',
  password : process.env['INFRAPASS'],
  database : '{{ .Project.Author }}_{{ .Project.Name }}'
});
const express = require('express');
{{.Project.Imports}}
const app = express();

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
