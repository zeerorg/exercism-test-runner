FROM golang:1.11

WORKDIR /home

COPY --from=exercism_build:latest /exercism /usr/bin/exercism
COPY --from=exercism_build:latest /root/.config/exercism/ /root/.config/exercism/
COPY test_script.sh ./script.sh

RUN chmod +x /usr/bin/exercism

ENTRYPOINT ["bash", "./script.sh"]

# fail: 06798ce0cf734b1db5ab79f7f3870b92
# pass: e9f052d8ec4949d9b755a97ce7543c75
