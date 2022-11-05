
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" >
    <title>留言版</title>
    <style type="text/css">
        span{
            width: 500px;
            height: 50px;
            border-width: 2px;
            border-style: dashed;
            border-color: #47f1da;
        }
    </style>
</head>
<body>
<a>
    <br><br><span>[{{.M.M1.SpeakerName}}]:{{.M.M1.Content}}@{{.M.M1.Time}}</span>
    <br><br><span>[{{.M.M2.SpeakerName}}]:{{.M.M2.Content}}@{{.M.M2.Time}}</span>
    <br><br><span>[{{.M.M3.SpeakerName}}]:{{.M.M3.Content}}@{{.M.M3.Time}}</span>
    <br><br><span>[{{.M.M4.SpeakerName}}]:{{.M.M4.Content}}@{{.M.M4.Time}}</span>
    <br><br><span>[{{.M.M5.SpeakerName}}]:{{.M.M5.Content}}@{{.M.M5.Time}}</span>
    <br><br><span>[{{.M.M6.SpeakerName}}]:{{.M.M6.Content}}@{{.M.M6.Time}}</span>
    <br><br><span>[{{.M.M7.SpeakerName}}]:{{.M.M7.Content}}@{{.M.M7.Time}}</span>
    <br><br><span>[{{.M.M8.SpeakerName}}]:{{.M.M8.Content}}@{{.M.M8.Time}}</span>
    <br><br><span>[{{.M.M9.SpeakerName}}]:{{.M.M9.Content}}@{{.M.M9.Time}}</span>
    <br><br><span>[{{.M.M10.SpeakerName}}]:{{.M.M10.Content}}@{{.M.M10.Time}}</span>
    <br><br><span>[{{.M.M11.SpeakerName}}]:{{.M.M11.Content}}@{{.M.M11.Time}}</span>
    <br><br><span>[{{.M.M12.SpeakerName}}]:{{.M.M12.Content}}@{{.M.M12.Time}}</span>
    <br><br><span>[{{.M.M13.SpeakerName}}]:{{.M.M13.Content}}@{{.M.M13.Time}}</span>
    <br><br><span>[{{.M.M14.SpeakerName}}]:{{.M.M14.Content}}@{{.M.M14.Time}}</span>
    <br><br><span>[{{.M.M15.SpeakerName}}]:{{.M.M15.Content}}@{{.M.M15.Time}}</span>
    <br><br><span>[{{.M.M16.SpeakerName}}]:{{.M.M16.Content}}@{{.M.M16.Time}}</span>
    <br><br><span>[{{.M.M17.SpeakerName}}]:{{.M.M17.Content}}@{{.M.M17.Time}}</span>
    <br><br><span>[{{.M.M18.SpeakerName}}]:{{.M.M18.Content}}@{{.M.M18.Time}}</span>
    <br><br><span>[{{.M.M19.SpeakerName}}]:{{.M.M19.Content}}@{{.M.M19.Time}}</span>
    <br><br><span>[{{.M.M20.SpeakerName}}]:{{.M.M20.Content}}@{{.M.M20.Time}}</span>
</a>
<br>
<br>
<a>
    <form  action="../../say/{{.User.Name}}/{{.User.Id}}" method="get">
        <input type="text" name="text" placeholder="留下你的想法">
        <input type="submit" value="发送">
    </form>
</a>
</body>
</html>
