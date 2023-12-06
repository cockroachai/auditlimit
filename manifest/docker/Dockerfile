FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
# ADD resource                $WORKDIR/resource
ADD ./temp/linux_amd64/main $WORKDIR/main
# COPY config.prod.yaml       $WORKDIR/config.yaml
ADD ./data                  $WORKDIR/data
RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
