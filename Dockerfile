FROM python:3
ADD / /
RUN pip install -r reqs.txt
CMD ["python", "./main.py"]