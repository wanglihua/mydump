package main

var displayJobStatusTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8"/>
    <style type="text/css">
        table.jobStatus {
            border-collapse: collapse;
            font-size: 12px;
        }

        table.jobStatus th, table.jobStatus td {
            border-style: solid;
            border-color: gray;
            border-width: 1px;
            padding: 5px;
            text-align: center;
        }
    </style>
</head>
<body>
<table class="jobStatus">
    <thead>
    <tr>
        <th style="width: 180px;">上一次运行</th>
        <th style="width: 180px;">下一次运行</th>
    </tr>
    </thead>
    <tbody>
    {{ range .entries }}
    {{ $job := .Job }}
    <tr>
        <td>{{ if not .Prev.IsZero }}{{ .Prev.Format "2006-01-02 15:04:05" }}{{ end }}</td>
        <td>{{ if not .Next.IsZero }}{{ .Next.Format "2006-01-02 15:04:05" }}{{ end }}</td>
    </tr>
    {{end}}
    </tbody>
</table>
</body>
</html>
`
