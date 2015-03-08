<html xmlns="http://www.w3.org/1999/xhtml">
		<head>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
      <title>hi4a</title>
      <style>
        table{border-collapse:collapse;border-spacing:0;border-left:1px solid #888;border-top:1px solid #888;background:#efefef;}
        th,td{border-right:1px solid #888;border-bottom:1px solid #888;padding:5px 15px;}
        th{font-weight:bold;background:#ccc;}
      </style>
		</head>
		<body>
      <br/>
			<div>
        <table>
          <thead>
            {{with .Imini}}
  						{{range .Columns}}
  							<th>{{.}}</th>
  						{{end}}
  					{{end}}
          </thead>
          <tbody>
            {{with .Imini}}
  						{{range .Data}}
  							<tr>
                  {{range .}}
                    <td>{{.}}</td>
      						{{end}}
                </tr>
  						{{end}}
  					{{end}}
          </tbody>
        </table>

			</div>
      <br/>
      <div>
      <table>
        <thead>
          {{with .Ihi4a}}
            {{range .Columns}}
              <th>{{.}}</th>
            {{end}}
          {{end}}
        </thead>
        <tbody>
          {{with .Ihi4a}}
            {{range .Data}}
              <tr>
                {{range .}}
                  <td>{{.}}</td>
                {{end}}
              </tr>
            {{end}}
          {{end}}
        </tbody>
      </table>
    </div>
		</body>
</html>
