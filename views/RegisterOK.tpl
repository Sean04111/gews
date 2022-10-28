<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册成功！</title>
  <link rel="stylesheet" href="../static/css/window.css">
</head>
<body>
<form class="box" action="../login/login" method="get">
  <h1 >恭喜你!{{.User.Name}}<br>注册成功</h1>
  <input type="submit" value="去登录">
</form>
</body>
</html>
