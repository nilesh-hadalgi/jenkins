FROM registry.access.redhat.com/ubi8/go-toolset:1.18.4-8.1669838000

# Labels
LABEL name="awsvmscheduler" \
    maintainer="tech@mycompany.com" \
    vendor="mycompany" \
    version="1.0.0" \
    release="1" \
    summary="This service enables AWS cloud vm start/stop." \
    description="This service enables AWS cloud vm start/stop."

ENV ec2_instanceIds "dummy-ids"

ENV ec2_command "start"

# copy code to the build path
COPY . .

RUN go mod download

RUN go build -o awsvmscheduler

CMD ["sh", "-c", "./awsvmscheduler -c  $ec2_command -i $ec2_instanceIds "]
