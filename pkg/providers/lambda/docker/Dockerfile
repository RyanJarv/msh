FROM public.ecr.aws/lambda/provided:al2

RUN yum install -y jq
RUN yum install -y python3-pip
RUN pip3 install -U awscli

WORKDIR /var/runtime/
COPY bootstrap bootstrap
RUN chmod 755 bootstrap

COPY main.sh main.sh
RUN chmod 755 main.sh

CMD [ "main.handler" ]