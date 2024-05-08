<html>
<head>
 <title>Cars</title>
</head>
<body>
 <h1>Cars</h1>
 <table>
 <tr>
 <th>Id</th>
 <th>Make</th>
 <th>Model</th>
 <th>Year</th>
 </tr>
 {{range .cars}}
 <tr>
 <td>{{.Id}}</td>
 <td>{{.Make}}</td>
 <td>{{.Model}}</td>
 <td>{{.Year}}</td>
 </tr>
{{end}}
 </table>
 <p>Total number of cars: {{.num}}</p>
</body>
</html>