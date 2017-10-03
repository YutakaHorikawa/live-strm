FROM centos:7

RUN set -x \
	yum install -y epel-release \
	&& yum -y update \
	&& yum install -y wget git gcc pcre-devel openssl openssl-devel \
	&& cd /usr/local/lib  && git clone https://github.com/arut/nginx-rtmp-module.git \
	&& cd ~/ && wget https://nginx.org/download/nginx-1.13.5.tar.gz && tar zxf nginx-1.13.5.tar.gz \
	&& cd ~/nginx-1.13.5 && ./configure --sbin-path=/usr/sbin/nginx --conf-path=/etc/nginx/nginx.conf --error-log-path=/var/log/nginx/error.log --http-log-path=/var/log/nginx/access.log --pid-path=/var/run/nginx.pid --lock-path=/var/run/nginx.lock --prefix=/usr/local/nginx --user=nginx --group=nginx --with-http_ssl_module --with-http_realip_module --with-http_addition_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --with-http_mp4_module --with-http_gzip_static_module --with-http_random_index_module --with-http_secure_link_module --with-http_stub_status_module --with-ipv6 --add-module=/usr/local/lib/nginx-rtmp-module \
	&& cd ~/nginx-1.13.5 && make && make install && groupadd nginx && useradd -g nginx nginx && usermod -s /bin/false nginx \
	&& mkdir -p /var/log/nginx \
	&& chown nginx:nginx /var/log/nginx


COPY nginx.conf /etc/nginx/

EXPOSE 1935 

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
#CMD ["tail", "-f", "/dev/null"]
