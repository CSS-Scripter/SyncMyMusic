FROM nginx
COPY './nginx/production_server.conf' '/etc/nginx/conf.d/default.conf'
COPY './builds/dist-production' '/var/www/html'