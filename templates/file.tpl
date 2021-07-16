<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go templates</title>
</head>
<body>
    {{ $wisdom := .}}
    <h3>Meaning of life is {{$wisdom}}</h3>
</body>
</html>