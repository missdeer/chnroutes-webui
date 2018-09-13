<html lang="zh_CN">
<head>
    <meta charset="UTF-8">
    <title>chnroutes WebUI</title>
    <meta name="viewport" content="width=device-width,initial-scale=1,user-scalable=0">
    <meta name="description" content="chnroutes WebUI">
    <style>
        body {
            text-align: center;
        }

        table {
            font-family: arial, sans-serif;
            border-collapse: collapse;
            width: 80%;
        }

        #file td, th {
            text-align: left;
            white-space: nowrap;
        }

        #file tr:nth-child(even) {
            background-color: #dddddd;
        }

        #app {
            margin: 0 auto
        }

        h2 {
            padding-top: 3%
        }

        input {
            margin: 1%;
            width: 70%
        }
    </style>
</head>
<body>
<div id="app">
    <h2>chnroutes WebUI</h2>
    <p>
    Gateway：
    <input v-model="gateway" required>
    </p>
    <button v-on:click="generate">生成文件</button>
    <div>        
        <table id="file" align="center">
            <tr>
                <th>平台</th>
                <th>文件</th>
            </tr>
            {{range $item := .items}}
            <tr>
                <td>{{ $item.Platform }}</td>
                <td><a href="{{ $item.URL}}">{{ $item.FileName }}</a></td>
            </tr>
            {{end}}
        </table>
    </div>
</div>
</body>

</html>