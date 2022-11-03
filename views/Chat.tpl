
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" >
    <title>Title</title>
    <link rel="stylesheet" href="../static/css/window.css">
</head>
<body>
<a class="boxshow">
    <span type="text1">[{{.M.M1.SpeakerName}}]:{{.M.M1.Content}}@{{.M.M1.Time}}</span>
    <span type="text2">[{{.M.M2.SpeakerName}}]:{{.M.M2.Content}}@{{.M.M2.Time}}</span>
    <span type="text3">[{{.M.M3.SpeakerName}}]:{{.M.M3.Content}}@{{.M.M3.Time}}</span>
    <span type="text4">[{{.M.M4.SpeakerName}}]:{{.M.M4.Content}}@{{.M.M4.Time}}</span>
    <span type="text5">[{{.M.M5.SpeakerName}}]:{{.M.M5.Content}}@{{.M.M5.Time}}</span>
    <span type="text6">[{{.M.M6.SpeakerName}}]:{{.M.M6.Content}}@{{.M.M6.Time}}</span>
    <span type="text7">[{{.M.M7.SpeakerName}}]:{{.M.M7.Content}}@{{.M.M7.Time}}</span>
    <span type="text8">[{{.M.M8.SpeakerName}}]:{{.M.M8.Content}}@{{.M.M8.Time}}</span>
    <span type="text9">[{{.M.M9.SpeakerName}}]:{{.M.M9.Content}}@{{.M.M9.Time}}</span>
    <span type="text10">[{{.M.M10.SpeakerName}}]:{{.M.M10.Content}}@{{.M.M10.Time}}</span>
    <span type="text11">[{{.M.M11.SpeakerName}}]:{{.M.M11.Content}}@{{.M.M11.Time}}</span>
    <span type="text12">[{{.M.M12.SpeakerName}}]:{{.M.M12.Content}}@{{.M.M12.Time}}</span>
</a>
<a class="boxsay">
    <form class="boxsay" action="../chat/say/{{.User.Name}}/{{.User.Id}}" method="get">
        <input type="text" name="content" placeholder="说点什么">
        <input type="submit" value="发送">
    </form>
</a>
</body>
</html>
