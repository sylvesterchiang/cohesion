<!DOCTYPE html>
<html>
<head>
    <title>No Cohesion</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap-theme.min.css">
</head>
<body>

    <div class="container">
        {{.clientid}}

        <input  type="text" class="playlistid" placeholder="playlist url">
        <button type="button" class="playlist-submit" class="btn btn-primary">Submit</button>

    </div>

    <div class="container">{{.name}}</div>
    <!--
    <div class="container">
        {{range $val := .tracks}}
        <a class="dropdown-item"></a>
        {{end}}
    </div>
    -->
    <script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
    <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
    <script>
        $(document).ready(function(){
            $('.playlist-submit').click(function(){
                $.post("playlist?id=" + $(".playlistid").val(),
                    function(data, status){
                        console.log(data);
                        console.log(status);
                    });
            });
        });
    </script>
</body>
</html>