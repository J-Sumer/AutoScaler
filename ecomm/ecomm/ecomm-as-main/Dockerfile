FROM httpd:2.4

ENV dir=/apache
WORKDIR ${dir}
# RUN apt-get install zip -y
# RUN wget -O ecomm.zip https://github.com/J-Sumer/ecomm-as/archive/refs/heads/main.zip
# RUN unzip ecomm.zip
COPY ./ /usr/local/apache2/htdocs/


# $ docker build -t my-apache2 .
# $ docker run -dit --name my-running-app -p 8080:80 my-apache2