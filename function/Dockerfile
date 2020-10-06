FROM golang:1.13

RUN apt-get update \
 && apt-get install -y openjdk-11-jre \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

# Install firebase emulator
RUN wget https://firebase.tools/bin/linux/latest -O /usr/local/bin/firebase \
 && chmod 755 /usr/local/bin/firebase

RUN firebase setup:emulators:firestore \
 && firebase setup:emulators:pubsub

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

COPY . ./

CMD ["firebase", "--project", "test", "emulators:exec", "--only", "firestore,pubsub", "make test"]
