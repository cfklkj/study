1.У��nginx ,unzip�Ƿ�װ
nginx
nginx -t
nginx -s reload
2.�����ļ���ָ���ļ�
url, zip, path, restart
unzip -o xxx -d path
wget url -O  path -T timeout -t times

3.֤��
certbot --nginx certonly -d domian
certbot renew

--test
vi /etc/nginx/conf.d/test.conf 
server {
listen 443 ssl http2 default_server;
listen [::]:443 ssl http2 default_server;
listen 80;
server_name im.guiruntang.club;
ssl on;
ssl_certificate "/etc/letsencrypt/live/www.com/fullchain.pem";
ssl_certificate_key "/etc/letsencrypt/live/www.com/privkey.pem";
}