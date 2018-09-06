FROM python:3.6-stretch

WORKDIR /home

COPY bin/exercism_linux_64 /usr/bin/exercism
COPY test_script.sh ./script.sh

RUN chmod +x /usr/bin/exercism \
        && exercism configure --token=3fe5769e-a849-4a33-9569-d652aa3cedfc --workspace=/home \
        && pip install pytest

ENTRYPOINT ["bash", "./script.sh"]

# Test : 0e3134698b1242f3998a6509b9378a9e
# Error : d7058833975546f9a5b1f20fcd48cb63
