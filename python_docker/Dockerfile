FROM python:3.6.6-alpine3.8
MAINTAINER Abhinav Tyagi

RUN apk update
RUN apk update && \
    apk add gcc python3-dev musl-dev

WORKDIR /

COPY ./requirements.txt /requirements.txt

RUN pip install -r requirements.txt

COPY ./run.py /

ENTRYPOINT ["python"]
CMD ["run.py"]
