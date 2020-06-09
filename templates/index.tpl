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

        input, select {
            margin: 1%;
            width: 30%
        }
    </style>
</head>
<body>
<div id="app">
    <h2>chnroutes WebUI</h2>
    <p>
    网关：
    <input v-model="gateway" placeholder="网关IP地址，留空则自动检测默认网关"><br>
    数据：
    <select v-model="datasource">
        <option value="chinaiplist" selected>https://github.com/17mon/china_ip_list</option>
        <option value="maxmind">https://dev.maxmind.com/geoip/geoip2/geolite2/</option>
        <option value="apnic">http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest</option>
    </select><br>
    平台：
    <select v-model="platform">
        <option value="all">全部</option>
        <option value="windows">Windows</option>
        <option value="mac">macOS</option>
        <option value="linux">Linux</option>
        <option value="android">Android</option>
        <option value="chinadns">ChinaDNS</option>
        <option value="routeros">RouterOS</option>
    </select>
    </p><p>
    <button v-on:click="generate">生成文件</button>
    </p>
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

<script src="//cdn.staticfile.org/vue/2.6.11/vue.min.js"></script>
<script>
var app = new Vue({
    el: "#app",
    data: {
        platform: "{{ .platform }}",
        gateway: "{{ .gateway }}",
        datasource: "{{ .datasource }}",
    },
	methods:{
        generate:function(){
            gw = this.gateway
            if (gw == "") {
                gw = "auto"
            }
            pf = this.platform
            if (pf == "all") {
                window.open( "/l/" + gw + "/" + this.datasource, "_self")
            } else {
                window.open( "/l/" + gw + "/" + this.datasource + "/" + this.platform, "_self")
            }
        }
    }
});
</script>
</html>