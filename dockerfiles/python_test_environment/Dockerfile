FROM python:3.6-stretch

WORKDIR /home

COPY --from=exercism_build:latest /exercism /usr/bin/exercism
COPY --from=exercism_build:latest /root/.config/exercism/ /root/.config/exercism/
COPY test_script.sh ./script.sh

RUN chmod +x /usr/bin/exercism \
        && pip install pytest

ENTRYPOINT ["bash", "./script.sh"]

# Test : 0e3134698b1242f3998a6509b9378a9e
# Error : d7058833975546f9a5b1f20fcd48cb63
