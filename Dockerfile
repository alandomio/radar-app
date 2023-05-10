# Immagine di base con Go e OpenCV
FROM golang:latest

# Installazione delle dipendenze di OpenCV
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    cmake \
    git \
    libgtk2.0-dev \
    pkg-config \
    libavcodec-dev \
    libavformat-dev \
    libswscale-dev \
    libv4l-dev \
    libjpeg-dev \
    libpng-dev \
    libtiff-dev \
    libatlas-base-dev \
    gfortran \
    && apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Clonazione del repository di OpenCV
RUN git clone https://github.com/opencv/opencv.git && \
    cd opencv && \
    git checkout 4.x  # Sostituisci <versione_opencv> con la versione di OpenCV desiderata

# Compilazione ed installazione di OpenCV
RUN cd opencv && \
    mkdir build && \
    cd build && \
    cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local .. && \
    make -j$(nproc) && \
    make install

# Copia del file pkg-config di OpenCV nel percorso appropriato
# RUN cp /usr/local/lib/pkgconfig/opencv.pc /usr/lib/pkgconfig/

# Impostazione delle variabili d'ambiente per Go
ENV CGO_CPPFLAGS="-I/usr/local/include"
ENV CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_highgui -lopencv_imgproc -lopencv_imgcodecs"

# Impostazione della directory di lavoro
WORKDIR /go/src/app

# Copia dei file del progetto nell'immagine
COPY . .

# Compilazione dell'applicazione Go
RUN go build -o main


# Avvia l'applicazione al momento dell'avvio del container
CMD ["./main"]
