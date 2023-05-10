# Immagine di base con Go e OpenCV
FROM golang:latest

ENV CUSTOM_BUILD_DEPS \
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
    gfortran 


# Installazione delle dipendenze di OpenCV
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    $CUSTOM_BUILD_DEPS && \ 
    # Clonazione del repository di OpenCV
    git clone https://github.com/opencv/opencv.git && \
    cd opencv && \
    git checkout 4.x && \
    # Compilazione ed installazione di OpenCV
    mkdir build && \
    cd build && \
    cmake -D BUILD_opencv_core=ON -D BUILD_opencv_imgproc=ON -D BUILD_opencv_imgcodecs=ON -D BUILD_opencv_highgui=ON -D BUILD_opencv_features2d=ON -D BUILD_opencv_calib3d=ON -D BUILD_opencv_objdetect=ON -D BUILD_opencv_photo=ON -D BUILD_opencv_video=ON -D BUILD_opencv_videoio=ON -D BUILD_opencv_dnn=ON -D BUILD_opencv_shape=ON -D BUILD_opencv_viz=ON -D BUILD_opencv_xfeatures2d=ON -D BUILD_opencv_stitching=ON -DOPENCV_GENERATE_PKGCONFIG=ON -D CMAKE_INSTALL_PREFIX=/usr/local .. && \
    cmake -D BUILD_opencv_imgproc=ON -D BUILD_opencv_imgcodecs=ON -DOPENCV_GENERATE_PKGCONFIG=ON -D CMAKE_INSTALL_PREFIX=/usr/local .. && \
    make -j$(4) && \
    make install && \
    pkg-config --cflags --libs opencv4 && \
    # remove openCV source code
    cd / && \
    rm -rf /opencv && \
    apt-get remove -y $CUSTOM_BUILD_DEPS  && \
    dpkg --purge $(dpkg -l | awk '/^rc/ { print $2 }') && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*



# Copia del file pkg-config di OpenCV nel percorso appropriato
# RUN cp /usr/local/lib/pkgconfig/opencv.pc /usr/lib/pkgconfig/

# Impostazione delle variabili d'ambiente per Go
ENV CGO_CPPFLAGS="-I/usr/local/include"
ENV CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_highgui -lopencv_imgproc -lopencv_imgcodecs"

# Crea il file openCV4.pc nella cartella /usr/local/lib/pkgconfig/ e inserisci le seguenti righe:
# prefix=/usr/local
# exec_prefix=${prefix}
# libdir=${exec_prefix}/lib
# includedir=${includedir_new}
# Name: OpenCV
# Description: Open Source Computer Vision Library
# Version: 4.x.x
# Cflags: -I${includedir}/opencv -I${includedir}/opencv2
# Libs: -L${libdir} -lopencv_calib3d -lopencv_imgproc -lopencv_contrib -lopencv_legacy -lopencv_core -lopencv_ml -lopencv_features2d -lopencv_objdetect -lopencv_flann -lopencv_video -lopencv_highgui


# Installazione di GoCV
RUN go get -u -d gocv.io/x/gocv





# Impostazione della directory di lavoro
WORKDIR /go/src/app

# Copia dei file del progetto nell'immagine
COPY . .

# Compilazione dell'applicazione Go
# RUN go build -o main


# Avvia l'applicazione al momento dell'avvio del container
# CMD ["./main"]
CMD [""]
