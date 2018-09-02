FROM python:3.6-stretch

COPY bin/exercism_linux_64 /usr/bin/exercism

RUN chmod +x /usr/bin/exercism \
        && exercism configure --token=3fe5769e-a849-4a33-9569-d652aa3cedfc --workspace=/home \
        && pip install pytest \
        && wget -O script.sh "https://gist.githubusercontent.com/zeerorg/98fa08f4a58f36cca9459db6e0db0401/raw/07d5ab47aad416bb8ce15ecdf34474963cd54c32/exercism_python_run.sh" \
        && mv script.sh /home/ \
        && chmod +x /home/script.sh

WORKDIR /home

ENTRYPOINT ["./script.sh"]

# Test : 0e3134698b1242f3998a6509b9378a9e
# Error : d7058833975546f9a5b1f20fcd48cb63
