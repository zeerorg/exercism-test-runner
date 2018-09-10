FROM golang:1.11

WORKDIR /home

COPY --from=exercism_build:latest /exercism /usr/bin/exercism
COPY test_script.sh ./script.sh

RUN chmod +x /usr/bin/exercism \
        && exercism configure --token=3fe5769e-a849-4a33-9569-d652aa3cedfc --workspace=/home

ENTRYPOINT ["bash", "./script.sh"]

# fail: 06798ce0cf734b1db5ab79f7f3870b92
# pass: e9f052d8ec4949d9b755a97ce7543c75
