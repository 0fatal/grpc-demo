FROM ubuntu:18.04

ENV PYTHONIOENCODING=utf-8

RUN apt-get update -y && \
    apt-get install python3-pip  -y

RUN mkdir -p /app/server

COPY ./fake_useragent_0.1.11.json /tmp/fake_useragent_0.1.11.json

WORKDIR /app/server

COPY ./requirements.txt .

RUN pip3 install --upgrade pip

RUN pip install -r requirements.txt

COPY . .

ENTRYPOINT ["python3"]

CMD ["main.py"]
