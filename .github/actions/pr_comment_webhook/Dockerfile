FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
COPY push.py ./
RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "push.py" ]
