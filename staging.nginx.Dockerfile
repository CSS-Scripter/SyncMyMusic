FROM nginx
COPY './nginx/staging_server.conf' '/etc/nginx/conf.d/default.conf'
COPY './builds/dist-staging' '/var/www/html'