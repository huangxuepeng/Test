server {
        listen       80;
        server_name  localhost;

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            root   /data/www/dist;
            index  index.html index.htm;
        }
        location /u {
            proxy_pass http://127.0.0.1:8080;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

        }

        #error_page  404              /404.html;
        # redirect server error pages to the static page /50x.html
        #
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }

        # proxy the PHP scripts to Apache listening on 127.0.0.1:80
        #
        #location ~ \.php$ {
        #    proxy_pass   http://127.0.0.1;
        #}

        #!/bin/bash -ilex
echo "更新前端代码"
rm -rf /data/www/*
cp -a -f /var/lib/jenkins/workspace/vue-test/. /data/www
cd /data/www
npm install
npm run build