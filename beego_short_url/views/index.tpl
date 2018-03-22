<!DOCTYPE html>

<html>
<head>

</head>

<body>
<table border="1px">
    {{range .url_list}}
        <tr>
           <td>shorturl: </td>
           <td><a href= {{.ShortUrl}}>{{.ShortUrl}}</a></td>
        </tr>
    {{end}}
</table>
</body>
</html>
