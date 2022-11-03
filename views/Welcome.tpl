<!DOCTYPE HTML>
<html lang="zh">
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1" />
<head>
    <title>Welcome!</title>
    <link rel="stylesheet" href="../static/css/window.css">
</head>
<body >
<h1 >你好!{{.User.Name}}<br></h1>
<form  class="box" action="../hotmain/" method="get">
    <input type="submit" value="看看抖音">
</form>
<form  class="box" action="../chat/chatmain/{{.User.Name}}/{{.User.Id}}" method="get">
    <input type="submit" value="看看抖音">
</form>
<br>
</body>
</html>